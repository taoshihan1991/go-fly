package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/config"
	"github.com/taoshihan1991/imaptool/controller"
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
	engine.GET("/mysql",middleware.JwtApiMiddleware,middleware.CasbinACL,  controller.MysqlGetConf)
	engine.POST("/mysql",middleware.JwtApiMiddleware,middleware.CasbinACL,  controller.MysqlSetConf)
	engine.GET("/visitor",middleware.JwtApiMiddleware, controller.GetVisitor)
	engine.GET("/visitors",middleware.JwtApiMiddleware, controller.GetVisitors)
	engine.GET("/setting_kefu_list",tmpl.PageKefuList)

	//前台接口
	engine.GET("/notice", controller.GetNotice)
	//配置文件
	engine.Run(baseServer)
}
