package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/models"
	"log"
	"sort"
	"time"
)
type vistor struct{
	conn *websocket.Conn
	name string
	id string
	avator string
	to_id string
}
type Message struct{
	conn *websocket.Conn
	c *gin.Context
	content []byte
	messageType int
}
var clientList = make(map[string]*vistor)
var kefuList = make(map[string][]*websocket.Conn)
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
	Refer string `json:"refer"`
}
//定时检测客户端是否在线
func init() {
	upgrader=websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	//go sendPingUpdateStatus()
	go singleBroadcaster()
	//go sendPingOnlineUsers()
	//sendPingToClient()
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
			for uid,visitor :=range clientList{
				if visitor.conn==conn{
					log.Println("删除用户",uid)
					delete(clientList,uid)
					models.UpdateVisitorStatus(uid,0)
					userInfo := make(map[string]string)
					userInfo["uid"] = uid
					userInfo["name"] = visitor.name
					msg := TypeMessage{
						Type: "userOffline",
						Data: userInfo,
					}
					str, _ := json.Marshal(msg)
					kefuConns:=kefuList[visitor.to_id]
					if kefuConns!=nil{
						for _,kefuConn:=range kefuConns{
							kefuConn.WriteMessage(websocket.TextMessage,str)
						}
					}
					sendPingOnlineUsers()
				}
			}
			log.Println(err)
			return
		}
		recevString=string(receive)
		log.Println("客户端:", recevString)
		message<-&Message{
			conn:conn,
			content: receive,
			c:c,
			messageType:messageType,
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
			str, _ := json.Marshal(msg)
			for uid, user := range clientList {
				err := user.conn.WriteMessage(websocket.TextMessage,str)
				if err != nil {
					delete(clientList, uid)
					models.UpdateVisitorStatus(uid,0)
				}
			}
			for kefuId, kfConns := range kefuList {

				var newkfConns =make([]*websocket.Conn,0)
				for _,kefuConn:=range kfConns{
					if(kefuConn==nil){
						continue
					}
					err:=kefuConn.WriteMessage(websocket.TextMessage,str)
					if err == nil {
						newkfConns=append(newkfConns,kefuConn)
					}
				}
				if newkfConns == nil {
					delete(kefuList, kefuId)
				}else{
					kefuList[kefuId]=newkfConns
				}
			}
			time.Sleep(15 * time.Second)
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
		time.Sleep(20 * time.Second)
	}
}
//定时推送当前在线用户
func sendPingOnlineUsers() {
	var visitorIds []string
	for visitorId, _ := range clientList {
		visitorIds=append(visitorIds,visitorId)
	}
	sort.Strings(visitorIds)

	for kefuId, kfConns := range kefuList {

		result := make([]map[string]string, 0)
		for _,visitorId:=range visitorIds{
			user:=clientList[visitorId]
			userInfo := make(map[string]string)
			userInfo["uid"] = user.id
			userInfo["username"] = user.name
			userInfo["avator"] = user.avator
			if user.to_id==kefuId{
				result = append(result, userInfo)
			}
		}
		msg := TypeMessage{
			Type: "allUsers",
			Data: result,
		}
		str, _ := json.Marshal(msg)
		var newkfConns =make([]*websocket.Conn,0)
		for _,kefuConn:=range kfConns{
			err:=kefuConn.WriteMessage(websocket.TextMessage,str)
			if err == nil {
				newkfConns=append(newkfConns,kefuConn)
			}
		}
		if len(newkfConns) == 0 {
			delete(kefuList, kefuId)
		}else{
			kefuList[kefuId]=newkfConns
		}
	}
}

//后端广播发送消息
func singleBroadcaster(){
	for {
		message:=<-message
		//log.Println("debug:",message)

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
				to_id:clientMsg.ToId,
			}
			clientList[clientMsg.Id] = user
			//插入数据表
			models.CreateVisitor(clientMsg.Name,clientMsg.Avator,message.c.ClientIP(),clientMsg.ToId,clientMsg.Id,clientMsg.Refer,clientMsg.City,clientMsg.ClientIp)
			userInfo := make(map[string]string)
			userInfo["uid"] = user.id
			userInfo["username"] = user.name
			userInfo["avator"] = user.avator
			msg := TypeMessage{
				Type: "userOnline",
				Data: userInfo,
			}
			str, _ := json.Marshal(msg)
			kefuConns:=kefuList[user.to_id]
			if kefuConns!=nil{
				for k,kefuConn:=range kefuConns{
					log.Println(k,"xxxxxxxx")
					kefuConn.WriteMessage(websocket.TextMessage,str)
				}
			}
			//客户上线发微信通知
			go SendServerJiang(userInfo["username"])
			sendPingOnlineUsers()
		//客服上线
		case "kfOnline":
			json.Unmarshal(msgData, &clientMsg)
			//客服id对应的连接
			var newKefuConns =[]*websocket.Conn{conn}
			kefuConns:=kefuList[clientMsg.Id]
			if kefuConns!=nil{
				newKefuConns=append(newKefuConns,kefuConns...)
			}
			log.Println(newKefuConns)
			kefuList[clientMsg.Id] = newKefuConns
			//发送给客户
			if len(clientList) == 0 {
				continue
			}
			sendPingOnlineUsers()
		//客服接手
		case "kfConnect":
			json.Unmarshal(msgData, &clientMsg)
			visitor,ok := clientList[clientMsg.ToId]
			if visitor==nil||!ok{
				continue
			}
			SendKefuOnline(clientMsg, visitor.conn)
		//心跳
		case "ping":
			msg := TypeMessage{
				Type: "pong",
			}
			str, _ := json.Marshal(msg)
			conn.WriteMessage(websocket.TextMessage,str)
		}

	}
}



