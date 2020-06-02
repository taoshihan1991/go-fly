package main

import (
	"encoding/json"
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
	//聊天服务
	mux.Handle("/chat_server",websocket.Handler(ChatServer))
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
//兼容之前的聊天服务
func ChatServer(w *websocket.Conn){
	var error error
	for {
		//接受消息
		var receive string
		if error = websocket.Message.Receive(w, &receive); error != nil {
			log.Println("接受消息失败", error)
			break
		}
		message := Message{}
		err := json.Unmarshal([]byte(receive), &message)
		if err != nil {
			log.Println(err)
		}
		chat := ChatUserMessage{}
		json.Unmarshal([]byte(receive), &chat)
		log.Println("客户端:", message)
		kfMessageData := KfMessageData{
			Kf_name: "客服小美",
			Avatar:  "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=4217138672,2588039002&fm=26&gp=0.jpg",
			Kf_id:   "KF2",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
			Content: chat.Data.Content,
		}
		switch message.Type.(string) {
		//用户初始化
		case "userInit":
			clientList[message.Uid.(string)] = w
			sendMsg := ChatKfMessage{
				Message_type: "kf_online",
				Data:         kfMessageData,
			}
			jsonStrByte, _ := json.Marshal(sendMsg)
			log.Println("服务端:", string(jsonStrByte))
			websocket.Message.Send(w, string(jsonStrByte))
			//正常发送消息
		case "chatMessage":

			sendMsg := ChatKfMessage{
				Message_type: "chatMessage",
				Data:         kfMessageData,
			}
			jsonStrByte, _ := json.Marshal(sendMsg)
			log.Println("服务端:", string(jsonStrByte))
			websocket.Message.Send(w, string(jsonStrByte))
			//回应ping
		case "ping":

			sendMsg := PingMessage{
				Type: "pong",
			}
			jsonStrByte, _ := json.Marshal(sendMsg)
			log.Println("服务端:", string(jsonStrByte))
			websocket.Message.Send(w, string(jsonStrByte))
		}
	}
}

type Message struct {
	Type   interface{}
	Uid    interface{}
	Name   interface{}
	Avatar interface{}
	Group  interface{}
}
type KfMessageData struct {
	Kf_name interface{} `json:"kf_name"`
	Avatar  interface{} `json:"avatar"`
	Kf_id   interface{} `json:"kf_id"`
	Time    interface{} `json:"time"`
	Content interface{} `json:"content"`
}
type UserMessageData struct {
	From_avatar interface{} `json:"from_avatar"`
	From_id     interface{} `json:"from_id"`
	From_name   interface{} `json:"from_name"`
	To_id       interface{} `json:"to_id"`
	To_name     interface{} `json:"to_name"`
	Content     interface{} `json:"content"`
}
type ChatKfMessage struct {
	Message_type interface{}   `json:"message_type"`
	Data         KfMessageData `json:"data"`
}
type ChatUserMessage struct {
	Message_type interface{}     `json:"message_type"`
	Data         UserMessageData `json:"data"`
}
type PingMessage struct {
	Type interface{} `json:"type"`
}

var clientList =make(map[string]*websocket.Conn)