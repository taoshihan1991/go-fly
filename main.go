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
	"log"
)
var (
	port string
	GoflyConfig config.Config
)
func main() {
	//获取参数中的数据
	flag.StringVar(&port, "port", "8080", "监听端口号")
	flag.Parse()
	if flag.NFlag() < 1 {
		flag.PrintDefaults()
	}
	baseServer := "0.0.0.0:"+port
	log.Println("start server...\r\ngo：http://" + baseServer)
	engine := gin.Default()
	engine.LoadHTMLGlob("static/html/*")
	engine.Static("/static", "./static")
	//首页
	engine.GET("/", controller.Index)
	//登陆界面
	engine.GET("/login", tmpl.PageLogin)
	//咨询界面
	engine.GET("/chat_page", tmpl.PageChat)
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
	engine.GET("/messages",middleware.JwtApiMiddleware, controller.GetVisitorMessage)
	//获取客服信息
	engine.GET("/kefuinfo",middleware.JwtApiMiddleware, controller.GetKefuInfo)
	engine.GET("/kefuinfo_setting",middleware.JwtApiMiddleware, controller.GetKefuInfoSetting)
	engine.POST("/kefuinfo",middleware.JwtApiMiddleware,middleware.CasbinACL, controller.PostKefuInfo)
	engine.DELETE("/kefuinfo",middleware.JwtApiMiddleware,middleware.CasbinACL, controller.DeleteKefuInfo)
	engine.GET("/kefulist",middleware.JwtApiMiddleware, controller.GetKefuList)
	//设置页
	engine.GET("/setting", tmpl.PageSetting)
	//设置mysql
	engine.GET("/setting_mysql", tmpl.PageSettingMysql)
	//网页部署
	engine.GET("/setting_deploy", tmpl.PageSettingDeploy)

	engine.GET("/mysql",middleware.JwtApiMiddleware,middleware.CasbinACL,  controller.MysqlGetConf)
	engine.POST("/mysql",middleware.JwtApiMiddleware,middleware.CasbinACL,  controller.MysqlSetConf)
	engine.GET("/visitor",middleware.JwtApiMiddleware, controller.GetVisitor)
	engine.GET("/visitors",middleware.JwtApiMiddleware, controller.GetVisitors)
	engine.GET("/setting_kefu_list",tmpl.PageKefuList)

	//前台接口
	engine.GET("/notice", controller.GetNotice)
	//前台引入js接口
	engine.GET("/webjs", tmpl.PageWebJs)
	//前台引入css接口
	engine.GET("/webcss", tmpl.PageWebCss)
	//文档服务
	docs.SwaggerInfo.Title = "GO-FLY接口文档"
	docs.SwaggerInfo.Description = "go-fly即时通讯web客服管理系统"
	docs.SwaggerInfo.Version = "0.0.6"
	docs.SwaggerInfo.Host = "127.0.0.1:"+port
	docs.SwaggerInfo.BasePath = "/"
	//docs.SwaggerInfo.Schemes = []string{"http"}
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.Run(baseServer)
}
