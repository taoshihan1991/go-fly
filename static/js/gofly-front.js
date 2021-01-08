var launchButtonFlag=false;
var titleTimer,titleNum=0;
var originTitle = document.title;
if (typeof GOFLY_URL=="undefined"){
    var GOFLY_URL="https://gofly.sopans.com";
}
if (typeof GOFLY_KEFU_ID=="undefined"){
    var GOFLY_KEFU_ID="";
}
if (typeof GOFLY_BTN_TEXT=="undefined"){
    var GOFLY_BTN_TEXT="Chat with me";
}
dynamicLoadCss(GOFLY_URL+"/static/css/gofly-front.css");
if (typeof $!="function"){
    dynamicLoadJs("https://cdn.jsdelivr.net/npm/jquery/dist/jquery.min.js",function () {
        dynamicLoadJs("https://cdn.bootcdn.net/ajax/libs/layer/3.1.1/layer.min.js",function () {
            clickBtn();
        });
    });
}else{
    dynamicLoadJs("https://cdn.bootcdn.net/ajax/libs/layer/3.1.1/layer.min.js",function () {
        clickBtn();
    });
}

function clickBtn(){
    $('body').append('<div id="launchButton" class="launchButton animateUpDown"><div class="launchButtonText">'+GOFLY_BTN_TEXT+'</div></div>');
    $("#launchButton").on("click",function() {
        if (launchButtonFlag) return;
        var width=$(window).width();
        if(width<768 || isIE()>0){
            window.open(GOFLY_URL+'/chatIndex?kefu_id='+GOFLY_KEFU_ID+'&refer='+window.document.title);
            return;
        }
        layer.open({
            type: 2,
            title: GOFLY_BTN_TEXT,
            closeBtn: 1, //不显示关闭按钮
            shade: 0,
            area: ['520px', '530px'],
            offset: 'rb', //右下角弹出
            anim: 2,
            content: [GOFLY_URL+'/chatIndex?kefu_id='+GOFLY_KEFU_ID+'&refer='+window.document.title, 'yes'], //iframe的url，no代表不显示滚动条
            end: function(){
                launchButtonFlag=false;
                $(".launchButton").show();
            }
        });
        launchButtonFlag=true;
        $(".launchButton").hide();
    });
    $("body").click(function () {
        clearTimeout(titleTimer);
        document.title = originTitle;
    });
}
function isIE() {
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
function showKefu(){
    if (launchButtonFlag) return;
    var width=$(window).width();
    if(width<768 || isIE()>0){
        window.open(GOFLY_URL+'/chatIndex?kefu_id='+GOFLY_KEFU_ID+'&refer='+window.document.title);
        return;
    }
    layer.open({
        type: 2,
        title: GOFLY_BTN_TEXT,
        closeBtn: 1, //不显示关闭按钮
        shade: [0],
        area: ['520px', '530px'],
        offset: 'rb', //右下角弹出
        anim: 2,
        content: [GOFLY_URL+'/chatIndex?kefu_id='+GOFLY_KEFU_ID+'&refer='+window.document.title, 'yes'], //iframe的url，no代表不显示滚动条
        end: function(){
            launchButtonFlag=false;
            $(".launchButton").show();
        }
    });
    launchButtonFlag=true;
    $(".launchButton").hide();

    $("body").click(function () {
        clearTimeout(titleTimer);
        document.title = originTitle;
    });
}
function dynamicLoadCss(url) {
    var head = document.getElementsByTagName('head')[0];
    var link = document.createElement('link');
    link.type='text/css';
    link.rel = 'stylesheet';
    link.href = url;
    head.appendChild(link);
}
function dynamicLoadJs(url, callback) {
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

function flashTitle() {
    titleNum++;
    if (titleNum >=3) {
        titleNum = 1;
    }
    if (titleNum == 1) {
        document.title = '【】' + originTitle;
    }
    if (titleNum == 2) {
        document.title = '【你有一条消息】' + originTitle;
    }
    titleTimer = setTimeout("flashTitle()", 500);
}
window.addEventListener('message',function(e){
    var msg=e.data;
    if(msg.type=="message"){
        flashTitle();//标题闪烁
    }
});
window.onfocus = function () {
    clearTimeout(titleTimer);
    document.title = originTitle;
};

