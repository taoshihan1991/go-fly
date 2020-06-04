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
func ActionChatMain(w http.ResponseWriter, r *http.Request) {
	render := tmpl.NewRender(w)
	render.Display("chat_main", nil)
}

//获取在线用户
func ChatUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json;charset=utf-8;")
	result := make([]map[string]string, 0)
	for uid, _ := range clientList {
		userInfo := make(map[string]string)
		userInfo["uid"] = uid
		result = append(result, userInfo)
	}
	msg, _ := json.Marshal(tools.JsonListResult{
		JsonResult: tools.JsonResult{Code: 200, Msg: "获取成功"},
		Result:     result,
	})
	w.Write(msg)
}

//兼容之前的聊天服务
func ChatServer(w *websocket.Conn) {
	var error error
	for {
		//接受消息
		var receive string
		if error = websocket.Message.Receive(w, &receive); error != nil {
			log.Println("接受消息失败", error)
			break
		}
		log.Println("客户端:", receive)

		var loginMsg Message
		json.Unmarshal([]byte(receive), &loginMsg)

		switch loginMsg.Type.(string) {
		//用户上线
		case "userInit":
			//用户id对应的连接
			clientList[loginMsg.Uid.(string)] = w
			if len(kefuList)==0{
				websocket.Message.Send(w, "无客服在线")
			}else{
				//发送给客服通知
				for _, conn := range kefuList {
					result := make([]map[string]string, 0)
					for uid, _ := range clientList {
						userInfo := make(map[string]string)
						userInfo["uid"] = uid
						result = append(result, userInfo)
					}
					msg:=NoticeMessage{
						Type: "notice",
						Data:result,
					}
					str,_:=json.Marshal(msg);sendStr:=string(str)
					websocket.Message.Send(conn,sendStr)
				}
			}
		//客服上线
		case "kfOnline":
			//客服id对应的连接
			kefuList[loginMsg.Uid.(string)] = w
			sendMsg := ChatKfMessage{
				Message_type: "kfOnline",
				Data: KfMessageData{
					Kf_name: loginMsg.Name,
					Avatar:  loginMsg.Avatar,
					Kf_id:   loginMsg.Uid,
					Time:    time.Now().Format("2006-01-02 15:04:05"),
					Content: "客服上线",
				},
			}
			jsonStrByte, _ := json.Marshal(sendMsg)
			//发送给客户
			log.Println("发送给客户",clientList,string(jsonStrByte))
			for _, conn := range clientList {
				websocket.Message.Send(conn, string(jsonStrByte))
			}
			//发送给客服通知
			result := make([]map[string]string, 0)
			for uid, _ := range clientList {
				userInfo := make(map[string]string)
				userInfo["uid"] = uid
				result = append(result, userInfo)
			}
			msg:=NoticeMessage{
				Type: "notice",
				Data:result,
			}
			str,_:=json.Marshal(msg);sendStr:=string(str)
			websocket.Message.Send(w,sendStr)
		case "chatMessage":

		}
	}
}
//客户登陆和客服登陆发送的消息
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
type NoticeMessage struct {
	Type interface{} `json:"type"`
	Data interface{} `json:"data"`
}

var clientList = make(map[string]*websocket.Conn)
var kefuList = make(map[string]*websocket.Conn)