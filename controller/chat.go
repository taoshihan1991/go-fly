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
var message = make(chan *Message)

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
	go sendPingUpdateStatus()
	go singleBroadcaster()
	go sendPingOnlineUsers()
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
					models.UpdateVisitorStatus(uid,0)
					SendNoticeToAllKefu()
				}
			}
			time.Sleep(10 * time.Second)
		}

	}()
}
//定时给更新数据库状态
func sendPingUpdateStatus() {
	for {
		visitors:=models.FindVisitorsOnline()
		for _,visitor :=range visitors{
			_,ok:=clientList[visitor.VisitorId]
			if !ok{
				models.UpdateVisitorStatus(visitor.VisitorId,0)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
//定时推送当前在线用户
func sendPingOnlineUsers() {
	for {
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
		for _, kfConn := range kefuList {
			kfConn.WriteMessage(websocket.TextMessage,str)
		}
		time.Sleep(10 * time.Second)
	}
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


//后端广播发送消息
func singleBroadcaster(){
	for {
		message:=<-message
		log.Println("debug:",message)

		var typeMsg TypeMessage
		var clientMsg ClientMessage
		json.Unmarshal(message.content, &typeMsg)
		conn:=message.conn
		if typeMsg.Type == nil || typeMsg.Data == nil {
			continue
		}
		msgType := typeMsg.Type.(string)
		msgData, _ := json.Marshal(typeMsg.Data)
		switch msgType {
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
				continue
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
			visitor,ok := clientList[clientMsg.ToId]
			if visitor==nil||!ok{
				continue
			}
			SendKefuOnline(clientMsg, visitor.conn)
		case "kfChatMessage":
			json.Unmarshal(msgData, &clientMsg)
			guest,ok:=clientList[clientMsg.ToId]
			if guest==nil||!ok{
				continue
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
				continue
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



