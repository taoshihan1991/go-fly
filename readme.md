# go-fly
基于Golang语言和MySQL实现的WEB在线客服系统

主要技术栈
 gin + jwt-go + websocket + go.uuid + gorm + cobra

### 更新日志


作者本人现在全职开发，后续更新转为了商务版，欢迎大家支持商务版 https://gofly.sopans.com

#### v0.6.0

+ 新增访客黑名单功能，可以根据访客id加入黑名单
+ 新增api接口可优雅关闭服务，在守护模式下相当于重启服务
+ 新增客服端搜索客服账号接口
+ 新增微信公众号访客展示是否关注公众号标签
+ 新增客服端首页展示系统公告，超管可以添加管理系统公告
+ 新增系统配置微信模板remark字段，模板消息会展示该字段
+ 修复优化访客表新增real_name字段，客服端首先展示该字段，客服备注姓名存入该字段
* 修复子进程退出次数太多，父守护进程也退出问题
* 修复优化传递商品卡片信息样式效果
* 修复优化访客端图片缩略展示，点击预览大图效果

#### v0.5.9
+ 新增系统配置项，系统管理员权限可以在后台配置本客服系统的标题、关键字、版权等基本信息，以及是否显示注册按钮等
+ 新增商户账号可以查看访客的基本信息，可以查看访客是否绑定了微信公众号
+ 新增商户账号下访客列表展示ip和对应的地址
+ 新增清除访客聊天记录
+ 新增商户配置项，上传微信域名验证文件功能
+ 新增生成微信公众号菜单可视化编辑功能
* 新增微信公众号主动发客服消息接口是否开启配置项
* 修复微信公众号和访客绑定功能，访客关注时新增绑定，访客取关时删除绑定
* 修复访客端超时，另开tab标签等操作时，弹窗确认重新reload页面

#### v0.5.8
+ 新增微信公众号网页oauth授权功能，网页获取微信用户的昵称和头像
+ 新增微信公众号关注时自动回复消息功能，与原来的自动欢迎拆分开
+ 新增访客页面二次跳转到落地域名的配置项
+ 新增微信公众号模板消息，可以给客服、访客发送新消息提醒模板消息
+ 新增微信公众号带参二维码，绑定客服与微信id
+ 新增商户后台配置模板id功能，增加客服消息、访客消息，访客上线三个模板配置
+ 修改客服相关接口的接口前缀
* 修复访客id分割错误问题，访客id连接符修改
* 修复微信公众号access_token获取接口次数超限制问题

#### v0.5.5
+ 新增访客端显示客服在线离线状态，判断客服离线状态
+ 新增自动欢迎内容增加富文本编辑器wangEditor
+ 新增独立链接模式传递访客名称，id和头像信息
+ 新增客服界面展示访客操作系统、浏览器等信息
+ 新增给访客打tag标签功能，可以根据tag标签搜索
+ 新增访客列表搜索功能，根据访客id或者名称进行搜索
+ 修改访客端消息时间的格式
* 修复通知邮件465端口不能发送问题，支持tls的smtp端口
* 修复访客端websocket连接最大次数限制不起作用问题

#### v0.5.1
...省略

此后转为商务版


##### V0.4.1

访客端咨询按钮的样式修改,滚动区域修改

访客端浏览器提醒自动消失

客服端可以编辑自动回复内容了

命令行参数中新增了关闭服务的功能如:./go-fly stop



##### V0.3.9

利用go1.16特性进行内嵌资源 , 把模板和js内嵌入二进制文件

增加安装界面,访问[域名]/install
进入安装界面,填写数据库信息,会自动写入配置并且导入数据库


##### V0.3.8

访客端输入框以及图标icon按钮修改

客服端界面icon修改

修复后端发消息空指针错误导致的进程退出

后端代码增加了允许跨域的http头,所以可以把nginx中的跨域相关http头可以去掉

##### V0.3.7

访客端增加自助服务点击后可以自动回复

访客端手机端咨询按钮移到右侧不遮挡底部

访客端前端修复多个layer冲突问题

后端修改守护进程方式,进程崩溃后可自动重启

后端增加定时清理频限防止内存泄露

后端增加通知频限和访客输入频限防止死锁

编译增加linux-x86_64/linux-i686版本支持

##### V0.3.6

修复访客端标题闪烁问题

优化访客端头像样式以及小键盘遮挡问题

优化发消息问题

新增访客关键词自动回复功能

客服端增加附件上传功能

客服端增加关键词自动回复功能

##### V0.3.5

新增分开系统自动断线与客服关闭连接

修复没有设置欢迎时tip显示错误问题

修复客服端发送消息错误提示不显示问题

修复一些界面问题

##### V0.3.4

修复发送死锁问题

##### V0.3.3

1.访客/客服端聊天界面样式修改

2.访客端展示客服头像信息

3.访客到来自动打开,以及参数控制

4.访客/客服端聊天信息默认折叠

5.客服端新消息提醒标识

6.客服端访客列表展示访客正在输入信息
   
##### V0.3.2
   1.修改访客界面样式，更加简洁扁平
   
   2.修改自动欢迎界面样式增加聊天框效果
   
   3.修改数据库时间字段类型,兼容mysql5.7
   
   4.修复数据库执行sql获取错误信息
   
##### V0.3.1
   1.修改在线咨询浮框样式
   
   2.修改数据库时间字段类型,兼容mysql5.5+
##### V0.2.9
    
   1.访客开多个窗口时 , 单点登录关闭旧ws连接
   
   2.访客切换窗口时可以自动重连
   
   3.访客到来时 , http接口和ws接口同时发送给客服上线信息
   
   4.客服后台定时拉取在线访客接口
   
   5.客服后台切换tab拉取在线访客


### 项目预览

![Image text](https://gitee.com/taoshihan/go-fly/raw/master/static/images/newintro1.jpg)

![Image text](https://img2020.cnblogs.com/blog/726254/202009/726254-20200902141707515-1201702349.jpg)

![Image text](https://gitee.com/taoshihan/go-fly/raw/master/static/images/newintro2.jpg)

![Image text](https://img2020.cnblogs.com/blog/726254/202009/726254-20200902141736713-1155907367.jpg)

![Image text](https://gitee.com/taoshihan/go-fly/raw/master/static/images/newintro3.jpg)


### 安装使用


#### 1. 先安装和运行mysql >=5.5版本 , 创建gofly数据库.
 
    create database gofly charset utf8;
   
   在config目录mysql.json中配置数据库
```php
{
	"Server":"127.0.0.1",
	"Port":"3306",
	"Database":"gofly",
	"Username":"go-fly",
	"Password":"go-fly"
}
```
#### 2. 二进制文件运行

   1) 下载地址 

        github: https://github.com/taoshihan1991/go-fly/releases/
   
        gitee(国内): https://gitee.com/taoshihan/go-fly/releases
   
   2) 文件解压缩 
   
        windows系统下,在cmd命令行,进入项目解压后目录; linux系统下创建目录执行如下 
        
            linux服务器:
            mkdir go-fly
            cd go-fly
            wget xxxxxxxxxxx.zip
            unzip xxxx.zip
            chmod 0777 -R ./
        
        导入数据库( 注意:会删除表并且清空数据 ) 
        
            windows: go-fly.exe install
            
            linux: ./go-fly install

        运行项目
        
            linux:   ./go-fly server [可选 -p 8082 -d]
           
            windows: go-fly.exe server [可选 -p 8082]
        
   3) 参数说明
    
        -p 指定端口
        
        -d linux下是否以daemon守护进程运行
        
        -h 查看帮助 
        
#### 3. 源码运行

1. 基于go module使用

   go env -w GO111MODULE=on
   
   go env -w GOPROXY=https://goproxy.cn,direct
   
   在任意目录 git clone https://github.com/taoshihan1991/go-fly.git
   
   进入go-fly 目录

2. 源码运行 go run go-fly.go server

3. 源码打包 go build go-fly.go 会生成go-fly可以执行文件

4. 导入数据库(会删除表清空数据) ./go-fly install

5. 二进制文件运行
 
   linux:   ./go-fly server [可选 -p 8082 -d]
   
   windows: go-fly.exe server [可选 -p 8082 -d]

6. 关闭程序
   ./go-fly stop  

   linux下使用ps命令结合kill命令杀掉进程
   
   ps -ef|grep go-fly 看到父子进程id
   
   kill 进程父进程id ； kill 进程子进程id

#### 4. 网页使用

   1.服务端安装成功后可把域名换成自己的域名或IP
   
   2.默认访问本地http://127.0.0.1:8081
   
```php
    //下面js路径和GOFLY_URL 都要改成自己的
    <script src="https://gofly.sopans.com/assets/js/gofly-front.js"></script>
    <script>
        GOFLY.init({
            GOFLY_URL:"https://gofly.sopans.com",
            GOFLY_KEFU_ID: "kefu2",
            GOFLY_BTN_TEXT: "客服在线 欢迎咨询",
            GOFLY_LANG:"cn"
        })
    </script>
```

   
### nginx部署

访问：https://gofly.sopans.com

1.参考支持https的部署示例 , 注意反向代理的端口号和证书地址 , 不使用https也可以访问 , 只是不会有浏览器通知弹窗

2.尽量按照下面的配置处理, 配置独立域名或者二级域名, 不建议在主域名加端口访问, 不建议主域名加目录访问 

3.<del>如果遇到域名跨域错误问题, 检查下面配置中add_header Access-Control-Allow-Origin这俩header头是否添加.</del>
代码里已经解决跨域 , nginx里不要加跨域头,否则会冲突报错

```php
server {
       listen 443 ssl http2;
        ssl on;
        ssl_certificate   conf.d/cert/4263285_gofly.sopans.com.pem;
        ssl_certificate_key  conf.d/cert/4263285_gofly.sopans.com.key;
        ssl_session_timeout 5m;
        ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
        ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
        ssl_prefer_server_ciphers on;
        #listen          80; 
        server_name  gofly.sopans.com;
        access_log  /var/log/nginx/gofly.sopans.com.access.log  main;
        location /static {
                root /var/www/html/go-fly;//自己的部署路径,静态文件直接nginx响应
        }
        location / {
                proxy_pass http://127.0.0.1:8081;
                    proxy_http_version 1.1;
                    proxy_set_header X-Real-IP $remote_addr;
                    proxy_set_header Upgrade $http_upgrade;
                    proxy_set_header Connection "upgrade";
                    proxy_set_header Origin "";
        }
}
server{
       listen 80;
        server_name  gofly.sopans.com;
        access_log  /var/log/nginx/gofly.sopans.com.access.log  main;
        location /static {
                root /var/www/html/go-fly;//自己的部署路径,静态文件直接nginx响应
        }        
        location / {
                proxy_pass http://127.0.0.1:8081;
                    proxy_http_version 1.1;
                    proxy_set_header X-Real-IP $remote_addr;
                    proxy_set_header Upgrade $http_upgrade;
                    proxy_set_header Connection "upgrade";
                    proxy_set_header Origin "";
        }
}
```
### 宝塔部署

原文地址：https://www.zqcnc.cn/post/99.html

#### 宝塔环境
1. 创建一个静态站点，地址为想要访问的域名
![](https://i.aweoo.com/imgs/2021/03/9662692b88b802f9.png)
2. 为该站点配置证书
![](https://i.aweoo.com/imgs/2021/03/9c2f91a215d37b2f.png)
3. 设置反向代理
![](https://i.aweoo.com/imgs/2021/03/61cdce0167949ff4.png)
4. 修改反代配置
![](https://i.aweoo.com/imgs/2021/03/2a5aa9783afa9a19.png)
**按照图示，将对应代码加入到配置文件中**
```shell
#PROXY-START/
location /
{
    proxy_pass http://127.0.0.1:8081;
	proxy_http_version 1.1;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header REMOTE-HOST $remote_addr;
    
	proxy_set_header Upgrade $http_upgrade;
	proxy_set_header Connection "upgrade";
	proxy_set_header Origin "";
    
    add_header X-Cache $upstream_cache_status;
    
    #Set Nginx Cache
    
    	add_header Cache-Control no-cache;
    expires 12h;
}

#PROXY-END/
```



### 联系作者

QQ: 630892807

微信：

<img src="https://gofly.sopans.com/static/images/wechat.jpg" width="280"  alt="微信"/>

### 版权声明

当前项目是完整功能代码 , 但是仍然仅支持个人演示测试 , 不包含线上使用 . 赞赏并联系作者后可得到作者授权 , 并且可以获取完整专属技术支持,包括安装/部署/bug修改以及后期功能升级 . 
使用本软件时,请遵守当地法律法规,任何违法用途一切后果请自行承担.