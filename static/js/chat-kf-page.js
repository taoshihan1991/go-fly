//首页组件
var chatKfIndex = {
    delimiters:["<{","}>"],
    data: function(){
        return {
            visitors: [],
        }
    },
    methods: {
        init(){
            this.getKefuInfo();
        },
        kfOnline() {
            let messsage = {}
            messsage.type = "kfOnline";
            messsage.data = this.$parent.kefuInfo;
            this.$parent.socket.send(JSON.stringify(messsage));
        },
        receiveMessage(e) {
            const retData = JSON.parse(e.data);
            switch (retData.type) {
                case "allUsers":
                    this.visitors = retData.data;
                    break;
                case "userOnline":
                    this.visitors.push(retData.data);
                    break;
            }
        },
        //初始化websocket
        initConn() {
            if(this.$parent.socket==null){
                this.$parent.socket = new ReconnectingWebSocket(this.$parent.server);
            }
            this.$parent.socket.onopen=this.kfOnline;
            this.$parent.socket.onmessage = this.receiveMessage;
        },
        //获取客服信息
        getKefuInfo(){
            let _this=this;
            $.ajax({
                type:"get",
                url:"/kefuinfo",
                headers:{
                    "token":TOKEN
                },
                success: function(data) {
                    if(data.code==200 && data.result!=null){
                        _this.$parent.kefuInfo.id=data.result.id;
                        _this.$parent.kefuInfo.name=data.result.name;
                        _this.$parent.kefuInfo.avator=data.result.avator;
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
    },
    created: function () {
        this.init();
    },
    template:$("#chatKfIndex").html()
};
//详情组件
var chatKfBox = {
    delimiters:["<{","}>"],
    data: function(){
        return {
            visitorId:null,
            msgList: [],
            messageContent: "",
            face: [],
        }
    },
    methods: {
        receiveMessage(e) {
            const retData = JSON.parse(e.data);
            switch (retData.type) {
                case "message":
                    alert(e.data);
                    break;
            }
        },
        init(){
            //获取当前客户消息
            this.visitorId=this.$route.params.visitorId;
            this.getMesssagesByVisitorId(this.$route.params.visitorId);
            this.$parent.socket.onmessage = this.receiveMessage;
        },
        //获取信息列表
        getMesssagesByVisitorId(visitorId) {
            let _this = this;
            $.ajax({
                type: "get",
                url: "/messages?visitorId=" + visitorId,
                headers: {
                    "token": TOKEN
                },
                success: function (data) {
                    if (data.code == 200 && data.result != null) {
                        let msgList = data.result;
                        _this.msgList = [];
                        for (let i = 0; i < msgList.length; i++) {
                            let visitorMes = msgList[i];
                            let content = {}
                            if (visitorMes["mes_type"] == "kefu") {
                                content.is_kefu = true;
                                content.avator = visitorMes["kefu_avator"];
                                content.name = visitorMes["kefu_name"];
                            } else {
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
                    if (data.code != 200) {
                        _this.$message({
                            message: data.msg,
                            type: 'error'
                        });
                    }
                }
            });
        },
        //发送给客户
        chatToUser() {
            this.messageContent=this.messageContent.trim("\r\n");
            if(this.messageContent==""||this.messageContent=="\r\n"||this.currentGuest==""){
                return;
            }
            let _this=this;
            let mes = {};
            mes.type = "kefu";
            mes.content = this.messageContent;
            mes.from_id = _this.$parent.kefuInfo.id;
            mes.to_id = this.visitorId;
            mes.content = this.messageContent;
            $.post("/message",mes,function(){
                _this.messageContent = "";
            });

            let content = {}
            content.avator = _this.$parent.kefuInfo.avator;
            content.name = _this.$parent.kefuInfo.name;
            content.content = replaceContent(this.messageContent);
            content.is_kefu = true;
            content.time = '';
            this.msgList.push(content);
            this.scrollBottom();
        },
        //滚到底部
        scrollBottom(){
            this.$nextTick(() => {
                $('.chatBox').scrollTop($(".chatBox")[0].scrollHeight);
            });
        },
    },
    created: function () {
        this.init();
    },
    template:$("#chatBox").html()
};
var routes = [
    { path: '/',component:chatKfIndex}, // 这个表示会默认渲染
    {path:'/chatKfBox/:visitorId',component:chatKfBox},
];
var router = new VueRouter({
    routes: routes
})

new Vue({
    router,
    el: '#app',

    data: function(){
        return{
            server:getWsBaseUrl()+"/chat_server",
            socket:null,
            kefuInfo:{},
        }
    },
    methods:{

    },
    created: function () {
    },
})