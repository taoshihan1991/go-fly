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
//除去html标签
function replaceHtml(str){
    return str.replace(/<[^>]*>/g, '');
}
//浏览器桌面通知
function notify(title, options, callback) {

    // 先检查浏览器是否支持
    if (!window.Notification) {
        console.log("浏览器不支持notify");
        return;
    }
    options.body=replaceHtml(options.body);
    console.log("浏览器notify权限:", Notification.permission);
    // 检查用户曾经是否同意接受通知
    if (Notification.permission === 'granted') {
        var notification = new Notification(title, options); // 显示通知
        if (notification && callback) {
            notification.onclick = function(event) {
                callback(notification, event);
            }
            setTimeout(function () {
                notification.close();
            },3000);
        }
    } else {
        Notification.requestPermission().then( (permission) =>function(){
            console.log("请求浏览器notify权限:", permission);
            if (permission === 'granted') {
                notification = new Notification(title, options); // 显示通知
                if (notification && callback) {
                    notification.onclick = function (event) {
                        callback(notification, event);
                    }
                    setTimeout(function () {
                        notification.close();
                    }, 3000);
                }
            } else if (permission === 'default') {
                console.log('用户关闭授权 可以再次请求授权');
            } else {
                console.log('用户拒绝授权 不能显示通知');
            }
        });
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
    content = (content || '')
        .replace(/face\[(.*?)\]/g, function (face) {  // 转义表情
            var alt = face.replace(/^face/g, '');
            return '<img alt="' + alt + '" title="' + alt + '" src="'+baseUrl + faces[alt] + '">';
        })
        .replace(/img\[(.*?)\]/g, function (face) {  // 转义图片
            var src = face.replace(/^img\[/g, '').replace(/\]/g, '');;
            return '<img onclick="bigPic(src,true)" src="' +baseUrl+ src + '" style="max-width: 150px"/></div>';
        })
        .replace(/\n/g, '<br>'); // 转义换行
    content=replaceAttachment(content);
    return content;
}
//替换附件展示
function replaceAttachment(str){
    return str.replace(/attachment\[(.*?)\]/g, function (result) {
        var mutiFiles=result.match(/attachment\[(.*?)\]/)
        if (mutiFiles.length<2){
            return result;
        }
        //return result;

        var info=JSON.parse(mutiFiles[1])
        var imgSrc="";
        switch(info.ext){
            case ".mp3":
                imgSrc="/static/images/ext/MP3.png";
                break;
            case ".zip":
                imgSrc="/static/images/ext/ZIP.png";
                break;
            case ".txt":
                imgSrc="/static/images/ext/TXT.png";
                break;
            case ".7z":
                imgSrc="/static/images/ext/7z.png";
                break;
            case ".bpm":
                imgSrc="/static/images/ext/BMP.png";
                break;
            case ".png":
                imgSrc="/static/images/ext/PNG.png";
                break;
            case ".jpg":
                imgSrc="/static/images/ext/JPG.png";
                break;
            case ".jpeg":
                imgSrc="/static/images/ext/JPEG.png";
                break;
            case ".pdf":
                imgSrc="/static/images/ext/PDF.png";
                break;
            case ".doc":
                imgSrc="/static/images/ext/DOC.png";
                break;
            case ".docx":
                imgSrc="/static/images/ext/DOCX.png";
                break;
            case ".rar":
                imgSrc="/static/images/ext/RAR.png";
                break;
            case ".xlsx":
                imgSrc="/static/images/ext/XLSX.png";
                break;
            case ".csv":
                imgSrc="/static/images/ext/XLSX.png";
                break;
            default:
                imgSrc="/static/images/ext/default.png";
                break;
        }
        var html= `<div onclick="window.open('`+info.path+`')" class="productCard">
                        <div><img src='`+imgSrc+`' style='width: 38px;height: 38px;' /></div>
                        <div class="productCardTitle">
                            <div class="productCardTitle">`+info.name+`</div>
                            <div style="font-size: 12px;color: #666">`+formatFileSize(info.size)+`</div>
                        </div>
                    </div>`;
        return html;
    })
}
function formatFileSize(fileSize) {
    if (fileSize < 1024) {
        return fileSize + 'B';
    } else if (fileSize < (1024*1024)) {
        var temp = fileSize / 1024;
        temp = temp.toFixed(2);
        return temp + 'KB';
    } else if (fileSize < (1024*1024*1024)) {
        var temp = fileSize / (1024*1024);
        temp = temp.toFixed(2);
        return temp + 'MB';
    } else {
        var temp = fileSize / (1024*1024*1024);
        temp = temp.toFixed(2);
        return temp + 'GB';
    }
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
//播放声音
function alertSound(id,src){
    var b = document.getElementById(id);
    if(src!=""){
        b.src=src;
    }
    var p = b.play();
    p && p.then(function(){}).catch(function(e){});
}
//日期格式化
function formatDate(dateString, format = 'yyyy-MM-dd HH:mm:ss') {
    const date = new Date(dateString);

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hour = String(date.getHours()).padStart(2, '0');
    const minute = String(date.getMinutes()).padStart(2, '0');
    const second = String(date.getSeconds()).padStart(2, '0');

    const formattedDate = format
        .replace(/yyyy/g, year)
        .replace(/MM/g, month)
        .replace(/dd/g, day)
        .replace(/HH/g, hour)
        .replace(/mm/g, minute)
        .replace(/ss/g, second);

    return formattedDate;
}
function copyText(text) {
    var target = document.createElement('input')
    target.value = text
    document.body.appendChild(target)
    target.select()
    document.execCommand("copy");
    document.body.removeChild(target);
    return true;
}
;
