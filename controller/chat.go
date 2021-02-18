package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/ws"
	"log"
	"sort"
	"time"
)

type vistor struct {
	conn   *websocket.Conn
	name   string
	id     string
	avator string
	to_id  string
}
type Message struct {
	conn        *websocket.Conn
	c           *gin.Context
	content     []byte
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
	Name      string `json:"name"`
	Avator    string `json:"avator"`
	Id        string `json:"id"`
	VisitorId string `json:"visitor_id"`
	Group     string `json:"group"`
	Time      string `json:"time"`
	ToId      string `json:"to_id"`
	Content   string `json:"content"`
	City      string `json:"city"`
	ClientIp  string `json:"client_ip"`
	Refer     string `json:"refer"`
}

//定时检测客户端是否在线
func init() {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	//go UpdateVisitorStatusCron()
	//go singleBroadcaster()
	//go sendPingOnlineUsers()
	//sendPingToClient()
}

func NewChatServer(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
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
			for uid, visitor := range clientList {
				if visitor.conn == conn {
					log.Println("删除用户", uid)
					delete(clientList, uid)
					models.UpdateVisitorStatus(uid, 0)
					userInfo := make(map[string]string)
					userInfo["uid"] = uid
					userInfo["name"] = visitor.name
					msg := TypeMessage{
						Type: "userOffline",
						Data: userInfo,
					}
					str, _ := json.Marshal(msg)
					kefuConns := kefuList[visitor.to_id]
					if kefuConns != nil {
						for _, kefuConn := range kefuConns {
							kefuConn.WriteMessage(websocket.TextMessage, str)
						}
					}
					//新版
					mKefuConns := ws.KefuList[visitor.to_id]
					if mKefuConns != nil {
						for _, kefu := range mKefuConns {
							kefu.Conn.WriteMessage(websocket.TextMessage, str)
						}
					}
					sendPingOnlineUsers()
				}
			}
			log.Println(err)
			return
		}
		recevString = string(receive)
		log.Println("客户端:", recevString)
		message <- &Message{
			conn:        conn,
			content:     receive,
			c:           c,
			messageType: messageType,
		}
	}
}

//发送给客户客服上线
func SendKefuOnline(clientMsg ClientMessage, conn *websocket.Conn) {
	sendMsg := TypeMessage{
		Type: "kfOnline",
		Data: ClientMessage{
			Name:    clientMsg.Name,
			Avator:  clientMsg.Avator,
			Id:      clientMsg.Id,
			Group:   clientMsg.Group,
			Time:    time.Now().Format("2006-01-02 15:04:05"),
			Content: "客服上线",
		},
	}
	jsonStrByte, _ := json.Marshal(sendMsg)
	conn.WriteMessage(websocket.TextMessage, jsonStrByte)
}

//发送通知
func SendNotice(msg string, conn *websocket.Conn) {
	sendMsg := TypeMessage{
		Type: "notice",
		Data: msg,
	}
	jsonStrByte, _ := json.Marshal(sendMsg)
	conn.WriteMessage(websocket.TextMessage, jsonStrByte)
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
				err := user.conn.WriteMessage(websocket.TextMessage, str)
				if err != nil {
					delete(clientList, uid)
					models.UpdateVisitorStatus(uid, 0)
				}
			}
			for kefuId, kfConns := range kefuList {

				var newkfConns = make([]*websocket.Conn, 0)
				for _, kefuConn := range kfConns {
					if kefuConn == nil {
						continue
					}
					err := kefuConn.WriteMessage(websocket.TextMessage, str)
					if err == nil {
						newkfConns = append(newkfConns, kefuConn)
					}
				}
				if newkfConns == nil {
					delete(kefuList, kefuId)
				} else {
					kefuList[kefuId] = newkfConns
				}
			}
			time.Sleep(15 * time.Second)
		}

	}()
}

//定时给更新数据库状态
func UpdateVisitorStatusCron() {
	for {
		visitors := models.FindVisitorsOnline()
		for _, visitor := range visitors {
			if visitor.VisitorId == "" {
				continue
			}
			_, ok := clientList[visitor.VisitorId]
			if !ok {
				models.UpdateVisitorStatus(visitor.VisitorId, 0)
			}
		}
		time.Sleep(60 * time.Second)
	}
}

//定时推送当前在线用户
func sendPingOnlineUsers() {
	var visitorIds []string
	for visitorId, _ := range clientList {
		visitorIds = append(visitorIds, visitorId)
	}
	sort.Strings(visitorIds)

	for kefuId, kfConns := range kefuList {

		result := make([]map[string]string, 0)
		for _, visitorId := range visitorIds {
			user := clientList[visitorId]
			userInfo := make(map[string]string)
			userInfo["uid"] = user.id
			userInfo["username"] = user.name
			userInfo["avator"] = user.avator
			if user.to_id == kefuId {
				result = append(result, userInfo)
			}
		}
		msg := TypeMessage{
			Type: "allUsers",
			Data: result,
		}
		str, _ := json.Marshal(msg)
		var newkfConns = make([]*websocket.Conn, 0)
		for _, kefuConn := range kfConns {
			err := kefuConn.WriteMessage(websocket.TextMessage, str)
			if err == nil {
				newkfConns = append(newkfConns, kefuConn)
			}
		}
		if len(newkfConns) == 0 {
			delete(kefuList, kefuId)
		} else {
			kefuList[kefuId] = newkfConns
		}
	}
}

//后端广播发送消息
//func singleBroadcaster() {
//	for {
//		message := <-message
//		//log.Println("debug:",message)
//
//		var typeMsg TypeMessage
//		var clientMsg ClientMessage
//		json.Unmarshal(message.content, &typeMsg)
//		conn := message.conn
//		if typeMsg.Type == nil || typeMsg.Data == nil {
//			continue
//		}
//		msgType := typeMsg.Type.(string)
//		msgData, _ := json.Marshal(typeMsg.Data)
//		switch msgType {
//		//用户上线
//		case "userInit":
//			json.Unmarshal(msgData, &clientMsg)
//			vistorInfo := models.FindVisitorByVistorId(clientMsg.VisitorId)
//			if vistorInfo.VisitorId == "" {
//				SendNotice("访客数据不存在", conn)
//				continue
//			}
//			//用户id对应的连接
//			user := &vistor{
//				conn:   conn,
//				name:   clientMsg.Name,
//				avator: clientMsg.Avator,
//				id:     clientMsg.VisitorId,
//				to_id:  clientMsg.ToId,
//			}
//			clientList[clientMsg.VisitorId] = user
//			//插入数据表
//			models.UpdateVisitor(clientMsg.VisitorId, 1, clientMsg.ClientIp, message.c.ClientIP(), clientMsg.Refer, "")
//			//models.CreateVisitor(clientMsg.Name,clientMsg.Avator,message.c.ClientIP(),clientMsg.ToId,clientMsg.VisitorId,clientMsg.Refer,clientMsg.City,clientMsg.ClientIp)
//			userInfo := make(map[string]string)
//			userInfo["uid"] = user.id
//			userInfo["username"] = user.name
//			userInfo["avator"] = user.avator
//			msg := TypeMessage{
//				Type: "userOnline",
//				Data: userInfo,
//			}
//			str, _ := json.Marshal(msg)
//
//			//新版
//			mKefuConns := ws.KefuList[user.to_id]
//			if mKefuConns != nil {
//				for _, kefu := range mKefuConns {
//					kefu.Conn.WriteMessage(websocket.TextMessage, str)
//				}
//			}
//
//			//兼容旧版
//			kefuConns := kefuList[user.to_id]
//			if kefuConns != nil {
//				for k, kefuConn := range kefuConns {
//					log.Println(k, "xxxxxxxx")
//					kefuConn.WriteMessage(websocket.TextMessage, str)
//				}
//			}
//
//			//客户上线发微信通知
//			go SendServerJiang(userInfo["username"])
//			sendPingOnlineUsers()
//		//客服上线
//		case "kfOnline":
//			json.Unmarshal(msgData, &clientMsg)
//			//客服id对应的连接
//			var newKefuConns = []*websocket.Conn{conn}
//			kefuConns := kefuList[clientMsg.Id]
//			if kefuConns != nil {
//				newKefuConns = append(newKefuConns, kefuConns...)
//			}
//			log.Println(newKefuConns)
//			kefuList[clientMsg.Id] = newKefuConns
//			//发送给客户
//			if len(clientList) == 0 {
//				continue
//			}
//			sendPingOnlineUsers()
//		//客服接手
//		case "kfConnect":
//			json.Unmarshal(msgData, &clientMsg)
//			visitor, ok := clientList[clientMsg.ToId]
//			if visitor == nil || !ok {
//				continue
//			}
//			SendKefuOnline(clientMsg, visitor.conn)
//		//心跳
//		case "ping":
//			msg := TypeMessage{
//				Type: "pong",
//			}
//			str, _ := json.Marshal(msg)
//			conn.WriteMessage(websocket.TextMessage, str)
//		}
//
//	}
//}
