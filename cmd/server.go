package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/zh-five/xdaemon"
	"goflylivechat/middleware"
	"goflylivechat/router"
	"goflylivechat/tools"
	"goflylivechat/ws"
	"log"
	"os"
)

var (
	port   string
	daemon bool
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start HTTP service",
	Example: "gochat server -p 8082",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	serverCmd.PersistentFlags().StringVarP(&port, "port", "p", "8081", "Port to listen on")
	serverCmd.PersistentFlags().BoolVarP(&daemon, "daemon", "d", false, "Run as daemon process")
}

func run() {
	// Daemon mode setup
	if daemon {
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
		d := xdaemon.NewDaemon(logFilePath + "gofly.log")
		d.MaxCount = 10
		d.Run()
	}

	baseServer := "0.0.0.0:" + port
	log.Println("Starting server...\nURL: http://" + baseServer)
	tools.Logger().Println("Starting server...\nURL: http://" + baseServer)

	// Gin engine setup
	engine := gin.Default()
	engine.LoadHTMLGlob("static/templates/*")
	engine.Static("/assets", "./static")
	engine.Static("/static", "./static")
	engine.Use(tools.Session("gofly"))
	engine.Use(middleware.CrossSite)

	// Middlewares
	engine.Use(middleware.NewMidLogger())

	// Routers
	router.InitViewRouter(engine)
	router.InitApiRouter(engine)

	// Background services
	tools.NewLimitQueue()
	ws.CleanVisitorExpire()
	go ws.WsServerBackend()

	// Start server
	engine.Run(baseServer)
}
