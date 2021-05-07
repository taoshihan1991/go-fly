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
        textareaFocused:false,
        replys:[],
        noticeName:"",
        noticeAvatar:"",
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
            console.log("ws:onopen");
            this.chatTitle=GOFLY_LANG[LANG]['connectok'];
            this.showTitle(this.chatTitle);

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
                this.chatTitle=msg.name+","+GOFLY_LANG[LANG]['chating'];
                $(".chatBox").append("<div class=\"chatTime\">"+this.chatTitle+"</div>");
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
                content.is_kefu = false;
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
            content.is_kefu = true;
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
        getMesssagesByVisitorId:function(isAll){
            let _this=this;
            $.ajax({
                type:"get",
                url:"/messages?visitorId="+this.visitor.visitor_id,
                success: function(data) {
                    if(data.code==200 && data.result!=null&&data.result.length!=0){
                        _this.msgListNum=data.result.length;
                        let msgList=data.result;
                        _this.msgList=[];
                        if(!isAll&&msgList.length>1){
                            var i=msgList.length-1
                        }else{
                            _this.msgListNum=0;
                            var i=0;
                        }
                        for(;i<msgList.length;i++){
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
                    }
                    if(data.code!=200){
                        _this.$message({
                            message: data.msg,
                            type: 'error'
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
            var _this=this;
            this.$nextTick(function(){
                $('.chatVisitorPage').scrollTop($(".chatVisitorPage")[0].scrollHeight);
            });
        },
        //软键盘问题
        textareaFocus:function(){
            this.scrollBottom()
            if(/Android|webOS|iPhone|iPad|BlackBerry/i.test(navigator.userAgent)) {
                $(".chatContext").css("margin-bottom","0");
                $(".chatBoxSend").css("position","static");
                this.textareaFocused=true;
            }
        },
        textareaBlur:function(){
            if(this.textareaFocused&&/Android|webOS|iPhone|iPad|BlackBerry/i.test(navigator.userAgent)) {
                var chatBoxSendObj=$(".chatBoxSend");
                var chatContextObj=$(".chatContext");
                if(this.textareaFocused&&chatBoxSendObj.css("position")!="fixed"){
                    chatContextObj.css("margin-bottom","105px");
                    chatBoxSendObj.css("position","fixed");
                    this.textareaFocused=false;
                }

            }
        },
        sendReply:function(title){
            this.messageContent=title;
            this.chatToUser();
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
                //debugger;
                _this.noticeName=res.result.username;
                _this.noticeAvatar=res.result.avatar;
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
                            _this.scrollBottom();
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
                $(".visitorFaceBtn").click(function(e){
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
            var _hmt = _hmt || [];
            (function() {
                var hm = document.createElement("script");
                hm.src = "https://hm.baidu.com/hm.js?82938760e00806c6c57adee91f39aa5e";
                var s = document.getElementsByTagName("script")[0];
                s.parentNode.insertBefore(hm, s);
            })();
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
        //自动
        getAutoReply:function(){
            var _this=this;
            $.get("/autoreply?kefu_id="+KEFU_ID,function(res) {
                if(res.code==200){
                    _this.replys=res.result;
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
            $(".chatBox").append("<div class=\"chatTime\"><span>"+title+"</span></div>");
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
        //获取欢迎
        this.getNotice();
        this.getAutoReply();
    }
})
