const KEFU={
    KEFU_URL:"",
    KEFU_KEFU_ID:"",
    KEFU_ENT:"",
    KEFU_LANG:"en",
    KEFU_EXTRA: {},
    KEFU_AUTO_OPEN:true,//是否自动打开
    KEFU_SHOW_TYPES:1,//展示样式，1：普通右下角，2：圆形icon
    KEFU_AUTO_SHOW:false,
    KEFU_WITHOUT_BOX:false,
    VISITOR_ID:"",
    VISITOR_NAME:"",
    VISITOR_AVATOR:"",
};
KEFU.launchButtonFlag=false;
KEFU.titleTimer=0;
KEFU.titleNum=0;
KEFU.noticeTimer=null;
KEFU.originTitle=document.title;
KEFU.chatPageTitle="KEFU";
KEFU.kefuName="";
KEFU.kefuAvator="";
KEFU.kefuIntroduce="";
KEFU.kefuDialogDelay="3000";
KEFU.offLine=false;
KEFU.TEXT={
    "cn":{
        "online_notice":"在线咨询",
        "offline_notice":"离线留言",
    },
    "en":{
        "online_notice":"chat with us",
        "offline_notice":"we are offline",
    },
    "ru":{
        "online_notice":"Разговор с нами",
        "offline_notice":"Мы оффлайн",
    },
};

KEFU.init=function(config){
    var _this=this;

    if (!config) { config = {}; }

    for (var key in KEFU) {
        if (typeof config[key] !== 'undefined') {
            this[key] = config[key];
        } else {
            this[key] = KEFU[key];
        }
    }

    if (typeof config.KEFU_URL!="undefined"){
        this.KEFU_URL=config.KEFU_URL.replace(/([\w\W]+)\/$/,"$1");
    }
    this.dynamicLoadCss(this.KEFU_URL+"/static/css/kefu-front.css?v="+Date.now());
    this.dynamicLoadCss(this.KEFU_URL+"/static/css/layui/css/layui.css?v="+Date.now());

    var refer=document.referrer?document.referrer:"none";
    this.KEFU_EXTRA.refer=refer;
    this.KEFU_EXTRA.host=document.location.href;
    this.KEFU_EXTRA=JSON.stringify(_this.KEFU_EXTRA);

    this.dynamicLoadJs(this.KEFU_URL+"/static/js/functions.js?v=1",function(){

        _this.dynamicLoadJs("https://cdn.staticfile.org/jquery/3.6.0/jquery.min.js",function () {
            jQuery.noConflict();
            _this.dynamicLoadJs(_this.KEFU_URL+"/static/js/layer/layer.js",function () {
                _this.jsCallBack();
            });
        });
        //}
    });


    _this.addEventlisten();
}
KEFU.jsCallBack=function(){
    var _this=this;
    _this.showPcTips();
    _this.addClickEvent();
}
//pc端的样式
KEFU.showPcTips=function(){
    var _this=this;

    //自动展开
    if(_this.KEFU_AUTO_OPEN&&_this.isIE()<=0){
        setTimeout(function () {
            _this.showKefu();
        },_this.kefuDialogDelay);
    }

    var html=`
    <div class='launchButtonBox'>
        <div id="launchButton" class="launchButton">
            <div id="launchIcon" class="launchIcon">1</div>
                <div class="launchButtonText">
                    <img src="`+_this.KEFU_URL+`/static/images/wechatLogo.png"/>
                    <span class='flyUsername'>在线咨询</span>
                </div>
        </div>
        <div id="launchButtonNotice" class="launchButtonNotice"></div>

    </div>
`
    jQuery('body').append(html);
    if(_this.KEFU_AUTO_OPEN){
        return;
    }
}

KEFU.addClickEvent=function(){
    var _this=this;
    var launchButton=jQuery("#launchButton");
    if(launchButton){
        launchButton.on("click",function() {
            if(_this.launchButtonFlag){
                return;
            }
            _this.KEFU_AUTO_SHOW=true;
            _this.showKefu();
            jQuery("#launchIcon").text(0).hide();
        });
    }

}

KEFU.postMessageToIframe=function(str){
    var msg={}
    msg.type='inputing_message';
    msg.content=str;
    this.postToIframe(msg);
}
KEFU.postToIframe=function(messageObj){
    var obj=document.getElementById('layui-layer-iframe19911116');
    if(!obj||!messageObj){
        return;
    }
    document.getElementById('layui-layer-iframe19911116').contentWindow.postMessage(messageObj, "*");
}
KEFU.addEventlisten=function(){
    var _this=this;
    window.addEventListener('message',function(e){
        var msg=e.data;
        if(msg.type=="message"){
            clearInterval(_this.noticeTimer);
            var width=jQuery(window).width();
            if(width>768){
                _this.flashTitle();//标题闪烁
            }
            if (_this.launchButtonFlag){
                return;
            }
            var welcomeHtml="<div id='launchNoticeClose' class='flyClose'>×</div>"
                +"<div class='flexBox'><div class='flyUser'><img class='flyAvatar' src='"+_this.kefuAvator+"'/>"
                +   "</div>"
                +"<div class='launchNoticeContent' id='launchNoticeContent'>"
                +replaceSpecialTag(msg.data.content,_this.KEFU_URL)+"</div></div>";

            var obj=jQuery("#launchButtonNotice");
            if(obj){
                obj.html(welcomeHtml).fadeIn();
                // setTimeout(function (obj) {
                //     obj.fadeOut();
                // },3000,obj);
            }
            var news=jQuery("#launchIcon").text();
            jQuery("#launchIcon").text(++news).show();
        }
        if(msg.type=="focus"){
            clearTimeout(_this.titleTimer);
            _this.titleTimer=0;
            document.title = _this.originTitle;
        }
        if(msg.type=="force_close"){
            kayer.close(kayer.index);
        }
    });
    window.onfocus = function () {
        clearTimeout(_this.titleTimer);
        _this.titleTimer=0;
        document.title = _this.originTitle;
        _this.postToIframe({type:"focus"});
    };
}
KEFU.dynamicLoadCss=function(url){
    var head = document.getElementsByTagName('head')[0];
    var link = document.createElement('link');
    link.type='text/css';
    link.rel = 'stylesheet';
    link.href = url;
    head.appendChild(link);
}
KEFU.dynamicLoadJs=function(url, callback){
    var head = document.getElementsByTagName('head')[0];
    var script = document.createElement('script');
    script.type = 'text/javascript';
    script.src = url;
    script.defer = true;
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


KEFU.isIE=function(){
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
KEFU.showPanel=function (){
    var width=jQuery(window).width();
    this.KEFU_AUTO_SHOW=true;
    if(this.isIE()>0){
        this.windowOpen();
        return;
    }
    if(width<768){
        this.layerOpen("100%","72%","rb");
        return;
    }
    var width=380;
    var height=580;
    var x=document.documentElement.clientWidth-parseInt(width)-20;
    var y=document.documentElement.clientHeight-parseInt(height)-20;
    this.layerOpen(width+"px",height+"px",[y,x]);

    this.launchButtonFlag=true;
}
KEFU.showKefu=function (){
    if (this.launchButtonFlag) return;
    var width=jQuery(window).width();
    if(this.isIE()>0){
        this.windowOpen();
        return;
    }
    if(width<768){
        this.layerOpen("100%","72%","rb");
        return;
    }
    var width=380;
    var height=580;
    var x=document.documentElement.clientWidth-parseInt(width)-20;
    var y=document.documentElement.clientHeight-parseInt(height)-20;
    this.layerOpen(width+"px",height+"px",[y,x]);
    this.launchButtonFlag=true;
    jQuery(".launchButtonBox").hide();
}
KEFU.layerOpen=function (width,height,offset){
    if (this.launchButtonFlag) return;
    var layBox=jQuery("#layui-layer19911116");
    if(layBox.css("display")=="none"){
        this.launchButtonFlag=true;
        layBox.show();
        return;
    }
    var onlineStatus="<i class='kfBarStatus'></i>";
    if(this.offLine){
        onlineStatus="<i class='offline kfBarStatus'></i>";
    }
    var title=`
    <div class="kfBar">
        <div class="kfBarAvator">
            <img src="`+this.KEFU_URL+`/static/images/4.jpg" class="kfBarLogo">

        </div>
        <div class="kfBarText">
            <div class="kfName">在线客服系统</div>
         </div>
    </div>
    `;
    var _this=this;
    if(!offset){
        offset="rb";
    }
    var chatUrl=this.KEFU_URL+'/chatIndex?kefu_id='+this.KEFU_KEFU_ID+'&lang='+this.KEFU_LANG+'&extra='+this.KEFU_EXTRA;
    if(this.VISITOR_ID!=""){
        chatUrl+="&visitor_id="+this.VISITOR_ID;
    }
    if(this.VISITOR_NAME!=""){
        chatUrl+="&visitor_name="+this.VISITOR_NAME;
    }
    if(this.VISITOR_AVATOR!=""){
        chatUrl+="&avator="+this.VISITOR_AVATOR;
    }

    kayer.index="19911115";
    kayer.open({
        type: 2,
        title: title,
        skin:"kfLayer",
        closeBtn: 1, //不显示关闭按钮
        shade: 0,
        area: [width, height],
        offset: offset, //右下角弹出
        anim: 2,
        content: [chatUrl , 'yes'], //iframe的url，no代表不显示滚动条
        move:false,
        success:function(layero, index){

            var layBox=jQuery("#layui-layer19911116");
            _this.launchButtonFlag=true;
            if(!_this.KEFU_WITHOUT_BOX&&_this.KEFU_AUTO_SHOW&&layBox.css("display")=="none"){
                layBox.show();
            }
            jQuery("#layui-layer-iframe19911116 .chatEntTitle").hide();
        },
        end: function(){
            _this.launchButtonFlag=false;
            jQuery(".launchButtonBox").show();
        },
        cancel: function(index, layero){
            jQuery("#layui-layer19911116").hide();
            _this.launchButtonFlag=false;
            jQuery(".launchButtonBox").show();
            return false;
        }
    });
}
KEFU.windowOpen=function (){
    window.open(this.KEFU_URL+'/chatIndex?kefu_id='+this.KEFU_KEFU_ID+'&lang='+this.KEFU_LANG+'&refer='+window.document.title+'&ent_id='+this.KEFU_ENT+'&extra='+this.KEFU_EXTRA);
}
KEFU.flashTitle=function () {
    flashTitle();
}

/**
 * 判断是否是手机访问
 * @returns {boolean}
 */
KEFU.isMobile=function () {
    if( /Mobile|Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent) ) {
        return true;
    }
    return false;
}

