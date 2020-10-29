KEFU_ID=KEFU_ID!=""? KEFU_ID:"kefu2";
new Vue({
    el: '#app',
    delimiters:["<{","}>"],
    data: {
        window:window,
        //server:getWsBaseUrl()+"/chat_server",
        server:getWsBaseUrl()+"/ws_visitor",
        socket:null,
        msgList:[],
        messageContent:"",
        chatTitle:"正在连接...",
        visitor:{},
        face:[],
        showKfonline:false,
        socketClosed:false,
        timer:null,
    },
    methods: {
        //初始化websocket
        initConn() {
            let socket = new ReconnectingWebSocket(this.server+"?visitor_id="+this.visitor.visitor_id);//创建Socket实例
            socket.maxReconnectAttempts = 30;
            this.socket = socket
            this.socket.onmessage = this.OnMessage;
            this.socket.onopen = this.OnOpen;
            this.socket.onclose = this.OnClose;
            //心跳
            this.ping();
        },
        OnOpen() {
            this.chatTitle="连接成功!"
            // let mes = {}
            // mes.type = "userInit";
            // this.visitor.refer=REFER;
            // mes.data = this.visitor;
            // this.socket.send(JSON.stringify(mes));
        },
        OnMessage(e) {
            const redata = JSON.parse(e.data);
            if (redata.type == "kfOnline") {
                let msg = redata.data
                if(this.showKfonline && this.visitor.to_id==msg.id){
                    return;
                }
                this.visitor.to_id=msg.id;
                this.chatTitle=msg.name+",正在与您沟通!"
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
                content.content =replaceContent(msg.content,true);
                content.is_kefu = false;
                content.time = msg.time;
                this.msgList.push(content);

                //this.saveHistory(content);
                this.scrollBottom();
                flashTitle();//标题闪烁
                clearInterval(this.timer);
                this.alertSound();//提示音
            }
            if (redata.type == "close") {
                this.chatTitle="连接关闭!请重新打开页面";
                $(".chatBox").append("<div class=\"chatTime\">"+this.chatTitle+"</div>");
                this.scrollBottom();
                this.socket.close();
                this.socketClosed=true;
            }
            window.parent.postMessage(redata);
        },
        //发送给客户
        chatToUser() {
            this.messageContent=this.messageContent.trim("\r\n");
            if(this.messageContent==""||this.messageContent=="\r\n"){
                this.$message({
                    message: '不能发送空白信息',
                    type: 'warning'
                });
                return;
            }
            if(this.socketClosed){
                this.$message({
                    message: '连接关闭!请重新打开页面',
                    type: 'warning'
                });
                return;
            }
            let _this=this;
            let mes = {};
            mes.type = "visitor";
            mes.content = this.messageContent;
            mes.from_id = this.visitor.visitor_id;
            mes.to_id = this.visitor.to_id;
            mes.content = this.messageContent;
            //发送消息
            $.post("/2/message",mes,function(res){
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
        OnClose() {
            this.chatTitle="连接关闭!请重新打开页面";
            $(".chatBox").append("<div class=\"chatTime\">"+this.chatTitle+"</div>");
        },
        //获取当前用户信息
        getUserInfo(){
            let obj=this.getCache("visitor");
            var visitor_id=""
            if(obj){
                visitor_id=obj.visitor_id;
            }
                let _this=this;
                //发送消息
                $.post("/visitor_login",{visitor_id:visitor_id,refer:REFER,to_id:KEFU_ID,client_ip:returnCitySN["cip"],},function(res){
                    if(res.code!=200){
                        _this.$message({
                            message: res.msg,
                            type: 'error'
                        });
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
        getMesssagesByVisitorId(){
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
                        _this.$nextTick(() => {
                            $(".chatBox").append("<div class=\"chatTime\">—— 以上是历史消息 ——</div>");
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
            this.$nextTick(() => {
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
                if (res.result != null) {
                    let msg = res.result;
                    var len=msg.length;
                    var i=0;
                    if(len>0){
                        _this.timer=setInterval(function(){
                            if(i>=len){
                                clearInterval(_this.timer);
                            }
                            let content = msg[i];
                            content.content = replaceContent(content.content);
                            _this.msgList.push(content);
                            _this.scrollBottom();
                            _this.sendSound();
                            i++;
                        },4000);
                    }

                }
            });
        },
        initCss(){
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
            });
        },
        //心跳
        ping(){
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
        init(){
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
        }
    },
    mounted() {
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
