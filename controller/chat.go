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
		userInfo["username"]=clientNameList[uid]
		result = append(result, userInfo)
	}
	msg, _ := json.Marshal(tools.JsonListResult{
		JsonResult: tools.JsonResult{Code: 200, Msg: "获取成功"},
		Result:     result,
	})
	w.Write(msg)
}
type NoticeMessage struct {
	Type interface{} `json:"type"`
	Data interface{} `json:"data"`
}
type TypeMessage struct {
	Type interface{} `json:"type"`
	Data interface{} `json:"data"`
}
type KfMessage struct {
	Kf_name string `json:"kf_name"`
	Avatar  string `json:"avatar"`
	Kf_id   string `json:"kf_id"`
	Kf_group string `json:"kf_group"`
	Time    string `json:"time"`
	Guest_id string `json:"guest_id"`
	Content string `json:"content"`
}
type UserMessage struct {
	From_avatar string `json:"from_avatar"`
	From_id     string `json:"from_id"`
	From_name   string `json:"from_name"`
	To_id       string `json:"to_id"`
	To_name     string `json:"to_name"`
	Time     string `json:"time"`
	Content     string `json:"content"`
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

		var typeMsg TypeMessage
		var kfMsg KfMessage
		var userMsg UserMessage
		json.Unmarshal([]byte(receive), &typeMsg)
		if typeMsg.Type==nil||typeMsg.Data==nil{
			break
		}
		msgType:=typeMsg.Type.(string)
		msgData, _ :=json.Marshal(typeMsg.Data)

		switch msgType {
		//用户上线
		case "userInit":
			json.Unmarshal(msgData,&userMsg)
			//用户id对应的连接
			clientList[userMsg.From_id] = w
			clientNameList[userMsg.From_id]=userMsg.From_name
			SendUserAllNotice(userMsg.From_id)
		//客服上线
		case "kfOnline":
			json.Unmarshal(msgData,&kfMsg)
			//客服id对应的连接
			kefuList[kfMsg.Kf_id] = w
			//发送给客户
			if len(clientList)==0{
				break
			}
			for _, conn := range clientList {
				SendKefuOnline(kfMsg,conn)
			}
			//发送给客服通知
			//SendOnekfuAllNotice(w)
		//客服接手
		case "kfConnect":
			json.Unmarshal(msgData,&kfMsg)
			SendKefuOnline(kfMsg,clientList[kfMsg.Guest_id])
		case "kfChatMessage":
			json.Unmarshal(msgData,&kfMsg)
			conn:=clientList[kfMsg.Guest_id]
			msg:=NoticeMessage{
				Type: "kfChatMessage",
				Data:KfMessage{
					Kf_name:  kfMsg.Kf_name,
					Avatar:   kfMsg.Avatar,
					Kf_id:    kfMsg.Kf_id,
					Time:     time.Now().Format("2006-01-02 15:04:05"),
					Guest_id: kfMsg.Guest_id,
					Content:  kfMsg.Content,
				},
			}
			str,_:=json.Marshal(msg);sendStr:=string(str)
			websocket.Message.Send(conn,sendStr)
		case "chatMessage":
			json.Unmarshal(msgData,&userMsg)
			conn:=kefuList[userMsg.To_id]
			msg:=NoticeMessage{
				Type: "chatMessage",
				Data:UserMessage{
					From_avatar: userMsg.From_avatar,
					From_id:     userMsg.From_id,
					From_name:   userMsg.From_name,
					To_id:       userMsg.To_id,
					To_name:     userMsg.To_name,
					Content:     userMsg.Content,
					Time:    time.Now().Format("2006-01-02 15:04:05"),
				},
			}
			str,_:=json.Marshal(msg);sendStr:=string(str)
			websocket.Message.Send(conn,sendStr)
		}
	}
}
//发送给所有客服客户上线
func SendUserAllNotice(uid string){
	if len(kefuList)!=0{
		//发送给客服通知
		for _, conn := range kefuList {
			result := make([]map[string]string, 0)
			userInfo := make(map[string]string)
			userInfo["uid"] = uid
			userInfo["username"]=clientNameList[uid]
			result = append(result, userInfo)
			msg:=NoticeMessage{
				Type: "notice",
				Data:result,
			}
			str,_:=json.Marshal(msg);sendStr:=string(str)
			websocket.Message.Send(conn,sendStr)
		}
	}
}
//发送给客户客服上线
func SendKefuOnline(kfMsg KfMessage,conn *websocket.Conn){
	sendMsg := TypeMessage{
		Type: "kfOnline",
		Data: KfMessage{
			Kf_name: kfMsg.Kf_name,
			Avatar:  kfMsg.Avatar,
			Kf_id:   kfMsg.Kf_id,
			Kf_group: kfMsg.Kf_group,
			Time:    time.Now().Format("2006-01-02 15:04:05"),
			Content: "客服上线",
		},
	}
	jsonStrByte, _ := json.Marshal(sendMsg)
	websocket.Message.Send(conn, string(jsonStrByte))
}
//发送给所有客服客户上线
func SendOnekfuAllNotice(conn *websocket.Conn){
	result := make([]map[string]string, 0)
	for uid, _ := range clientList {
		userInfo := make(map[string]string)
		userInfo["uid"] = uid
		userInfo["username"]=clientNameList[uid]
		result = append(result, userInfo)
	}
	msg:=NoticeMessage{
		Type: "notice",
		Data:result,
	}
	str,_:=json.Marshal(msg);sendStr:=string(str)
	websocket.Message.Send(conn,sendStr)
}
func SendUserChat(){}
func SendKefuChat(){}
var clientList = make(map[string]*websocket.Conn)
var clientNameList = make(map[string]string)
var kefuList = make(map[string]*websocket.Conn)