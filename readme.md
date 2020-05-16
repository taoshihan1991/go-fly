# go-imap
邮箱imap网页版客户端工具，基于GO语言实现。

1.使用第三方类库go-imap解析imap协议

2.使用http包 ，template包，实现http服务下的网页展示

3.使用goroutine在主界面并发请求左右栏的数据

4.使用text/net包下的encoding和transform等配合解决乱码问题

5.使用go modoule解决依赖问题

6.充分实践了struct，interface，map，slice，for range等基础知识

###项目预览
![Image text](https://img2020.cnblogs.com/blog/726254/202005/726254-20200516183509721-1526715752.png)

![Image text](https://img2020.cnblogs.com/blog/726254/202005/726254-20200516183521692-598821905.png)

![Image text](https://img2020.cnblogs.com/blog/726254/202005/726254-20200516183534997-76603458.png)


###安装使用
1.git clone https://github.com/taoshihan1991/imaptool.git

2.进入目录执行 go mod tidy

3.源码运行 go run server.go

4.源码打包 go build server.go

