function getBaseUrl() {
    var ishttps = 'https:' == document.location.protocol ? true : false;
    var url = window.location.host;
    if (ishttps) {
        url = 'https://' + url;
    } else {
        url = 'http://' + url;
    }
    return url;
}
function getWsBaseUrl() {
    var ishttps = 'https:' == document.location.protocol ? true : false;
    var url = window.location.host;
    if (ishttps) {
        url = 'wss://' + url;
    } else {
        url = 'ws://' + url;
    }
    return url;
}
function notify(title, options, callback) {
    // 先检查浏览器是否支持
    if (!window.Notification) {
        return;
    }
    var notification;
    // 检查用户曾经是否同意接受通知
    if (Notification.permission === 'granted') {
        notification = new Notification(title, options); // 显示通知

    } else {
        var promise = Notification.requestPermission();
    }

    if (notification && callback) {
        notification.onclick = function(event) {
            callback(notification, event);
        }
        setTimeout(function () {
            notification.close();
        },3000);
    }
}
var titleTimer=0;
var titleNum=0;
var originTitle = document.title;
function flashTitle() {
    if(titleTimer!=0){
        return;
    }
    titleTimer = setInterval(function(){
        titleNum++;
        if (titleNum == 3) {
            titleNum = 1;
        }
        if (titleNum == 1) {
            document.title = '【】' + originTitle;
        }
        if (titleNum == 2) {
            document.title = '【new message】' + originTitle;
        }
    }, 500);

}
function clearFlashTitle() {
    clearInterval(titleTimer);
    document.title = originTitle;
}
var faceTitles = ["[微笑]", "[嘻嘻]", "[哈哈]", "[可爱]", "[可怜]", "[挖鼻]", "[吃惊]", "[害羞]", "[挤眼]", "[闭嘴]", "[鄙视]", "[爱你]", "[泪]", "[偷笑]", "[亲亲]", "[生病]", "[太开心]", "[白眼]", "[右哼哼]", "[左哼哼]", "[嘘]", "[衰]", "[委屈]", "[吐]", "[哈欠]", "[抱抱]", "[怒]", "[疑问]", "[馋嘴]", "[拜拜]", "[思考]", "[汗]", "[困]", "[睡]", "[钱]", "[失望]", "[酷]", "[色]", "[哼]", "[鼓掌]", "[晕]", "[悲伤]", "[抓狂]", "[黑线]", "[阴险]", "[怒骂]", "[互粉]", "[心]", "[伤心]", "[猪头]", "[熊猫]", "[兔子]", "[ok]", "[耶]", "[good]", "[NO]", "[赞]", "[来]", "[弱]", "[草泥马]", "[神马]", "[囧]", "[浮云]", "[给力]", "[围观]", "[威武]", "[奥特曼]", "[礼物]", "[钟]", "[话筒]", "[蜡烛]", "[蛋糕]"];
function placeFace() {
    var faces=[];
    for(var i=0;i<faceTitles.length;i++){
        faces[faceTitles[i]]="/static/images/face/"+i+".gif";
    }
    return faces;
}
function replaceContent (content,baseUrl) {// 转义聊天内容中的特殊字符
    if(typeof baseUrl=="undefined"){
        baseUrl="";
    }
    var faces=placeFace();
    var html = function (end) {
        return new RegExp('\\n*\\[' + (end || '') + '(pre|div|span|p|table|thead|th|tbody|tr|td|ul|li|ol|li|dl|dt|dd|h2|h3|h4|h5)([\\s\\S]*?)\\]\\n*', 'g');
    };
    content = (content || '').replace(/&(?!#?[a-zA-Z0-9]+;)/g, '&amp;')
        .replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/'/g, '&#39;').replace(/"/g, '&quot;') // XSS
        .replace(/face\[([^\s\[\]]+?)\]/g, function (face) {  // 转义表情
            var alt = face.replace(/^face/g, '');
            return '<img alt="' + alt + '" title="' + alt + '" src="'+baseUrl + faces[alt] + '">';
        })
        .replace(/img\[([^\s\[\]]+?)\]/g, function (face) {  // 转义图片
            var src = face.replace(/^img\[/g, '').replace(/\]/g, '');;
            return '<img onclick="bigPic(src,true)" src="' +baseUrl+ src + '" style="max-width: 100%"/></div>';
        })
        .replace(/file\[([^\s\[\]]+?)\]/g, function (face) {  // 转义图片
            var src = face.replace(/^file\[/g, '').replace(/\]/g, '');;
            return '<div class="folderBtn" onclick="window.open(\''+baseUrl+src+'\')"  style="font-size:25px;"/></div>';
        })
        .replace(/\[([^\s\[\]]+?)\]+link\[([^\s\[\]]+?)\]/g, function (face) {  // 转义超链接
            var text = face.replace(/link\[.*?\]/g, '').replace(/\[|\]/g, '');
            var src = face.replace(/^\[([^\s\[\]]+?)\]+link\[/g, '').replace(/\]/g, '');
            return '<a href="'+src+'" target="_blank" />【'+text+'】</a>';
        })
        .replace(html(), '\<$1 $2\>').replace(html('/'), '\</$1\>') // 转移HTML代码
        .replace(/\n/g, '<br>') // 转义换行

    return content;
}
function bigPic(src,isVisitor){
    if (isVisitor) {
        window.open(src);
        return;
    }
}
function filter (obj){
    var imgType = ["image/jpeg","image/png","image/jpg","image/gif"];
    var filetypes = imgType;
    var isnext = false;
    for (var i = 0; i < filetypes.length; i++) {
        if (filetypes[i] == obj.type) {
            return true;
        }
    }
    return false;
}
function sleep(time) {
    var startTime = new Date().getTime() + parseInt(time, 10);
    while(new Date().getTime() < startTime) {}
}
function checkLang(){
    var langs=["cn","en"];
    var lang=getQuery("lang");
    if(lang!=""&&langs.indexOf(lang) > 0 ){
        return lang;
    }
    return "cn";
}
function getQuery(key) {
    var query = window.location.search.substring(1);
    var key_values = query.split("&");
    var params = {};
    key_values.map(function (key_val){
        var key_val_arr = key_val.split("=");
        params[key_val_arr[0]] = key_val_arr[1];
    });
    if(typeof params[key]!="undefined"){
        return params[key];
    }
    return "";
}
function utf8ToB64(str) {
    return window.btoa(unescape(encodeURIComponent(str)));
}
function b64ToUtf8(str) {
    return decodeURIComponent(escape(window.atob(str)));
}
;
