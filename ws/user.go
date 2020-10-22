package ws

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
	//获取GET参数,创建WS
	var kefu User
	kefu.Id = kefuInfo.Name
	kefu.Name = kefuInfo.Nickname
	kefu.Avator = kefuInfo.Avator
	kefu.Conn = conn
	AddKefuToList(&kefu)

	for {
		//接受消息
		var receive []byte
		messageType, receive, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
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
		for _, kefu := range kefuConns {
			msg := TypeMessage{
				Type: "pong",
			}
			str, _ := json.Marshal(msg)
			err := kefu.Conn.WriteMessage(websocket.TextMessage, str)
			if err != nil {
				newKefuConns = append(newKefuConns, kefu)
			}
		}
	}
	log.Println(newKefuConns)
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
