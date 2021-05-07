var app=new Vue({
    el: '#app',
    delimiters:["<{","}>"],
    data: {
        visible:false,
        chatTitleType:"info",
        fullscreenLoading:true,
        leftTabActive:"first",
        rightTabActive:"visitorInfo",
        users:[],
        usersMap:[],
        server:getWsBaseUrl()+"/ws_kefu?token="+localStorage.getItem("token"),
        //server:getWsBaseUrl()+"/chat_server",
        socket:null,
        messageContent:"",
        currentGuest:"",
        msgList:[],
        chatTitle:"暂时未处理咨询",
        chatInputing:"",
        kfConfig:{
            id : "kf_1",
            name : "客服丽丽",
            avator : "",
            to_id : "",
        },
        visitor:{
            visitor_id:"",
            refer:"",
            client_ip:"",
            city:"",
            status:"",
            source_ip:"",
            created_at:"",
        },
        visitorExtra:[],
        visitors:[],
        visitorCount:0,
        visitorCurrentPage:1,
        visitorPageSize:10,
        face:[],
        transKefuDialog:false,
        otherKefus:[],
        replyGroupDialog:false,
        replyContentDialog:false,
        editReplyContentDialog:false,
        replySearch:"",
        replySearchList:[],
        replySearchListActive:[],
        groupName:"",
        groupId:"",
        replys:[],
        replyId:"",
        replyContent:"",
        replyTitle:"",
        ipBlacks:[],
        sendDisabled:false,
    },
    methods: {
        //跳转
        openUrl(url) {
            window.location.href = url;
        },
        sendKefuOnline(){
            let mes = {}
            mes.type = "kfOnline";
            mes.data = this.kfConfig;
            this.socket.send(JSON.stringify(mes));
        },
        //心跳
        ping(){
            let _this=this;
            let mes = {}
            mes.type = "ping";
            mes.data = "";
            setInterval(function () {
                if(_this.socket!=null){
                    _this.socket.send(JSON.stringify(mes));
                }
            },20000)
            setInterval(function(){
                _this.getOnlineVisitors();
            },120000);
        },
        //初始化websocket
        initConn() {
            let socket = new ReconnectingWebSocket(this.server);//创建Socket实例
            this.socket = socket
            this.socket.onmessage = this.OnMessage;
            this.socket.onopen = this.OnOpen;
        },
        OnOpen() {
            this.sendKefuOnline();
        },
        OnMessage(e) {
            const redata = JSON.parse(e.data);
            switch (redata.type){
                case "inputing":
                    this.handleInputing(redata.data);
                    //this.sendKefuOnline();
                    break;
                case "allUsers":
                    this.handleOnlineUsers(redata.data);
                    //this.sendKefuOnline();
                    break;
                case "userOnline":
                    this.addOnlineUser(redata.data);


                    break;
                case "userOffline":
                    this.removeOfflineUser(redata.data);
                    //this.sendKefuOnline();
                    break;
                case "notice":
                    //发送通知
                    var _this=this;
                    window.parent.postMessage({
                        name:redata.data.username,
                        body: redata.data.content,
                        icon: redata.data.avator
                    });
                    _this.alertSound();
                    break;
            }


            if (redata.type == "message") {
                let msg = redata.data
                let content = {}
                let _this=this;
                content.avator = msg.avator;
                content.name = msg.name;
                content.content = replaceContent(msg.content);
                content.is_kefu = msg.is_kefu=="yes"? true:false;
                content.time = msg.time;
                if (msg.id == this.currentGuest) {
                    this.msgList.push(content);
                }

                for(let i=0;i<this.users.length;i++){
                    if(this.users[i].uid==msg.id){
                        this.$set(this.users[i],'last_message',msg.content);
                        if(this.visitor.visitor_id!=msg.id){
                            this.$set(this.users[i],'hidden_new_message',false);
                        }
                    }
                }
                this.scrollBottom();
                if(content.is_kefu){
                    return;
                }
                window.parent.postMessage({
                    name:msg.name,
                    body: msg.content,
                    icon: msg.avator

                });
                _this.alertSound();
                _this.chatInputing="";
            }
        },
        //接手客户
        talkTo(guestId,name) {
            this.currentGuest = guestId;
            //this.chatTitle=name+"|"+guestId+",正在处理中...";

            //发送给客户
            let mes = {}
            mes.type = "kfConnect";
            this.kfConfig.to_id=guestId;
            mes.data = this.kfConfig;
            this.socket.send(JSON.stringify(mes));

            //获取当前访客信息
            this.getVistorInfo(guestId);
            //获取当前客户消息
            this.getMesssagesByVisitorId(guestId);
            for(var i=0;i<this.users.length;i++){
                if(this.users[i].uid==guestId){
                    this.$set(this.users[i],'hidden_new_message',true);
                }
            }
        },
        //发送给客户
        chatToUser() {
            this.messageContent=this.messageContent.trim("\r\n");
            this.messageContent=this.messageContent.replace("\n","");
            this.messageContent=this.messageContent.replace("\r\n","");
            if(this.messageContent==""||this.messageContent=="\r\n"||this.currentGuest==""){
                return;
            }
            if(this.sendDisabled){
                return;
            }
            this.sendDisabled=true;
            let _this=this;
            let mes = {};
            mes.type = "kefu";
            mes.content = this.messageContent;
            mes.from_id = this.kfConfig.id;
            mes.to_id = this.currentGuest;
            mes.content = this.messageContent;
            $.post("/2/message",mes,function(res){
                _this.sendDisabled=false;
                if(res.code!=200){
                    _this.$message({
                        message: res.msg,
                        type: 'error'
                    });
                    return;
                }
                _this.messageContent = "";
               _this.sendSound();
            });

            // let content = {}
            // content.avator = this.kfConfig.avator;
            // content.name = this.kfConfig.name;
            // content.content = replaceContent(this.messageContent);
            // content.is_kefu = true;
            // content.time = '';
            // this.msgList.push(content);
            _this.sendDisabled=false;
            this.scrollBottom();
        },
        //处理当前在线用户列表
        addOnlineUser:function (retData) {
            var flag=false;
            retData.last_message=retData.last_message;
            retData.status=1;
            retData.name=retData.username;
            retData.hidden_new_message=true;
            for(let i=0;i<this.users.length;i++){
                if(this.users[i].uid==retData.uid){
                    flag=true;
                }
            }
            if(!flag){
                this.users.unshift(retData);
            }
            for(let i=0;i<this.visitors.length;i++){
                if(this.visitors[i].visitor_id==retData.uid){
                    this.visitors[i].status=1;
                    break;
                }
            }
            if(this.visitor.visitor_id==retData.uid){
                this.getVistorInfo(retData.uid)
            }

        },
        //处理当前在线用户列表
        removeOfflineUser:function (retData) {
            for(let i=0;i<this.users.length;i++){
                if(this.users[i].uid==retData.uid){
                    this.users.splice(i,1);
                }
            }
            let vid=retData.uid;
            for(let i=0;i<this.visitors.length;i++){
                if(this.visitors[i].visitor_id==vid){
                    this.visitors[i].status=0;
                    break;
                }
            }
        },
        //处理当前在线用户列表
        handleOnlineUsers:function (retData) {
            if (this.currentGuest == "") {
                this.chatTitle = "连接成功,等待处理中...";
            }
            this.usersMap=[];
            for(let i=0;i<retData.length;i++){
                this.usersMap[retData[i].uid]=retData[i].username;
                retData[i].last_message="新访客";
            }
            if(this.users.length==0){
                this.users = retData;
            }
            for(let i=0;i<this.visitors.length;i++){
                let vid=this.visitors[i].visitor_id;
                if(typeof this.usersMap[vid]=="undefined"){
                    this.visitors[i].status=0;
                }else{
                    this.visitors[i].status=1;
                }
            }

        },
        //处理正在输入
        handleInputing:function (retData) {
            if(retData.from==this.visitor.visitor_id){
                this.chatInputing="|正在输入："+retData.content+"...";
                if(retData.content==""){
                    this.chatInputing="";
                }
            }
            for(var i=0;i<this.users.length;i++){
                if(this.users[i].uid==retData.from){
                    this.$set(this.users[i],'last_message',retData.content+"...");
                }
            }
        },
        //获取客服信息
        getKefuInfo(){
            let _this=this;
            $.ajax({
                type:"get",
                url:"/kefuinfo",
                headers:{
                    "token":localStorage.getItem("token")
                },
                success: function(data) {
                    if(data.code==200 && data.result!=null){
                        _this.kfConfig.id=data.result.id;
                        _this.kfConfig.name=data.result.name;
                        _this.kfConfig.avator=data.result.avator;
                        _this.initConn();
                    }
                    if(data.code!=200){
                        _this.$message({
                            message: data.msg,
                            type: 'error'
                        });
                    }
                }
            });
        },
        //获取客服信息
        getOnlineVisitors(){
            let _this=this;
            $.ajax({
                type:"get",
                url:"/visitors_kefu_online",
                headers:{
                    "token":localStorage.getItem("token")
                },
                success: function(data) {
                    if(data.code==200 && data.result!=null){
                        _this.users=data.result;
                        for(var i=0;i<_this.users.length;i++){
                            _this.$set(_this.users[i],'hidden_new_message',true);
                        }
                    }
                    if(data.code!=200){
                        _this.$message({
                            message: data.msg,
                            type: 'error'
                        });
                    }
                    if(data.code==400){
                        window.location.href="/login";
                    }
                }
            });
        },
        //获取信息列表
        getMesssagesByVisitorId(visitorId,isAll){
            let _this=this;
            $.ajax({
                type:"get",
                url:"/messages?visitorId="+visitorId,
                headers:{
                    "token":localStorage.getItem("token")
                },
                success: function(data) {
                    if(data.code==200 && data.result!=null){
                        let msgList=data.result;
                        _this.msgList=[];
                        if(!isAll&&msgList.length>10){
                            var i=msgList.length-10
                        }else{
                            var i=0;
                        }
                        for(;i<msgList.length;i++){
                            let visitorMes=msgList[i];
                            let content = {}
                            if(visitorMes["mes_type"]=="kefu"){
                                content.is_kefu = true;
                                content.avator = visitorMes["kefu_avator"];
                                content.name = visitorMes["kefu_name"];
                            }else{
                                content.is_kefu = false;
                                content.avator = visitorMes["visitor_avator"];
                                content.name = visitorMes["visitor_name"];
                            }
                            content.content = replaceContent(visitorMes["content"]);
                            content.time = visitorMes["time"];
                            _this.msgList.push(content);
                            _this.scrollBottom();
                        }
                    }
                    if(data.code!=200){
                        _this.$message({
                            message: data.msg,
                            type: 'error'
                        });
                    }
                    if(data.code==400){
                        window.location.href="/login";
                    }
                }
            });
        },
        //获取客服信息
        getVistorInfo(vid){
            let _this=this;
            $.ajax({
                type:"get",
                url:"/visitor",
                data:{visitorId:vid},
                headers:{
                    "token":localStorage.getItem("token")
                },
                success: function(data) {
                    if(data.result!=null){
                        let r=data.result;
                        _this.visitor=r;
                        // _this.visitor.created_at=r.created_at;
                        // _this.visitor.refer=r.refer;
                        // _this.visitor.city=r.city;
                        // _this.visitor.client_ip=r.client_ip;
                        // _this.visitor.source_ip=r.source_ip;
                        _this.visitor.status=r.status==1?"在线":"离线";

                        //_this.visitor.visitor_id=r.visitor_id;
                        _this.chatTitle="#"+r.id+"|"+r.name;
                        _this.chatTitleType="success";
                        _this.visitorExtra=[];
                        if(r.extra!=""){
                            var extra=JSON.parse(b64ToUtf8(r.extra));
                            if (typeof extra=="string"){
                                extra=JSON.parse(extra);
                            }
                            for(var key in extra){
                                if(extra[key]==""){
                                    extra[key]="无";
                                }
                                if(key=="visitorAvatar"||key=="visitorName") continue;
                                var temp={key:key,val:extra[key]}
                                _this.visitorExtra.push(temp);
                            }
                        }

                    }
                    if(data.code!=200){
                        _this.$message({
                            message: data.msg,
                            type: 'error'
                        });
                    }
                }
            });
        },
        //关闭访客
        closeVisitor(visitorId){
            let _this=this;
            $.ajax({
                type:"get",
                url:"/2/message_close",
                data:{visitor_id:visitorId},
                headers:{
                    "token":localStorage.getItem("token")
                },
                success: function(data) {
                    if(data.code!=200){
                        _this.$message({
                            message: data.msg,
                            type: 'error'
                        });
                    }
                }
            });
        },
        //处理tab切换
        handleTabClick(tab, event){
            let _this=this;
            if(tab.name=="first"){
                this.getOnlineVisitors();
            }
            if(tab.name=="second"){
                this.getVisitorPage(1);
            }
            if(tab.name=="blackList"){
            }
        },
        //所有访客分页展示
        visitorPage(page){
            this.getVisitorPage(page);
        },
        //获取访客分页
        getVisitorPage(page){
            let _this=this;
            $.ajax({
                type:"get",
                url:"/visitors",
                data:{page:page},
                headers:{
                    "token":localStorage.getItem("token")
                },
                success: function(data) {
                    if(data.result.list!=null){
                        _this.visitors=data.result.list;
                        _this.visitorCount=data.result.count;
                        _this.visitorPageSize=data.result.pagesize;
                    }
                    if(data.code!=200){
                        _this.$message({
                            message: data.msg,
                            type: 'error'
                        });
                    }
                }
            });
        },
        replaceContent(content){
            return replaceContent(content)
        },
        //滚到底部
        scrollBottom(){
            this.$nextTick(() => {
                $('.chatBox').scrollTop($(".chatBox")[0].scrollHeight);
            });
        },
        //jquery
        initJquery(){
            this.$nextTick(() => {
                var _this=this;
                $(function () {
                    //展示表情
                    var faces=placeFace();
                    $.each(faceTitles, function (index, item) {
                        _this.face.push({"name":item,"path":faces[item]});
                    });
                    $(".faceBtn").click(function(){
                        var status=$('.faceBox').css("display");
                        if(status=="block"){
                            $('.faceBox').hide();
                        }else{
                            $('.faceBox').show();
                        }
                        return false;
                    });
                });
            });
            var _hmt = _hmt || [];
            (function() {
                var hm = document.createElement("script");
                hm.src = "https://hm.baidu.com/hm.js?82938760e00806c6c57adee91f39aa5e";
                var s = document.getElementsByTagName("script")[0];
                s.parentNode.insertBefore(hm, s);
            })();
        },
        //表情点击事件
        faceIconClick(index){
            $('.faceBox').hide();
            this.messageContent+="face"+this.face[index].name;
        },
        //上传图片
        uploadImg (url){
            let _this=this;
            $('#uploadImg').after('<input type="file" accept="image/gif,image/jpeg,image/jpg,image/png" id="uploadImgFile" name="file" style="display:none" >');
            $("#uploadImgFile").click();
            $("#uploadImgFile").change(function (e) {
                var formData = new FormData();
                var file = $("#uploadImgFile")[0].files[0];
                formData.append("imgfile",file); //传给后台的file的key值是可以自己定义的
                filter(file) && $.ajax({
                    url: url || '',
                    type: "post",
                    data: formData,
                    contentType: false,
                    processData: false,
                    dataType: 'JSON',
                    mimeType: "multipart/form-data",
                    success: function (res) {
                        if(res.code!=200){
                            _this.$message({
                                message: res.msg,
                                type: 'error'
                            });
                        }else{
                            _this.messageContent+='img[/' + res.result.path + ']';
                            _this.chatToUser();
                        }
                    },
                    error: function (data) {
                        console.log(data);
                    }
                });
            });
        },
        //上传文件
        uploadFile:function (url){
            let _this=this;
            $('#uploadFile').after('<input type="file"  id="uploadRealFile" name="file2" style="display:none" >');
            $("#uploadRealFile").click();
            $("#uploadRealFile").change(function (e) {
                var formData = new FormData();
                var file = $("#uploadRealFile")[0].files[0];
                formData.append("realfile",file); //传给后台的file的key值是可以自己定义的
                console.log(formData);
                $.ajax({
                    url: url || '',
                    type: "post",
                    data: formData,
                    contentType: false,
                    processData: false,
                    dataType: 'JSON',
                    mimeType: "multipart/form-data",
                    success: function (res) {

                        if(res.code!=200){
                            _this.$message({
                                message: res.msg,
                                type: 'error'
                            });
                        }else{
                            _this.messageContent+='file[/' + res.result.path + ']';
                            _this.chatToUser();
                        }
                    },
                    error: function (data) {
                        console.log(data);
                    }
                });
            });
        },
        addIpblack(ip){
            let _this=this;
            $.ajax({
                type:"post",
                url:"/ipblack",
                data:{ip:ip},
                headers:{
                    "token":localStorage.getItem("token")
                },
                success: function(data) {
                    if(data.code!=200){
                        _this.$message({
                            message: data.msg,
                            type: 'error'
                        });
                    }else{
                        _this.$message({
                            message: data.msg,
                            type: 'success'
                        });
                    }
                }
            });
        },
        //粘贴上传图片
        onPasteUpload(event){
            let items = event.clipboardData && event.clipboardData.items;
            let file = null
            if (items && items.length) {
                // 检索剪切板items
                for (var i = 0; i < items.length; i++) {
                    if (items[i].type.indexOf('image') !== -1) {
                        file = items[i].getAsFile()
                    }
                }
            }
            if (!file) {
                return;
            }
            let _this=this;
            var formData = new FormData();
            formData.append('imgfile', file);
            $.ajax({
                url: '/uploadimg',
                type: "post",
                data: formData,
                contentType: false,
                processData: false,
                dataType: 'JSON',
                mimeType: "multipart/form-data",
                success: function (res) {
                    if(res.code!=200){
                        _this.$message({
                            message: res.msg,
                            type: 'error'
                        });
                    }else{
                        _this.messageContent+='img[/' + res.result.path + ']';
                        _this.chatToUser();
                    }
                },
                error: function (data) {
                    console.log(data);
                }
            });
        },
        openUrl(url){
            window.open(url);
        },
        //提示音
        alertSound(){
            var b = document.getElementById("chatMessageAudio");
            var p = b.play();
            p && p.then(function(){}).catch(function(e){});
        },
        sendSound(){
            var b = document.getElementById("chatMessageSendAudio");
            var p = b.play();
            p && p.then(function(){}).catch(function(e){});
        },
        //转移客服
        transKefu(){
            this.transKefuDialog=true;
            var _this=this;
            this.sendAjax("/other_kefulist","get",{},function(result){
                _this.otherKefus=result;
            });
        },
        //转移访客客服
        transKefuVisitor(kefu,visitorId){
            var _this=this;
            this.sendAjax("/trans_kefu","get",{kefu_id:kefu,visitor_id:visitorId},function(result){
                //_this.otherKefus=result;
                _this.transKefuDialog = false
            });
        },
        //保存回复分组
        addReplyGroup(){
            var _this=this;
            this.sendAjax("/reply","post",{group_name:_this.groupName},function(result){
                //_this.otherKefus=result;
                _this.replyGroupDialog = false
                _this.groupName="";
                _this.getReplys();
            });
        },
        //添加回复内容
        addReplyContent(){
            var _this=this;
            this.sendAjax("/reply_content","post",{group_id:_this.groupId,item_name:_this.replyTitle,content:_this.replyContent},function(result){
                //_this.otherKefus=result;
                _this.replyContentDialog = false
                _this.replyContent="";
                _this.getReplys();
            });
        },
        //获取快捷回复
        getReplys(){
            var _this=this;
            this.sendAjax("/replys","get",{},function(result){
                _this.replys=result;
            });
        },
        //删除回复
        deleteReplyGroup(id){
            var _this=this;
            this.sendAjax("/reply?id="+id,"delete",{},function(result){
                _this.getReplys();
            });
        },
        //删除回复
        deleteReplyContent(id){
            var _this=this;
            this.sendAjax("/reply_content?id="+id,"delete",{},function(result){
                _this.getReplys();
            });
        },
        //编辑回复
        editReplyContent(save,id,title,content){
            var _this=this;
            if(save=='yes'){
                var data={
                    reply_id:this.replyId,
                    reply_title:this.replyTitle,
                    reply_content:this.replyContent
                }
                this.sendAjax("/reply_content_save","post",data,function(result){
                    _this.editReplyContentDialog=false;
                    _this.getReplys();
                });
            }else{
                this.editReplyContentDialog=true;
                this.replyId=id;
                this.replyTitle=title;
                this.replyContent=content;
            }

        },
        //搜索回复
        searchReply(){
            var _this=this;
            _this.replySearchListActive=[];
            if(this.replySearch==""){
                _this.replySearchList=[];
            }
            this.sendAjax("/reply_search","post",{search:this.replySearch},function(result){
                _this.replySearchList=result;
                for (var i in result) {
                    _this.replySearchListActive.push(result[i].group_id);
                }
            });
        },
        //获取黑名单
        getIpblacks(){
            var _this=this;
            this.sendAjax("/ipblacks","get",{},function(result){
                _this.ipBlacks=result;
            });
        },
        //删除黑名单
        delIpblack(ip){
            let _this=this;
            this.sendAjax("/ipblack?ip="+ip,"DELETE",{ip:ip},function(result){
                _this.sendAjax("/ipblacks","get",{},function(result){
                    _this.ipBlacks=result;
                });
            });
        },
        //划词搜索
        selectText(){
            return false;
            var _this=this;
            $('body').click(function(){
                try{
                    var selecter = window.getSelection().toString();
                    if (selecter != null && selecter.trim() != ""){
                        _this.replySearch=selecter.trim();
                        _this.searchReply();
                    }else{
                        _this.replySearch="";
                    }
                } catch (err){
                    var selecter = document.selection.createRange();
                    var s = selecter.text;
                    if (s != null && s.trim() != ""){
                        _this.replySearch=s.trim();
                        _this.searchReply();
                    }else{
                        _this.replySearch="";
                    }
                }
                var status=$('.faceBox').css("display");
                if(status=="block"){
                    $('.faceBox').hide();
                }
            });
        },
        sendAjax(url,method,params,callback){
            let _this=this;
            $.ajax({
                type: method,
                url: url,
                data:params,
                headers: {
                    "token": localStorage.getItem("token")
                },
                error:function(res){
                    var data=JSON.parse(res.responseText);
                    console.log(data);
                    if(data.code!=200){
                        _this.$message({
                            message: data.msg,
                            type: 'error'
                        });
                    }
                },
                success: function(data) {
                    if(data.code!=200){
                        _this.$message({
                            message: data.msg,
                            type: 'error'
                        });
                    }else if(data.result!=null){
                        callback(data.result);
                    }else{
                        callback(data);
                    }
                }
            });
        },
    },
    mounted() {
        document.addEventListener('paste', this.onPasteUpload)
    },
    created: function () {
        //jquery
        this.initJquery();
        this.getKefuInfo();
        this.getOnlineVisitors();
        this.getReplys();
        this.getIpblacks();
        this.selectText();
        //心跳
        this.ping();
    }
})
