package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/taoshihan1991/imaptool/config"
	"github.com/taoshihan1991/imaptool/controller"
	"github.com/taoshihan1991/imaptool/docs"
	"github.com/taoshihan1991/imaptool/middleware"
	"github.com/taoshihan1991/imaptool/tmpl"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)
var (
	port string
	daemon bool
	GoflyConfig config.Config
)
func init(){
	//获取参数中的数据
	flag.StringVar(&port, "port", "8080", "监听端口号")
	flag.BoolVar(&daemon, "d", false, "是否为守护进程模式")
	flag.Parse()
	if flag.NFlag() < 1 {
		flag.PrintDefaults()
	}
	if daemon==true{
		if os.Getppid() != 1{
			// 将命令行参数中执行文件路径转换成可用路径
			filePath, _ := filepath.Abs(os.Args[0])
			cmd := exec.Command(filePath, os.Args[1:]...)
			// 将其他命令传入生成出的进程
			cmd.Stdin = os.Stdin // 给新进程设置文件描述符，可以重定向到文件中
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Start() // 开始执行新进程，不等待新进程退出
			os.Exit(0)
		}
	}
}
func main() {

	baseServer := "0.0.0.0:"+port
	log.Println("start server...\r\ngo：http://" + baseServer)
	engine := gin.Default()
	engine.LoadHTMLGlob("static/html/*")
	engine.Static("/static", "./static")
	//首页
	engine.GET("/", controller.Index)
	engine.GET("/index", tmpl.PageIndex)
	//登陆界面
	engine.GET("/login", tmpl.PageLogin)
	//咨询界面
	engine.GET("/chat_page",middleware.SetLanguage, tmpl.PageChat)
	//登陆验证
	engine.POST("/check", controller.LoginCheckPass)
	//框架界面
	engine.GET("/main",middleware.JwtPageMiddleware,tmpl.PageMain)
	//框架界面
	engine.GET("/chat_main",middleware.JwtPageMiddleware,tmpl.PageChatMain)
	//验证权限
	engine.POST("/check_auth",middleware.JwtApiMiddleware, controller.MainCheckAuth)
	//前后聊天
	engine.GET("/chat_server", controller.NewChatServer)
	//获取消息
	engine.GET("/messages", controller.GetVisitorMessage)
	//发送单条消息
	engine.POST("/message",controller.SendMessage)
	//获取未读消息数
	engine.GET("/message_status",controller.GetVisitorMessage)
	//设置消息已读
	engine.POST("/message_status",controller.GetVisitorMessage)

	//获取客服信息
	engine.GET("/kefuinfo",middleware.JwtApiMiddleware, controller.GetKefuInfo)
	engine.GET("/kefuinfo_setting",middleware.JwtApiMiddleware, controller.GetKefuInfoSetting)
	engine.POST("/kefuinfo",middleware.JwtApiMiddleware,middleware.RbacAuth, controller.PostKefuInfo)
	engine.DELETE("/kefuinfo",middleware.JwtApiMiddleware,middleware.RbacAuth, controller.DeleteKefuInfo)
	engine.GET("/kefulist",middleware.JwtApiMiddleware, controller.GetKefuList)
	//设置页
	engine.GET("/setting", tmpl.PageSetting)
	//设置mysql
	engine.GET("/setting_mysql", tmpl.PageSettingMysql)
	//角色列表
	engine.GET("/roles",middleware.JwtApiMiddleware, controller.GetRoleList)
	engine.GET("/roles_list", tmpl.PageRoleList)

	//网页部署
	engine.GET("/setting_deploy", tmpl.PageSettingDeploy)
	//邮箱列表
	engine.GET("/mail_list", tmpl.PageMailList)
	//邮件夹列表
	engine.GET("/folders", controller.GetFolders)

	engine.GET("/mysql",middleware.JwtApiMiddleware,middleware.RbacAuth,  controller.MysqlGetConf)
	engine.POST("/mysql",middleware.JwtApiMiddleware,middleware.RbacAuth,  controller.MysqlSetConf)
	engine.POST("/visitor",controller.PostVisitor)
	engine.GET("/visitor",middleware.JwtApiMiddleware, controller.GetVisitor)
	engine.GET("/visitors",middleware.JwtApiMiddleware, controller.GetVisitors)
	engine.GET("/setting_kefu_list",tmpl.PageKefuList)

	//前台接口
	engine.GET("/notice",middleware.SetLanguage, controller.GetNotice)
	//前台引入js接口
	engine.GET("/webjs", tmpl.PageWebJs)
	//前台引入css接口
	engine.GET("/webcss", tmpl.PageWebCss)
	//文档服务
	docs.SwaggerInfo.Title = "GO-FLY接口文档"
	docs.SwaggerInfo.Description = "go-fly即时通讯web客服管理系统 , 测试账户:kefu2 测试密码:123 类型:kefu"
	docs.SwaggerInfo.Version = "0.0.7"
	//docs.SwaggerInfo.Host = "127.0.0.1:"+port
	docs.SwaggerInfo.Host = "gofly.sopans.com"
	docs.SwaggerInfo.BasePath = "/"
	//docs.SwaggerInfo.Schemes = []string{"http"}
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logFile, _ := os.OpenFile("./fatal.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	tools.RedirectStderr(logFile)
	engine.Run(baseServer)
}
