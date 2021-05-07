var GOFLY={
    GOFLY_URL:"https://gofly.sopans.com",
    GOFLY_KEFU_ID:"",
    GOFLY_BTN_TEXT:"Chat with me",
    GOFLY_LANG:"en",
    GOFLY_EXTRA: {},
    GOFLY_AUTO_OPEN:true,
    GOFLY_AUTO_SHOW:false,
    GOFLY_WITHOUT_BTN:false,
};
GOFLY.launchButtonFlag=false;
GOFLY.titleTimer=0;
GOFLY.titleNum=0;
GOFLY.noticeTimer=null;
GOFLY.originTitle=document.title;
GOFLY.chatPageTitle="GOFLY";
GOFLY.kefuName="";
GOFLY.kefuAvator="";
GOFLY.init=function(config){
    var _this=this;
    if(typeof config=="undefined"){
        return;
    }

    if (typeof config.GOFLY_URL!="undefined"){
        this.GOFLY_URL=config.GOFLY_URL.replace(/([\w\W]+)\/$/,"$1");
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
    if (typeof config.GOFLY_AUTO_SHOW!="undefined"){
        this.GOFLY_AUTO_SHOW=config.GOFLY_AUTO_SHOW;
    }
    if (typeof config.GOFLY_WITHOUT_BTN!="undefined"){
        this.GOFLY_WITHOUT_BTN=config.GOFLY_WITHOUT_BTN;
    }
    var refer=document.referrer?document.referrer:"无";
    this.GOFLY_EXTRA.refer=refer;
    this.GOFLY_EXTRA.host=document.location.href;
    this.GOFLY_EXTRA=JSON.stringify(_this.GOFLY_EXTRA);

    this.dynamicLoadJs(this.GOFLY_URL+"/assets/js/functions.js?v=1",function(){
        if (typeof config.GOFLY_LANG!="undefined"){
            _this.GOFLY_LANG=config.GOFLY_LANG;
        }else{
            _this.GOFLY_LANG=checkLang();
        }
        _this.GOFLY_EXTRA=utf8ToB64(_this.GOFLY_EXTRA);
    });
    if (typeof $!="function"){
        this.dynamicLoadJs("https://cdn.staticfile.org/jquery/3.6.0/jquery.min.js",function () {
            _this.dynamicLoadJs("https://cdn.staticfile.org/layer/3.4.0/layer.min.js",function () {
                _this.jsCallBack();
            });
        });
    }else{
        this.dynamicLoadJs("https://cdn.staticfile.org/layer/3.4.0/layer.min.js",function () {
            _this.jsCallBack();
        });
    }
    _this.addEventlisten();
}
GOFLY.jsCallBack=function(){
    this.showKefuBtn();
    this.addClickEvent();
    this.getNotice();
}
GOFLY.showKefuBtn=function(){
    var _this=this;
    if(_this.GOFLY_WITHOUT_BTN){
        return;
    }
    var html="<div class='launchButtonBox'>" +
        '<div id="launchButton" class="launchButton">' +
        '<div id="launchIcon" class="launchIcon animateUpDown">1</div> ' +
        '<div class="launchButtonText">'+_this.GOFLY_BTN_TEXT+'</div></div>' +
        '<div id="launchButtonNotice" class="launchButtonNotice"></div>' +
        '</div>';
    $('body').append(html);
}
GOFLY.addClickEvent=function(){
    var _this=this;
    $(".launchButton").on("click",function() {
        _this.GOFLY_AUTO_SHOW=true;
        _this.showKefu();
        $("#launchIcon").text(0).hide();
    });

    $("body").on("click","#launchNoticeClose",function() {
        $("#launchButtonNotice").hide();
    });

    $("body").click(function () {
        clearTimeout(_this.titleTimer);
        document.title = _this.originTitle;
    });
}
GOFLY.addEventlisten=function(){
    var _this=this;
    window.addEventListener('message',function(e){
        var msg=e.data;
        if(msg.type=="message"){
            clearInterval(_this.noticeTimer);
            var width=$(window).width();
            if(width>768){
                _this.flashTitle();//标题闪烁
            }
            if (_this.launchButtonFlag){
                return;
            }
            var welcomeHtml="<div class='flyUser'><img class='flyAvatar' src='"+_this.GOFLY_URL+msg.data.avator+"'/> <span class='flyUsername'>"+msg.data.name+"</span>" +
                "<span id='launchNoticeClose' class='flyClose'>×</span>" +
                "</div>";
            welcomeHtml+="<div id='launchNoticeContent'>"+replaceContent(msg.data.content,_this.GOFLY_URL)+"</div>";
            $("#launchButtonNotice").html(welcomeHtml).show();
            var news=$("#launchIcon").text();
            $("#launchIcon").text(++news).show();
        }
        if(msg.type=="focus"){
            clearTimeout(_this.titleTimer);
            _this.titleTimer=0;
            document.title = _this.originTitle;
        }
    });
    window.onfocus = function () {
        clearTimeout(_this.titleTimer);
        _this.titleTimer=0;
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

GOFLY.getNotice=function(){
    var _this=this;
    $.get(this.GOFLY_URL+"/notice?kefu_id="+this.GOFLY_KEFU_ID,function(res) {
        if(res.result.status=='offline'){
            _this.chatPageTitle="<div class='launchPointer offline'></div>";
        }else{
            _this.chatPageTitle="<div class='launchPointer'></div>";
            setTimeout(function(){
                var userInfo="<img style='margin-top: 5px;' class='flyAvatar' src='"+_this.GOFLY_URL+res.result.avatar+"'/> <span class='flyUsername'>"+res.result.username+"</span>"
                $('.launchButtonText').html(userInfo);
            },3000);
        }
        _this.kefuAvator=res.result.avatar;
        _this.kefuName=res.result.username;
        _this.chatPageTitle+="<img src='"+_this.GOFLY_URL+res.result.avatar+"' class='flyAvatar'>"+res.result.username;
        if(_this.GOFLY_AUTO_OPEN&&_this.isIE()<=0){
            _this.showKefu();
            $(".launchButtonBox").show();
            _this.launchButtonFlag=false;
        }
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

                    var obj=$("#launchButtonNotice");
                    if(obj[0]){
                        obj[0].innerHTML=welcomeHtml;
                        obj.show();
                    }
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
GOFLY.showPanel=function (){
    var width=$(window).width();
    this.GOFLY_AUTO_SHOW=true;
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
}
GOFLY.layerOpen=function (width,height){
    if (this.launchButtonFlag) return;
    var layBox=$("#layui-layer19911116");
    if(layBox.css("display")=="none"){
        layBox.show();
        return;
    }
    var _this=this;
    layer.index="19911115";
    layer.open({
        type: 2,
        title: this.chatPageTitle,
        closeBtn: 1, //不显示关闭按钮
        shade: 0,
        area: [width, height],
        offset: 'rb', //右下角弹出
        anim: 2,
        content: [this.GOFLY_URL+'/chatIndex?kefu_id='+this.GOFLY_KEFU_ID+'&lang='+this.GOFLY_LANG+'&refer='+window.document.title+'&extra='+this.GOFLY_EXTRA , 'yes'], //iframe的url，no代表不显示滚动条
        success:function(){
            var layBox=$("#layui-layer19911116");
            if(_this.GOFLY_AUTO_SHOW&&layBox.css("display")=="none"){
                _this.launchButtonFlag=true;
                layBox.show();
            }
        },
        end: function(){
            _this.launchButtonFlag=false;
            $(".launchButtonBox").show();
        },
        cancel: function(index, layero){
            $("#layui-layer19911116").hide();
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
    if(this.titleTimer!=0){
        return;
    }
    this.titleTimer = setInterval("GOFLY.flashTitleFunc()", 500);
}
GOFLY.flashTitleFunc=function(){
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
}


