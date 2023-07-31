### 郑重提示
禁止将本项目用于含病毒、木马、色情、赌博、诈骗、违禁用品、假冒产品、虚假信息、数字货币、金融等违法违规业务

当前项目仅供个人学习测试，禁止一切线上商用行为，禁止一切违法使用！！！




### 项目简介

Golang语言开源客服系统，主要使用了gin + jwt-go + websocket + go.uuid + gorm + cobra + VueJS + ElementUI + MySQL等技术


### 安装使用


* 先安装和运行mysql数据库 ，版本>=5.5 ，创建数据库
 
```
 create database gofly charset utf8mb4;
```
   
*  配置数据库链接信息，在config目录mysql.json中
```php
{
	"Server":"127.0.0.1",
	"Port":"3306",
	"Database":"gofly",
	"Username":"go-fly",
	"Password":"go-fly"
}
```
* 安装配置Golang运行环境，请参照下面的命令去执行
```php
wget https://studygolang.com/dl/golang/go1.20.2.linux-amd64.tar.gz
tar -C /usr/local -xvf go1.20.2.linux-amd64.tar.gz
mv go1.20.2.linux-amd64.tar.gz /tmp
echo "PATH=\$PATH:/usr/local/go/bin" >> /etc/profile
echo "PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
source /etc/profile
go version
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```
* 下载代码

    在任意目录 git clone https://github.com/taoshihan1991/go-fly.git
    
    进入go-fly 目录
   
* 导入数据库 go run go-fly.go install

* 源码运行 go run go-fly.go server

* 源码打包 go build -o kefu  会生成kefu可以执行文件

* 二进制文件运行

   linux:   ./kefu server [可选 -p 8082 -d]
   
   windows: kefu.exe server [可选 -p 8082 -d]
   
* 关闭程序

   killall kefu


程序正常运行后，监听端口8081，可以直接ip+端口8081访问

也可以配置域名访问，反向代理到8081端口，就能隐藏端口号
### 客服对接
聊天链接

http://127.0.0.1:8081/chatIndex?kefu_id=kefu2

弹窗使用

```
    (function(a, b, c, d) {
        let h = b.getElementsByTagName('head')[0];let s = b.createElement('script');
        s.type = 'text/javascript';s.src = c+"/static/js/kefu-front.js";s.onload = s.onreadystatechange = function () {
            if (!this.readyState || this.readyState === "loaded" || this.readyState === "complete") d(c);
        };h.appendChild(s);
    })(window, document,"http://127.0.0.1:8081",function(u){
        KEFU.init({
            KEFU_URL:u,
            KEFU_KEFU_ID: "kefu2",
        })
    });

```
### 版权声明

当前项目是完整功能代码 , 但是仍然仅支持个人演示测试 , 不包含线上使用 ，禁止一切商用行为。
使用本软件时,请遵守当地法律法规,任何违法用途一切后果请自行承担.