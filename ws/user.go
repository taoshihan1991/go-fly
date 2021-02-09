package ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/models"
	"log"
)

func NewKefuServer(c *gin.Context) {
	kefuId, _ := c.Get("kefu_id")
	kefuInfo := models.FindUserById(kefuId)
	if kefuInfo.ID == 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
		return
	}

	//go kefuServerBackend()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	//获取GET参数,创建WS
	var kefu User
	kefu.Id = kefuInfo.Name
	kefu.Name = kefuInfo.Nickname
	kefu.Avator = kefuInfo.Avator
	kefu.Role_id = kefuInfo.RoleId
	kefu.Conn = conn
	AddKefuToList(&kefu)

	for {
		//接受消息
		var receive []byte
		messageType, receive, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			go SendPingToKefuClient()
			return
		}

		message <- &Message{
			conn:        conn,
			content:     receive,
			context:     c,
			messageType: messageType,
		}
	}
}
func AddKefuToList(kefu *User) {
	var newKefuConns = []*User{kefu}
	kefuConns := KefuList[kefu.Id]
	if kefuConns != nil {
		for _, otherKefu := range kefuConns {
			msg := TypeMessage{
				Type: "many pong",
			}
			str, _ := json.Marshal(msg)
			err := otherKefu.Conn.WriteMessage(websocket.TextMessage, str)
			if err == nil {
				newKefuConns = append(newKefuConns, otherKefu)
			}
		}
	}
	log.Println("xxxxxxxxxxxxxxxxxxxxxxxx", newKefuConns)
	KefuList[kefu.Id] = newKefuConns
}

//后端广播发送消息
func kefuServerBackend() {
	for {
		message := <-message
		var typeMsg TypeMessage
		json.Unmarshal(message.content, &typeMsg)
		conn := message.conn
		if typeMsg.Type == nil || typeMsg.Data == nil {
			continue
		}
		msgType := typeMsg.Type.(string)
		log.Println("客户端:", string(message.content))

		switch msgType {
		//心跳
		case "ping":
			msg := TypeMessage{
				Type: "pong",
			}
			str, _ := json.Marshal(msg)
			Mux.Lock()
			conn.WriteMessage(websocket.TextMessage, str)
			Mux.Unlock()
		}

	}
}

//给超管发消息
func SuperAdminMessage(str []byte) {
	return
	//给超管发
	for _, kefuUsers := range KefuList {
		for _, kefuUser := range kefuUsers {
			if kefuUser.Role_id == "2" {
				kefuUser.Conn.WriteMessage(websocket.TextMessage, str)
			}
		}
	}
}

//给指定客服发消息
func OneKefuMessage(toId string, str []byte) {
	//新版
	mKefuConns := KefuList[toId]
	if mKefuConns != nil {
		for _, kefu := range mKefuConns {
			kefu.Conn.WriteMessage(websocket.TextMessage, str)
		}
	}
	SuperAdminMessage(str)
}

//给客服客户端发送消息判断客户端是否在线
func SendPingToKefuClient() {
	msg := TypeMessage{
		Type: "many pong",
	}
	str, _ := json.Marshal(msg)
	for kefuId, kfConns := range KefuList {
		var newKefuConns = []*User{}
		for _, kefuConn := range kfConns {
			if kefuConn == nil {
				continue
			}
			err := kefuConn.Conn.WriteMessage(websocket.TextMessage, str)
			if err == nil {
				newKefuConns = append(newKefuConns, kefuConn)
			}
		}
		if len(newKefuConns) > 0 {
			KefuList[kefuId] = newKefuConns
		} else {
			delete(KefuList, kefuId)
		}
	}
}
