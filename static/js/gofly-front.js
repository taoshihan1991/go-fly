var GOFLY={
    GOFLY_URL:"https://gofly.sopans.com",
    GOFLY_KEFU_ID:"",
    GOFLY_BTN_TEXT:"Chat with me",
    GOFLY_LANG:"en",
    GOFLY_EXTRA:"",
    GOFLY_AUTO_OPEN:false,
};
GOFLY.launchButtonFlag=false;
GOFLY.titleTimer=0;
GOFLY.titleNum=0;
GOFLY.noticeTimer=null;
GOFLY.originTitle=document.title;
GOFLY.chatPageTitle="GOFLY";
GOFLY.init=function(config){
    var _this=this;
    if(typeof config=="undefined"){
        return;
    }

    if (typeof config.GOFLY_URL!="undefined"){
        this.GOFLY_URL=config.GOFLY_URL;
    }
    this.dynamicLoadCss(this.GOFLY_URL+"/static/css/gofly-front.css?v=1");

    if (typeof config.GOFLY_KEFU_ID!="undefined"){
        this.GOFLY_KEFU_ID=config.GOFLY_KEFU_ID;
    }
    if (typeof config.GOFLY_BTN_TEXT!="undefined"){
        this.GOFLY_BTN_TEXT=config.GOFLY_BTN_TEXT;
    }
    if (typeof config.GOFLY_EXTRA!="undefined"){
        this.GOFLY_EXTRA=config.GOFLY_EXTRA;
    }
    if (typeof config.GOFLY_AUTO_OPEN!="undefined"){
        this.GOFLY_AUTO_OPEN=config.GOFLY_AUTO_OPEN;
    }
    if(this.GOFLY_EXTRA==""){
        var refer=document.referrer?document.referrer:"无";
        this.GOFLY_EXTRA='{"refer":"'+refer+'","host":"'+document.location.href+'"}';
    }
    this.dynamicLoadJs(this.GOFLY_URL+"/static/js/functions.js?v=1",function(){
        if (typeof config.GOFLY_LANG!="undefined"){
            _this.GOFLY_LANG=config.GOFLY_LANG;
        }else{
            _this.GOFLY_LANG=checkLang();
        }
        _this.GOFLY_EXTRA=utf8ToB64(_this.GOFLY_EXTRA);
    });

    if (typeof $!="function"){
        this.dynamicLoadJs("https://cdn.jsdelivr.net/npm/jquery/dist/jquery.min.js",function () {
            _this.dynamicLoadJs("https://cdn.bootcdn.net/ajax/libs/layer/3.1.1/layer.min.js",function () {
                _this.clickBtn();
            });
        });
    }else{
        this.dynamicLoadJs("https://cdn.bootcdn.net/ajax/libs/layer/3.1.1/layer.min.js",function () {
            _this.clickBtn();
        });
    }

    window.addEventListener('message',function(e){
        var msg=e.data;
        if(msg.type=="message"){
            _this.flashTitle();//标题闪烁
            $("#launchNoticeContent").html(replaceContent(msg.data.content,_this.GOFLY_URL));
            $("#launchButtonNotice").show();
        }
    });
    window.onfocus = function () {
        clearTimeout(this.titleTimer);
        console.log(1);
        document.title = _this.originTitle;
    };
}
GOFLY.dynamicLoadCss=function(url){
    var head = document.getElementsByTagName('head')[0];
    var link = document.createElement('link');
    link.type='text/css';
    link.rel = 'stylesheet';
    link.href = url;
    head.appendChild(link);
}
GOFLY.dynamicLoadJs=function(url, callback){
    var head = document.getElementsByTagName('head')[0];
    var script = document.createElement('script');
    script.type = 'text/javascript';
    script.src = url;
    if(typeof(callback)=='function'){
        script.onload = script.onreadystatechange = function () {
            if (!this.readyState || this.readyState === "loaded" || this.readyState === "complete"){
                callback();
                script.onload = script.onreadystatechange = null;
            }
        };
    }
    head.appendChild(script);
}

GOFLY.clickBtn=function (){
    var _this=this;
    var html="<div class='launchButtonBox'>" +
        '<div id="launchButton" class="launchButton">' +
        '<div id="launchIcon" class="launchIcon animateUpDown">1</div> ' +
        '<div class="launchButtonText">'+_this.GOFLY_BTN_TEXT+'</div></div>' +
        '<div id="launchButtonNotice" class="launchButtonNotice">您好:<br/>极简强大的开源免费Go语言在线客服单页营销系统，来了解一下？</div>' +
        '</div>';
    $('body').append(html);
    $(".launchButton").on("click",function() {
        _this.showKefu();
    });
    if(this.GOFLY_AUTO_OPEN){
        _this.showKefu();
        $(".launchButtonBox").show();
        this.launchButtonFlag=false;
    }
    $("body").on("click","#launchNoticeClose",function() {
        $("#launchButtonNotice").hide();
    });
    _this.getNotice();
}
GOFLY.getNotice=function(){
    var _this=this;
    $.get(this.GOFLY_URL+"/notice?kefu_id="+this.GOFLY_KEFU_ID,function(res) {
        _this.chatPageTitle="<img style='margin-top: 5px;' src='"+_this.GOFLY_URL+res.result.avatar+"' class='flyAvatar'>"+res.result.username;
        if (res.result.welcome != null) {
            var msg = res.result.welcome;
            var len=msg.length;
            var i=0;
            if(len>0){

                _this.noticeTimer=setInterval(function(){
                    if(i>=len||typeof msg[i]=="undefined"||msg[i]==null){
                        clearInterval(_this.noticeTimer);
                        return;
                    }
                    var content = msg[i];
                    if(typeof content.content =="undefined"){
                        return;
                    }
                    var welcomeHtml="<div class='flyUser'><img class='flyAvatar' src='"+_this.GOFLY_URL+res.result.avatar+"'/> <span class='flyUsername'>"+res.result.username+"</span>" +
                        "<span id='launchNoticeClose' class='flyClose'>×</span>" +
                        "</div>";
                    welcomeHtml+="<div id='launchNoticeContent'>"+replaceContent(content.content,_this.GOFLY_URL)+"</div>";
                    $("#launchButtonNotice").html(welcomeHtml).show();
                    i++;
                    $("#launchIcon").text(i).show();
                },4000);
            }

        }
    });
}
GOFLY.isIE=function(){
    var userAgent = navigator.userAgent; //取得浏览器的userAgent字符串
    var isIE = userAgent.indexOf("compatible") > -1 && userAgent.indexOf("MSIE") > -1; //判断是否IE<11浏览器
    var isEdge = userAgent.indexOf("Edge") > -1 && !isIE; //判断是否IE的Edge浏览器
    var isIE11 = userAgent.indexOf('Trident') > -1 && userAgent.indexOf("rv:11.0") > -1;
    if(isIE) {
        var reIE = new RegExp("MSIE (\\d+\\.\\d+);");
        reIE.test(userAgent);
        var fIEVersion = parseFloat(RegExp["$1"]);
        if(fIEVersion == 7) {
            return 7;
        } else if(fIEVersion == 8) {
            return 8;
        } else if(fIEVersion == 9) {
            return 9;
        } else if(fIEVersion == 10) {
            return 10;
        } else {
            return 6;//IE版本<=7
        }
    } else if(isEdge) {
        return 'edge';//edge
    } else if(isIE11) {
        return 11; //IE11
    }else{
        return -1;//不是ie浏览器
    }
}
GOFLY.showKefu=function (){
    if (this.launchButtonFlag) return;
    var width=$(window).width();
    if(this.isIE()>0){
        this.windowOpen();
        return;
    }
    if(width<768){
        this.layerOpen("100%","100%");
        return;
    }
    this.layerOpen("400px","530px");
    this.launchButtonFlag=true;
    $(".launchButtonBox").hide();
    var _this=this;
    $("body").click(function () {
        clearTimeout(_this.titleTimer);
        document.title = _this.originTitle;
    });
    window.onfocus = function () {
        clearTimeout(_this.titleTimer);
        document.title = _this.originTitle;
    };
}
GOFLY.layerOpen=function (width,height){
    if (this.launchButtonFlag) return;
    if($("#layui-layer1").css("display")=="none"){
        $("#layui-layer1").show();
        return;
    }
    var _this=this;
    layer.open({
        type: 2,
        title: this.chatPageTitle,
        closeBtn: 1, //不显示关闭按钮
        shade: 0,
        area: [width, height],
        offset: 'rb', //右下角弹出
        anim: 2,
        content: [this.GOFLY_URL+'/chatIndex?kefu_id='+this.GOFLY_KEFU_ID+'&lang='+this.GOFLY_LANG+'&refer='+window.document.title+'&extra='+this.GOFLY_EXTRA , 'yes'], //iframe的url，no代表不显示滚动条
        end: function(){
            _this.launchButtonFlag=false;
            $(".launchButtonBox").show();
        },
        cancel: function(index, layero){
            $("#layui-layer1").hide();
            _this.launchButtonFlag=false;
            $(".launchButtonBox").show();
            return false;
        }
    });
}
GOFLY.windowOpen=function (){
   window.open(this.GOFLY_URL+'/chatIndex?kefu_id='+this.GOFLY_KEFU_ID+'&lang='+this.GOFLY_LANG+'&refer='+window.document.title+'&extra='+this.GOFLY_EXTRA);
}
GOFLY.flashTitle=function () {
    this.titleNum++;
    if (this.titleNum >=3) {
        this.titleNum = 1;
    }
    if (this.titleNum == 1) {
        document.title = '【】' + this.originTitle;
    }
    if (this.titleNum == 2) {
        document.title = '【new message】' + this.originTitle;
    }
    this.titleTimer = setTimeout("GOFLY.flashTitle()", 500);
}


