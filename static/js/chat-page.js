KEFU_ID=KEFU_ID!=""? KEFU_ID:"kefu2";
new Vue({
    el: '#app',
    delimiters:["<{","}>"],
    data: {
        window:window,
        server:getWsBaseUrl()+"/ws_visitor",
        socket:null,
        msgList:[],
        msgListNum:[],
        messageContent:"",
        chatTitle:GOFLY_LANG[LANG]['connecting'],
        visitor:{},
        face:[],
        showKfonline:false,
        socketClosed:false,
        focusSendConn:false,
        timer:null,
        sendDisabled:false,
        flyLang:GOFLY_LANG[LANG],
        showIconBtns:false,
        showFaceIcon:false,
        isIframe:false,
        kefuInfo:{},
        showLoadMore:false,
        messages:{
            page:1,
            pagesize:5,
            list:[],
        },
    },
    methods: {
        //初始化websocket
        initConn:function() {
            let socket = new ReconnectingWebSocket(this.server+"?visitor_id="+this.visitor.visitor_id);//创建Socket实例
            this.socket = socket
            this.socket.onmessage = this.OnMessage;
            this.socket.onopen = this.OnOpen;
            this.socket.onclose = this.OnClose;
            //心跳
            this.ping();
        },
        OnOpen:function() {
            console.log("ws:onopen");
            //获取欢迎
            this.getNotice();
            this.socketClosed=false;
            this.focusSendConn=false;
        },
        OnMessage:function(e) {
            console.log("ws:onmessage");
            this.socketClosed=false;
            this.focusSendConn=false;
            const redata = JSON.parse(e.data);
            if (redata.type == "kfOnline") {
                let msg = redata.data
                if(this.showKfonline && this.visitor.to_id==msg.id){
                    return;
                }
                this.visitor.to_id=msg.id;
                this.showTitle(msg.name+","+GOFLY_LANG[LANG]['chating']);
                this.scrollBottom();
                this.showKfonline=true;
            }
            if (redata.type == "transfer") {
                var kefuId = redata.data
                if(!kefuId){
                    return;
                }
                this.visitor.to_id=kefuId;
            }
            if (redata.type == "notice") {
                let msg = redata.data
                if(!msg){
                    return;
                }
                this.chatTitle=msg
                $(".chatBox").append("<div class=\"chatTime\">"+this.chatTitle+"</div>");
                this.scrollBottom();
            }
            if (redata.type == "message") {
                let msg = redata.data
                this.visitor.to_id=msg.id;

                let content = {}
                content.avator = msg.avator;
                content.name = msg.name;
                content.content =replaceContent(msg.content);
                content.is_kefu = true;
                content.time = msg.time;
                this.msgList.push(content);

                notify(msg.name, {
                    body: msg.content,
                    icon: msg.avator
                },function(notification) {
                    window.focus();
                    notification.close();
                });
                this.scrollBottom();
                flashTitle();//标题闪烁
                clearInterval(this.timer);
                this.alertSound();//提示音
            }
            if (redata.type == "close") {
                this.chatTitle=GOFLY_LANG[LANG]['closemes'];
                $(".chatBox").append("<div class=\"chatTime\">"+this.chatTitle+"</div>");
                this.scrollBottom();
                this.socket.close();
                //this.socketClosed=true;
                this.focusSendConn=true;
            }
            if (redata.type == "force_close") {
                this.chatTitle=GOFLY_LANG[LANG]['forceclosemes'];
                $(".chatBox").append("<div class=\"chatTime\">"+this.chatTitle+"</div>");
                this.scrollBottom();
                this.socket.close();
                this.socketClosed=true;
            }
            if (redata.type == "auto_close") {
                this.chatTitle=GOFLY_LANG[LANG]['autoclosemes'];
                $(".chatBox").append("<div class=\"chatTime\">"+this.chatTitle+"</div>");
                this.scrollBottom();
                this.socket.close();
                this.socketClosed=true;
            }
            window.parent.postMessage(redata,"*");
        },
        //发送给客户
        chatToUser:function() {
            var messageContent=this.messageContent.trim("\r\n");
            messageContent=messageContent.replace("\n","");
            messageContent=messageContent.replace("\r\n","");
            if(messageContent==""||messageContent=="\r\n"){
                this.messageContent="";
                return;
            }
            this.messageContent=messageContent;
            if(this.socketClosed){
                this.$message({
                    message: '连接关闭!请重新打开页面',
                    type: 'warning'
                });
                return;
            }
            this.sendDisabled=true;
            let _this=this;

            let content = {}
            content.avator=_this.visitor.avator;
            content.content = replaceContent(_this.messageContent);
            content.name = _this.visitor.name;
            content.is_kefu = false;
            content.time = _this.getNowDate();
            content.show_time=false;
            _this.msgList.push(content);
            _this.scrollBottom();

            let mes = {};
            mes.type = "visitor";
            mes.content = this.messageContent;
            mes.from_id = this.visitor.visitor_id;
            mes.to_id = this.visitor.to_id;
            mes.content = this.messageContent;
            //发送消息
            $.post("/2/message",mes,function(res){
                _this.sendDisabled=false;
                if(res.code!=200){
                    _this.msgList.pop();
                    _this.$message({
                        message: res.msg,
                        type: 'error'
                    });
                    return;
                }
                _this.messageContent = "";
                clearInterval(_this.timer);
                _this.sendSound();
            });

        },
        OnClose:function() {
            console.log("ws:onclose");
            this.focusSendConn=true;
            //this.socketClosed=true;
            // this.chatTitle="连接关闭!请重新打开页面";
            // $(".chatBox").append("<div class=\"chatTime\">"+this.chatTitle+"</div>");
            // this.scrollBottom();
        },
        //获取当前用户信息
        getUserInfo:function(){
            let obj=this.getCache("visitor");
            var visitor_id=""
            var to_id=KEFU_ID;
            if(obj){
                visitor_id=obj.visitor_id;
                //to_id=obj.to_id;
            }
                let _this=this;
                var extra=getQuery("extra");
                //发送消息
                $.post("/visitor_login",{visitor_id:visitor_id,refer:REFER,to_id:to_id,extra:extra},function(res){
                    if(res.code!=200){
                        _this.$message({
                            message: res.msg,
                            type: 'error'
                        });
                        _this.sendDisabled=true;
                        return;
                    }
                    _this.visitor=res.result;
                    _this.getHistoryMessage();
                    _this.setCache("visitor",res.result);
                    //_this.getMesssagesByVisitorId();
                    _this.initConn();
                });
            // }else{
            //     this.visitor=obj;
            //     this.initConn();
            // }
        },
        //获取信息列表
        getHistoryMessage:function(){
            let params={
                page:this.messages.page,
                pagesize: this.messages.pagesize,
                visitor_id: this.visitor.visitor_id,
            }
            let _this=this;
            $.get("/2/messagesPages",params,function(res){
                let msgList=res.result.list;
                if(msgList.length>=_this.messages.pagesize){
                    _this.showLoadMore=true;
                }else{
                    _this.showLoadMore=false;
                }
                for(let i in msgList){
                    let item = msgList[i];
                    let content = {}
                    if (item["mes_type"] == "kefu") {
                        item.is_kefu = true;
                        item.avator=item["kefu_avator"];

                    } else {
                        item.is_kefu = false;
                        item.avator=item["visitor_avator"];
                    }
                    item.time = item["create_time"];
                    item.content=replaceContent(item["content"]);
                    _this.msgList.unshift(item);
                }
                if(_this.messages.page==1){
                    _this.scrollBottom();
                }
                _this.messages.page++;
            });
        },
        //滚动到底部
        scrollBottom:function(){
            var _this=this;
            this.$nextTick(function(){
                $('.chatVisitorPage').scrollTop($(".chatVisitorPage")[0].scrollHeight);
            });
        },
        //获取日期
        getNowDate : function() {// 获取日期
            var d = new Date(new Date());
            return d.getFullYear() + '-' + this.digit(d.getMonth() + 1) + '-' + this.digit(d.getDate())
                + ' ' + this.digit(d.getHours()) + ':' + this.digit(d.getMinutes()) + ':' + this.digit(d.getSeconds());
        },
        //补齐数位
        digit : function (num) {
            return num < 10 ? '0' + (num | 0) : num;
        },
        setCache : function (key,obj){
            if(navigator.cookieEnabled&&typeof window.localStorage !== 'undefined'){
                localStorage.setItem(key, JSON.stringify(obj));
            }
        },getCache : function (key){
            if(navigator.cookieEnabled&&typeof window.localStorage !== 'undefined') {
                return JSON.parse(localStorage.getItem(key));
            }
        },
        //获取自动欢迎语句
        getNotice : function (){
            let _this=this;
            $.get("/notice?kefu_id="+KEFU_ID,function(res) {
                var code=res.code;
                if(code!=200) return;
                _this.kefuInfo=res.result;
                _this.showTitle(_this.kefuInfo.nickname+" 为您服务");
                if(!_this.kefuInfo.welcome) return;
                var msg={
                    content:replaceContent(_this.kefuInfo.welcome),
                    avator:_this.kefuInfo.avatar,
                    name :_this.kefuInfo.nickname,
                    time:_this.getNowDate(),
                    is_kefu:true,
                }
                _this.msgList.push(msg);
                _this.scrollBottom();
                _this.alertSound();
            });
        },
        initCss:function(){
            var _this=this;
            $(function () {
                //$(".chatContext").css("max-height",$(window).height());
                // if (top.location != location){
                //     $(".chatContext").css("max-height",$(window).height()-65);
                // }
                //展示表情
                var faces=placeFace();
                $.each(faceTitles, function (index, item) {
                    _this.face.push({"name":item,"path":faces[item]});
                });


                var windheight = $(window).height();
                $(window).resize(function(){
                    var docheight = $(window).height();  /*唤起键盘时当前窗口高度*/
                    //_this.scrollBottom();
                    $('body').scrollTop(99999999);
                    // if(docheight < windheight){            /*当唤起键盘高度小于未唤起键盘高度时执行*/
                    //     $(".chatBoxSend").css("position","static");
                    // }else{
                    //     $(".chatBoxSend").css("position","fixed");
                    // }
                });
            });
        },
        //心跳
        ping:function(){
            let _this=this;
            let mes = {}
            mes.type = "ping";
            mes.data = "visitor:"+_this.visitor.visitor_id;
            setInterval(function () {
                if(_this.socket!=null){
                    _this.socket.send(JSON.stringify(mes));
                }
            },10000);
        },
        //初始化
        init:function(){
            var _this=this;
            this.initCss();
            $('body').click(function(){
                clearFlashTitle();
                window.parent.postMessage({type:"focus"},"*");
                $('.faceBox').hide();
            });
            window.onfocus = function () {
                //_this.scrollBottom();
                clearFlashTitle();
                window.parent.postMessage({type:"focus"},"*");
                if(_this.socketClosed){
                    return;
                }
                if(!_this.focusSendConn){
                    return;
                }
                _this.initConn();
                _this.scrollBottom();
            }

        },
        //表情点击事件
        faceIconClick:function(index){
            $('.faceBox').hide();
            this.messageContent+="face"+this.face[index].name;
        },
        //上传图片
        uploadImg:function (url){
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
                            var data=JSON.stringify({
                                name:res.result.name,
                                ext:res.result.ext,
                                size:res.result.size,
                                path:'/' + res.result.path,
                            })
                            _this.messageContent+='attachment['+data+']';
                            _this.chatToUser();
                        }
                    },
                    error: function (data) {
                        console.log(data);
                    }
                });
            });
        },
        //粘贴上传图片
        onPasteUpload:function(event){
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
        //提示音
        alertSound:function(){
            alertSound("chatMessageAudio",'/static/images/alert2.ogg');
        },
        sendSound:function(){
            alertSound("chatMessageSendAudio",'/static/images/sent.ogg');
        },
        sendAjax:function(url,method,params,callback){
            let _this=this;
            $.ajax({
                type: method,
                url: url,
                data:params,
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
        showTitle:function(title){
            $(".chatBox").append("<div class='chatNotice'><div class=\"chatNoticeContent\"><span>"+title+"</span></div></div>");
            this.scrollBottom();
        },
    },
    mounted:function() {
        document.addEventListener('paste', this.onPasteUpload)
        document.addEventListener('scroll',this.textareaBlur)
    },
    created: function () {
        this.init();
        this.getUserInfo();
        //加载历史记录
        //this.msgList=this.getHistory();
        //滚动底部
        //this.scrollBottom();

    }
})
