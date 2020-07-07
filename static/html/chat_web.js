var loadJs=function(url,callback){
    var script = document.createElement('script'), fn = callback || function(){};
    script.type = 'text/javascript';
    //IE
    if(script.readyState){
        script.onreadystatechange = function(){
            if( script.readyState == 'loaded' || script.readyState == 'complete' ){
                script.onreadystatechange = null;
                fn();
            }
        };
    }else{
        //其他浏览器
        script.onload = function(){
            fn();
        };
    }
    script.src = url;
    document.getElementsByTagName('head')[0].appendChild(script);
};
loadJs("https://cdn.jsdelivr.net/npm/jquery/dist/jquery.min.js",function(){
    loadJs("https://cdn.bootcdn.net/ajax/libs/layer/3.1.1/layer.min.js" ,function () {
        $(function () {
            //====================================================================
            var div =document.createElement('div');
            div.id ='goflyKefu';
            div.className +='goflyKefu';
            document.body.appendChild(div);
            var w =document.getElementById('goflyKefu');
            w.innerHTML='<div style="border-radius:5px;position: fixed;right: 10px;bottom: 10px;background: #66b1ff;padding: 10px 30px;color: #fff;cursor: pointer;">在线咨询</div>';

            $("#goflyKefu").click(function () {
                $("#goflyKefu").hide();
                layer.open({
                    type: 2,
                    title: '在线咨询',
                    shadeClose: true,
                    shade: false,
                    maxmin: true, //开启最大化最小化按钮
                    area: ['660px', '600px'],
                    content: ['http://gofly.sopans.com/chat_page','no'],
                    end: function(){
                        $("#goflyKefu").show();
                    }
                });
            });
            //---------------------------------------------------------------
        })
    });
});


