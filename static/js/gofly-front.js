var launchButtonFlag=false;
$("#launchButton").click(function() {
    if (launchButtonFlag) return;
    layer.open({
        type: 2,
        title: "Chat with us",
        closeBtn: 1, //不显示关闭按钮
        shade: [0],
        area: ['520px', '530px'],
        offset: 'rb', //右下角弹出
        anim: 2,
        content: [GOFLY_URL+'/chat_page?refer='+window.location.host, 'no'], //iframe的url，no代表不显示滚动条
        end: function(){
            launchButtonFlag=false;
            $(".launchButton").show();
        }
    });
    launchButtonFlag=true;
    $(".launchButton").hide();
});

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
