package controller

import (
	"encoding/json"
	"github.com/taoshihan1991/imaptool/tmpl"
	"github.com/taoshihan1991/imaptool/tools"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)
//聊天主界面
func ActionChatMain(w http.ResponseWriter, r *http.Request){
	render:=tmpl.NewRender(w)
	render.Display("chat_main",nil)
}
//获取在线用户
func ChatUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "text/json;charset=utf-8;")
	result:=make([]map[string]string,0)
	for uid,_:=range clientList{
		userInfo:=make(map[string]string)
		userInfo["uid"]=uid
		result=append(result,userInfo)
	}
	msg, _ := json.Marshal(tools.JsonListResult{
		JsonResult: tools.JsonResult{Code: 200, Msg: "获取成功"},
		Result:     result,
	})
	w.Write(msg)
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