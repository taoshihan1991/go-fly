package main

import (
	"github.com/taoshihan1991/imaptool/controller"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)
func main() {
	baseServer:="127.0.0.1:8080"
	log.Println("start server...\r\ngo：http://"+baseServer)
	mux:=&http.ServeMux{}
	//根路径
	mux.HandleFunc("/", controller.ActionIndex)
	//邮件夹
	mux.HandleFunc("/list", controller.ActionFolder)
	//登陆界面
	mux.HandleFunc("/login", controller.ActionLogin)
	//验证接口
	mux.HandleFunc("/check", controller.LoginCheck)
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
	//聊天界面
	mux.HandleFunc("/chat_main", controller.ActionChatMain)
	//新邮件提醒服务
	mux.HandleFunc("/push_mail", controller.PushMailServer)
	//聊天界面
	mux.HandleFunc("/chat_page",controller.ActionChatPage)
	//聊天服务
	mux.Handle("/chat_server",websocket.Handler(controller.ChatServer))
	//获取在线用户
	mux.HandleFunc("/chat_users", controller.ChatUsers)
	//设置mysql
	mux.HandleFunc("/setting_mysql",controller.ActionChatPage)
	//后台任务
	controller.TimerSessFile()
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
	s.ListenAndServe()
}
