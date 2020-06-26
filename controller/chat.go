package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"time"
)
type vistor struct{
	conn *websocket.Conn
	name string
	id string
	avator string
}
var clientList = make(map[string]*vistor)
var kefuList = make(map[string]*websocket.Conn)

type TypeMessage struct {
	Type interface{} `json:"type"`
	Data interface{} `json:"data"`
}
type KfMessage struct {
	Kf_name  string `json:"kf_name"`
	Avatar   string `json:"avatar"`
	Kf_id    string `json:"kf_id"`
	Kf_group string `json:"kf_group"`
	Time     string `json:"time"`
	Guest_id string `json:"guest_id"`
	Content  string `json:"content"`
}
type UserMessage struct {
	From_avatar string `json:"from_avatar"`
	From_id     string `json:"from_id"`
	From_name   string `json:"from_name"`
	To_id       string `json:"to_id"`
	To_name     string `json:"to_name"`
	Time        string `json:"time"`
	Content     string `json:"content"`
}
//定时检测客户端是否在线
func init() {
	upgrader=websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	sendPingToClient()
}

func NewChatServer(c *gin.Context){
	conn,err:=upgrader.Upgrade(c.Writer,c.Request,nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	for {
		//接受消息
		var receive []byte
		var recevString string
		messageType, receive, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		recevString=string(receive)
		log.Println("客户端:", recevString,messageType)
		var typeMsg TypeMessage
		var kfMsg KfMessage
		var userMsg UserMessage
		json.Unmarshal(receive, &typeMsg)
		if typeMsg.Type == nil || typeMsg.Data == nil {
			break
		}
		msgType := typeMsg.Type.(string)
		msgData, _ := json.Marshal(typeMsg.Data)
		switch msgType {
		//获取当前在线的所有用户
		case "getOnlineUsers":
			getOnlineUser(conn,messageType)
		//用户上线
		case "userInit":
			json.Unmarshal(msgData, &userMsg)
			//用户id对应的连接
			user:=&vistor{
				conn:conn,
				name: userMsg.From_name,
				avator: userMsg.From_avatar,
				id:userMsg.From_id,
			}
			clientList[userMsg.From_id] = user
			SendNoticeToAllKefu()
		//客服上线
		case "kfOnline":
			json.Unmarshal(msgData, &kfMsg)
			//客服id对应的连接
			kefuList[kfMsg.Kf_id] = conn
			//发送给客户
			if len(clientList) == 0 {
				break
			}
			//for _, conn := range clientList {
			//	SendKefuOnline(kfMsg, conn)
			//}
			//发送给客服通知
			//SendOnekfuAllNotice(w)
		//客服接手
		case "kfConnect":
			json.Unmarshal(msgData, &kfMsg)
			kefuList[kfMsg.Kf_id] = conn
			SendKefuOnline(kfMsg, clientList[kfMsg.Guest_id].conn)
		case "kfChatMessage":
			json.Unmarshal(msgData, &kfMsg)
			conn := clientList[kfMsg.Guest_id].conn
			if kfMsg.Guest_id == "" || conn == nil {
				return
			}
			msg := TypeMessage{
				Type: "kfChatMessage",
				Data: KfMessage{
					Kf_name:  kfMsg.Kf_name,
					Avatar:   kfMsg.Avatar,
					Kf_id:    kfMsg.Kf_id,
					Time:     time.Now().Format("2006-01-02 15:04:05"),
					Guest_id: kfMsg.Guest_id,
					Content:  kfMsg.Content,
				},
			}
			str, _ := json.Marshal(msg)
			conn.WriteMessage(messageType,str)
		case "chatMessage":
			json.Unmarshal(msgData, &userMsg)
			conn := kefuList[userMsg.To_id]
			msg := TypeMessage{
				Type: "chatMessage",
				Data: UserMessage{
					From_avatar: userMsg.From_avatar,
					From_id:     userMsg.From_id,
					From_name:   userMsg.From_name,
					To_id:       userMsg.To_id,
					To_name:     userMsg.To_name,
					Content:     userMsg.Content,
					Time:        time.Now().Format("2006-01-02 15:04:05"),
				},
			}
			str, _ := json.Marshal(msg)
			conn.WriteMessage(messageType,str)
		}
	}
}

//发送给客户客服上线
func SendKefuOnline(kfMsg KfMessage, conn *websocket.Conn) {
	sendMsg := TypeMessage{
		Type: "kfOnline",
		Data: KfMessage{
			Kf_name:  kfMsg.Kf_name,
			Avatar:   kfMsg.Avatar,
			Kf_id:    kfMsg.Kf_id,
			Kf_group: kfMsg.Kf_group,
			Time:     time.Now().Format("2006-01-02 15:04:05"),
			Content:  "客服上线",
		},
	}
	jsonStrByte, _ := json.Marshal(sendMsg)
	conn.WriteMessage(1,jsonStrByte)
}

//发送给所有客服客户上线
func SendOnekfuAllNotice(conn *websocket.Conn) {
	result := make([]map[string]string, 0)
	for _, user := range clientList {
		userInfo := make(map[string]string)
		userInfo["uid"] = user.id
		userInfo["username"] = user.name
		userInfo["avator"] = user.avator
		result = append(result, userInfo)
	}
	msg := TypeMessage{
		Type: "notice",
		Data: result,
	}
	str, _ := json.Marshal(msg)
	conn.WriteMessage(1,str)
}

//定时给客户端发送消息判断客户端是否在线
func sendPingToClient() {
	msg := TypeMessage{
		Type: "ping",
	}
	go func() {
		for {
			log.Println("check online users...")
			str, _ := json.Marshal(msg)
			for uid, user := range clientList {
				err := user.conn.WriteMessage(1,str)
				if err != nil {
					delete(clientList, uid)
					SendNoticeToAllKefu()
				}
			}
			time.Sleep(10 * time.Second)
		}

	}()
}
func SendNoticeToAllKefu() {
	if len(kefuList) != 0 {
		//发送给客服通知
		for _, conn := range kefuList {
			msg := TypeMessage{
				Type: "notice",
			}
			str, _ := json.Marshal(msg)
			conn.WriteMessage(1,str)
		}
	}
}

//获取当前的在线用户
func getOnlineUser(w *websocket.Conn,messageType int) {
	result := make([]map[string]string, 0)
	for _, user := range clientList {
		userInfo := make(map[string]string)
		userInfo["uid"] = user.id
		userInfo["username"] = user.name
		userInfo["avator"] = user.avator
		result = append(result, userInfo)
	}
	msg := TypeMessage{
		Type: "getOnlineUsers",
		Data: result,
	}
	str, _ := json.Marshal(msg)
	w.WriteMessage(messageType,str)
}




