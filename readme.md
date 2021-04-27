# go-fly
基于Golang语言和MySQL实现的WEB在线客服系统

主要技术栈
 gin + jwt-go + websocket + go.uuid + gorm + cobra

用法
docker build -t go-fly:master .  
docker run -p 8081:8081 -v config:/app/config --name go-fly go-fly:master

### 更新日志
##### V0.4.1

访客端咨询按钮的样式修改了

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

3.如果遇到域名跨域错误问题, 检查下面配置中add_header Access-Control-Allow-Origin这俩header头是否添加

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
location  ~* \.(php|jsp|cgi|asp|aspx)$
{
	add_header Access-Control-Allow-Origin *;
	add_header Access-Control-Allow-Methods 'GET, POST, OPTIONS';
    proxy_pass http://127.0.0.1:8081;
	proxy_http_version 1.1;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header REMOTE-HOST $remote_addr;
    
	proxy_set_header Upgrade $http_upgrade;
	proxy_set_header Connection "upgrade";
	proxy_set_header Origin "";
}
location /
{
	add_header Access-Control-Allow-Origin *;
	add_header Access-Control-Allow-Methods 'GET, POST, OPTIONS';
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


### 感谢赞助

2021年04月25日 **P7  88.88元

2021年04月19日 **指  1400元(多商户)

2021年04月01日 **科技  66.66元

2021年03月15日 **无畏  8.88元

2021年03月15日 **彬  77元

2021年03月10日 ABC  100元

2021年02月20日 广西***社  1000元

2021年02月19日 **辉  1000元

2021年02月04日 **宏  10.24元

2021年02月03日 pony  188元

2021年01月22日 **~  1000元(多商户)

2021年01月20日 **生  8.88元

2021年01月17日 **白  8.88元

2021年01月13日 **~  500元(多商户)

2020年12月31日 **强  8.88元

2020年12月24日 **松  8.88元

2020年12月23日 **渊  10元

2020年12月16日 **彬  8.8元

2020年11月30日 **宇  88元

### 打赏作者


<img src="https://gofly.sopans.com/static/images/alipay.jpg" width="280"  alt="支付宝"/>
<img src="https://gofly.sopans.com/static/images/weixin.jpg" width="280"  alt="微信"/>

### 版权声明

当前项目是完整功能代码 , 但是仍然仅支持个人演示测试 , 不包含线上使用 . 赞赏并联系作者后可得到作者授权 , 并且可以获取完整专属技术支持,包括安装/部署/bug修改以及后期功能升级 . 
使用本软件时,请遵守当地法律法规,任何违法用途一切后果请自行承担.