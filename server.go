package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/controller"
	"github.com/taoshihan1991/imaptool/middleware"
	"github.com/taoshihan1991/imaptool/tmpl"
	"log"
	"net/http"
	"time"
)
var (
	port string
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
	//获取客服信息
	engine.GET("/kefuinfo",middleware.JwtApiMiddleware, controller.GetKefuInfo)
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
	//------------------old code-----------------------------
	mux := &http.ServeMux{}
	//根路径
	mux.HandleFunc("/", controller.ActionIndex)
	//邮件夹
	mux.HandleFunc("/list", controller.ActionFolder)
	//邮件夹接口
	mux.HandleFunc("/folders", controller.FoldersList)
	//新邮件夹接口
	mux.HandleFunc("/folder_dirs", controller.FolderDir)
	//邮件接口
	mux.HandleFunc("/mail", controller.FolderMail)
	//详情界面
	mux.HandleFunc("/view", controller.ActionDetail)
	//写信界面
	mux.HandleFunc("/write", controller.ActionWrite)
	//框架界面
	mux.HandleFunc("/main", controller.ActionMain)
	//设置界面
	mux.HandleFunc("/setting", controller.ActionSetting)
	//设置账户接口
	mux.HandleFunc("/setting_account", controller.SettingAccount)
	//发送邮件接口
	mux.HandleFunc("/send", controller.FolderSend)
	//新邮件提醒服务
	mux.HandleFunc("/push_mail", controller.PushMailServer)
	//mux.Handle("/chat_server", websocket.Handler(controller.ChatServer))
	//后台任务
	//监听端口
	//http.ListenAndServe(":8080", nil)
	//var myHandler http.Handler
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//---------------old code end------------------
	engine.Run(baseServer)
	s.ListenAndServe()
}
