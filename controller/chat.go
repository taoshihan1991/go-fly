package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/models"
	"log"
	"time"
)
type vistor struct{
	conn *websocket.Conn
	name string
	id string
	avator string
}
type Message struct{
	conn *websocket.Conn
	c *gin.Context
	content []byte
}
var clientList = make(map[string]*vistor)
var kefuList = make(map[string]*websocket.Conn)
var message = make(chan *Message, 10)

type TypeMessage struct {
	Type interface{} `json:"type"`
	Data interface{} `json:"data"`
}
type ClientMessage struct {
	Name  string `json:"name"`
	Avator   string `json:"avator"`
	Id    string `json:"id"`
	Group string `json:"group"`
	Time     string `json:"time"`
	ToId string `json:"to_id"`
	Content  string `json:"content"`
	City  string `json:"city"`
	ClientIp  string `json:"client_ip"`
}
//定时检测客户端是否在线
func init() {
	upgrader=websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	go singleBroadcaster()
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
		_, receive, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		recevString=string(receive)
		log.Println("客户端:", recevString)
		message<-&Message{
			conn:conn,
			content: receive,
			c:c,
		}
	}
}

//发送给客户客服上线
func SendKefuOnline(clientMsg ClientMessage, conn *websocket.Conn) {
	sendMsg := TypeMessage{
		Type: "kfOnline",
		Data: ClientMessage{
			Name:  clientMsg.Name,
			Avator:   clientMsg.Avator,
			Id:    clientMsg.Id,
			Group: clientMsg.Group,
			Time:     time.Now().Format("2006-01-02 15:04:05"),
			Content:  "客服上线",
		},
	}
	jsonStrByte, _ := json.Marshal(sendMsg)
	conn.WriteMessage(websocket.TextMessage,jsonStrByte)
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
				err := user.conn.WriteMessage(websocket.TextMessage,str)
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
			conn.WriteMessage(websocket.TextMessage,str)
		}
	}
}

//获取当前的在线用户
func getOnlineUser(w *websocket.Conn) {
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
	w.WriteMessage(websocket.TextMessage,str)
}
//后端广播发送消息
func singleBroadcaster(){
	for {
		message:=<-message
		var typeMsg TypeMessage
		var clientMsg ClientMessage
		json.Unmarshal(message.content, &typeMsg)
		conn:=message.conn
		if typeMsg.Type == nil || typeMsg.Data == nil {
			break
		}
		msgType := typeMsg.Type.(string)
		msgData, _ := json.Marshal(typeMsg.Data)
		switch msgType {
		//获取当前在线的所有用户
		case "getOnlineUsers":
			getOnlineUser(conn)
		//用户上线
		case "userInit":
			json.Unmarshal(msgData, &clientMsg)
			//用户id对应的连接
			user:=&vistor{
				conn:conn,
				name: clientMsg.Name,
				avator: clientMsg.Avator,
				id:clientMsg.Id,
			}
			clientList[clientMsg.Id] = user
			//插入数据表
			models.CreateVisitor(clientMsg.Name,clientMsg.Avator,message.c.ClientIP(),clientMsg.ToId,clientMsg.Id,message.c.Request.Referer(),clientMsg.City,clientMsg.ClientIp)
			SendNoticeToAllKefu()
		//客服上线
		case "kfOnline":
			json.Unmarshal(msgData, &clientMsg)
			//客服id对应的连接
			kefuList[clientMsg.Id] = conn
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
			json.Unmarshal(msgData, &clientMsg)
			kefuList[clientMsg.Id] = conn
			SendKefuOnline(clientMsg, clientList[clientMsg.ToId].conn)
		case "kfChatMessage":
			json.Unmarshal(msgData, &clientMsg)
			guest,ok:=clientList[clientMsg.ToId]
			if guest==nil||!ok{
				return
			}
			conn := guest.conn

			msg := TypeMessage{
				Type: "kfChatMessage",
				Data: ClientMessage{
					Name:  clientMsg.Name,
					Avator:   clientMsg.Avator,
					Id:    clientMsg.Id,
					Time:     time.Now().Format("2006-01-02 15:04:05"),
					ToId: clientMsg.ToId,
					Content:  clientMsg.Content,
				},
			}
			str, _ := json.Marshal(msg)
			conn.WriteMessage(websocket.TextMessage,str)
		case "chatMessage":
			json.Unmarshal(msgData, &clientMsg)
			conn,ok := kefuList[clientMsg.ToId]
			if conn==nil||!ok{
				return
			}
			msg := TypeMessage{
				Type: "chatMessage",
				Data: ClientMessage{
					Avator: clientMsg.Avator,
					Id:     clientMsg.Id,
					Name:   clientMsg.Name,
					ToId:       clientMsg.ToId,
					Content:     clientMsg.Content,
					Time:        time.Now().Format("2006-01-02 15:04:05"),
				},
			}
			str, _ := json.Marshal(msg)
			conn.WriteMessage(websocket.TextMessage,str)
		}
	}
}



