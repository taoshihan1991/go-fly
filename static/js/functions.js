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
    }
}
var titleTimer,titleNum=0;
var originTitle = document.title;
function flashTitle() {
    titleNum++;
    if (titleNum == 3) {
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

function replaceContent (content) {// 转义聊天内容中的特殊字符

}