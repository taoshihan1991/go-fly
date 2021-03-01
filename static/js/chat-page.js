KEFU_ID=KEFU_ID!=""? KEFU_ID:"kefu2";
new Vue({
    el: '#app',
    delimiters:["<{","}>"],
    data: {
        window:window,
        server:getWsBaseUrl()+"/ws_visitor",
        socket:null,
        msgList:[],
        messageContent:"",
        chatTitle:GOFLY_LANG[LANG]['connecting'],
        visitor:{},
        face:[],
        showKfonline:false,
        socketClosed:false,
        timer:null,
        sendDisabled:false,
        flyLang:GOFLY_LANG[LANG],
    },
    methods: {
        //初始化websocket
        initConn:function() {
            let socket = new ReconnectingWebSocket(this.server+"?visitor_id="+this.visitor.visitor_id);//创建Socket实例
            socket.maxReconnectAttempts = 30;
            this.socket = socket
            this.socket.onmessage = this.OnMessage;
            this.socket.onopen = this.OnOpen;
            this.socket.onclose = this.OnClose;
            //心跳
            this.ping();
        },
        OnOpen:function() {
            this.chatTitle=GOFLY_LANG[LANG]['connectok'];
            this.socketClosed=false;
        },
        OnMessage:function(e) {
            this.socketClosed=false;
            const redata = JSON.parse(e.data);
            if (redata.type == "kfOnline") {
                let msg = redata.data
                if(this.showKfonline && this.visitor.to_id==msg.id){
                    return;
                }
                this.visitor.to_id=msg.id;
                this.chatTitle=msg.name+","+GOFLY_LANG[LANG]['chating'];
                $(".chatBox").append("<div class=\"chatTime\">"+this.chatTitle+"</div>");
                this.scrollBottom();
                this.showKfonline=true;
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
                content.is_kefu = false;
                content.time = msg.time;
                this.msgList.push(content);

                notify(msg.name, {
                    body: msg.content,
                    icon: msg.avator
                });
                this.scrollBottom();
                flashTitle();//标题闪烁
                clearInterval(this.timer);
                this.alertSound();//提示音
            }
            if (redata.type == "close") {
                this.chatTitle="系统关闭连接!请重新打开页面";
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
                    _this.$message({
                        message: res.msg,
                        type: 'error'
                    });
                    return;
                }
                let content = {}
                content.avator=_this.visitor.avator;
                content.content = replaceContent(_this.messageContent);
                content.name = _this.visitor.name;
                content.is_kefu = true;
                content.time = res.result.data.time;
                _this.msgList.push(content);
                //_this.saveHistory(content);
                _this.scrollBottom();
                _this.messageContent = "";
                clearInterval(_this.timer);
                _this.sendSound();
            });

        },
        //正在输入
        inputNextText:function(){
            if(this.socketClosed||!this.socket){
                return;
            }
            //console.log(this.messageContent);
            var message = {}
            message.type = "inputing";
            message.data = {
                from : this.visitor.visitor_id,
                to : this.visitor.to_id,
                content:this.messageContent
            };
            this.socket.send(JSON.stringify(message));
        },
        OnClose:function() {
            this.socketClosed=true;
            // this.chatTitle="连接关闭!请重新打开页面";
            // $(".chatBox").append("<div class=\"chatTime\">"+this.chatTitle+"</div>");
            // this.scrollBottom();
        },
        //获取当前用户信息
        getUserInfo:function(){
            let obj=this.getCache("visitor");
            var visitor_id=""
            if(obj){
                visitor_id=obj.visitor_id;
            }
                let _this=this;
                var extra=getQuery("extra");
                if(extra!=""){
                    var extraJson=JSON.parse(window.atob(extra))
                    for(var key in extraJson){
                        if(extraJson[key]==""){
                            _this.$message({
                                message: "用户扩展信息错误",
                                type: 'error'
                            });
                            return;
                        }
                    }
                }
                //发送消息
                $.post("/visitor_login",{visitor_id:visitor_id,refer:REFER,to_id:KEFU_ID,extra:extra},function(res){
                    if(res.code!=200){
                        _this.$message({
                            message: res.msg,
                            type: 'error'
                        });
                        _this.sendDisabled=true;
                        return;
                    }
                    _this.visitor=res.result;
                    _this.setCache("visitor",res.result);
                    _this.getMesssagesByVisitorId();
                    _this.initConn();
                });
            // }else{
            //     this.visitor=obj;
            //     this.initConn();
            // }
        },
        //获取信息列表
        getMesssagesByVisitorId:function(){
            let _this=this;
            $.ajax({
                type:"get",
                url:"/messages?visitorId="+this.visitor.visitor_id,
                success: function(data) {
                    if(data.code==200 && data.result!=null&&data.result.length!=0){
                        let msgList=data.result;
                        _this.msgList=[];
                        for(let i=0;i<msgList.length;i++){
                            let visitorMes=msgList[i];
                            let content = {}
                            if(visitorMes["mes_type"]=="kefu"){
                                content.is_kefu = false;
                                content.avator = visitorMes["kefu_avator"];
                                content.name = visitorMes["kefu_name"];
                            }else{
                                content.is_kefu = true;
                                content.avator = visitorMes["visitor_avator"];
                                content.name = visitorMes["visitor_name"];
                            }
                            content.content = replaceContent(visitorMes["content"]);
                            content.time = visitorMes["time"];
                            _this.msgList.push(content);
                            _this.scrollBottom();
                        }
                        _this.$nextTick(function(){
                            $(".chatBox").append("<div class=\"chatTime\">"+GOFLY_LANG[LANG]['historymes']+"</div>");
                        });
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
        //滚动到底部
        scrollBottom:function(){
            this.$nextTick(function(){
                //debugger;
                $('body').scrollTop($("body")[0].scrollHeight);
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
            if(typeof(Storage) !== "undefined"){
                localStorage.setItem(key, JSON.stringify(obj));
            }
        },getCache : function (key){
            return JSON.parse(localStorage.getItem(key));
        },
        //获取自动欢迎语句
        getNotice : function (){
            let _this=this;
            $.get("/notice?kefu_id="+KEFU_ID,function(res) {
                //debugger;
                if (res.result.welcome != null) {
                    let msg = res.result.welcome;
                    var len=msg.length;
                    var i=0;
                    if(len>0){
                        _this.timer=setInterval(function(){
                            if(i>=len||typeof msg[i]=="undefined"||msg[i]==null){
                                clearInterval(_this.timer);
                                return;
                            }
                            let content = msg[i];
                            if(typeof content.content =="undefined"){
                                return;
                            }
                            content.content = replaceContent(content.content);
                            _this.msgList.push(content);
                            if(_this.msgList.length>=4){
                                _this.scrollBottom();
                            }
                            if(i==0){
                                _this.alertSound();
                            }
                            i++;
                        },4000);
                    }

                }
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
                $(".faceBtn").click(function(e){
                    var status=$('.faceBox').css("display");
                    if(status=="block"){
                        $('.faceBox').hide();
                    }else{
                        $('.faceBox').show();
                    }
                    return false;
                });

                var windheight = $(window).height();
                $(window).resize(function(){
                    var docheight = $(window).height();  /*唤起键盘时当前窗口高度*/
                    console.log(docheight,windheight);
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
            },60000);
        },
        //初始化
        init:function(){
            var _this=this;
            this.initCss();
            $("#app").click(function(){
                clearTimeout(titleTimer);
                document.title = originTitle;
            });
            $('body').click(function(){
                clearTimeout(titleTimer);
                document.title = originTitle;

                $('.faceBox').hide();
            });
            window.onfocus = function () {
                _this.scrollBottom();
                if(!_this.socketClosed){
                    return;
                }
                _this.initConn();
                _this.chatTitle="连接已重连";
                $(".chatBox").append("<div class=\"chatTime\">"+_this.chatTitle+"</div>");
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
            var b = document.getElementById("chatMessageAudio");
            if (b.canPlayType('audio/ogg; codecs="vorbis"')) {
                b.type= 'audio/mpeg';
                b.src= '/static/images/alert2.ogg';
                var p = b.play();
                p && p.then(function () {
                }).catch(function (e) {
                });
            }
        },
        sendSound:function(){
            var b = document.getElementById("chatMessageSendAudio");
            if (b.canPlayType('audio/ogg; codecs="vorbis"')) {
                b.type= 'audio/mpeg';
                b.src= '/static/images/sent.ogg';
                var p = b.play();
                p && p.then(function(){}).catch(function(e){});
            }
        }
    },
    mounted:function() {
        document.addEventListener('paste', this.onPasteUpload)
    },
    created: function () {
        this.init();
        this.getUserInfo();
        //加载历史记录
        //this.msgList=this.getHistory();
        //滚动底部
        //this.scrollBottom();
        //获取欢迎
        this.getNotice();

    }
})
