# go-fly
基于GO语言实现的web客服即时通讯与客服管理系统。

1.使用gin http框架实现restful风格的API

2.使用jwt-go配合gin中间件实现无状态的jwt登陆认证

3.使用casbin配合gin中间件实现权限控制

4.使用gin以及template包的模板语法进行展示界面

5.使用go modoule解决依赖问题

6.使用go-imap实现邮件的列表展示和读取

7.使用go-smtp实现发送邮件

8.使用github.com/gorilla/websocket实现即时通讯

9.使用gorm配合mysql实现数据存储

10.充分实践了struct，interface，map，slice，for range,groutine和channel管道等基础知识

### 项目预览

![Image text](https://img2020.cnblogs.com/blog/726254/202006/726254-20200628213102562-970752675.jpg)

![Image text](https://img2020.cnblogs.com/blog/726254/202006/726254-20200628213122308-1498566060.jpg)

![Image text](https://img2020.cnblogs.com/blog/726254/202006/726254-20200628213137942-667076789.jpg)


### 安装使用


1. 先安装和运行mysql , 创建go-fly数据库，并导入*.sql创建表结构与数据，在config

2. 在go get github.com/taoshihan1991/go-fly ,进入$GOPATH/src/github.com/taoshihan1991/go-fly 目录
   
   目录mysql.json中配置数据库
```php
{
	"Server":"127.0.0.1",
	"Port":"3306",
	"Database":"go-fly",
	"Username":"go-fly",
	"Password":"go-fly"
}
```


3. 源码运行 go run server.go

4. 源码打包 go build server.go

