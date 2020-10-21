package websocket

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/models"
	"log"
)

func NewKefuServer(c *gin.Context) {
	kefuId, _ := c.Get("kefu_name")
	kefuInfo := models.FindUser(kefuId.(string))
	if kefuInfo.ID == 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
		return
	}

	go kefuServerBackend()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	for {
		//接受消息
		var receive []byte
		messageType, receive, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		//获取GET参数,创建WS
		var kefu User
		kefu.id = kefuInfo.Name
		kefu.name = kefuInfo.Nickname
		kefu.avator = kefuInfo.Avator
		kefu.conn = conn
		AddKefuToList(&kefu)

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
	kefuConns := kefuList[kefu.id]
	if kefuConns != nil {
		newKefuConns = append(newKefuConns, kefuConns...)
	}
	kefuList[kefu.id] = newKefuConns
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
		log.Println("客户端:", msgType)

		switch msgType {
		//心跳
		case "ping":
			msg := TypeMessage{
				Type: "pong",
			}
			str, _ := json.Marshal(msg)
			conn.WriteMessage(websocket.TextMessage, str)
		}

	}
}
