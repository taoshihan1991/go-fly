package cmd

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/taoshihan1991/imaptool/controller"
	"github.com/taoshihan1991/imaptool/middleware"
	"github.com/taoshihan1991/imaptool/router"
	"github.com/taoshihan1991/imaptool/tools"
	"github.com/zh-five/xdaemon"
	"log"
	"os"
)

var (
	Port   string
	daemon bool
)
var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "example:go-fly server -p 8081",
	Example: "go-fly server -c config/",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	serverCmd.PersistentFlags().StringVarP(&Port, "port", "p", "8081", "监听端口号")
	serverCmd.PersistentFlags().BoolVarP(&daemon, "daemon", "d", false, "是否为守护进程模式")
}
func run() {
	if daemon == true {
		logFilePath := ""
		if dir, err := os.Getwd(); err == nil {
			logFilePath = dir + "/logs/"
		}
		_, err := os.Stat(logFilePath)
		if os.IsNotExist(err) {
			if err := os.MkdirAll(logFilePath, 0777); err != nil {
				log.Println(err.Error())
			}
		}
		d := xdaemon.NewDaemon(logFilePath + "go-fly.log")
		d.MaxCount = 5
		d.Run()
	}

	baseServer := "0.0.0.0:" + Port
	controller.Port = Port
	log.Println("start server...\r\ngo：http://" + baseServer)
	tools.Logger().Println("start server...\r\ngo：http://" + baseServer)

	engine := gin.Default()
	engine.LoadHTMLGlob("static/html/*")
	engine.Static("/static", "./static")
	engine.Use(tools.Session("gofly"))
	engine.Use(middleware.CrossSite)
	//性能监控
	pprof.Register(engine)

	//记录日志
	engine.Use(middleware.NewMidLogger())
	router.InitViewRouter(engine)
	router.InitApiRouter(engine)

	//logFile, _ := os.OpenFile("./fatal.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	//tools.RedirectStderr(logFile)

	//tcp服务
	//go controller.NewTcpServer(tcpBaseServer)
	engine.Run(baseServer)
}
