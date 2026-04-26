package ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"goflylivechat/models"
	"goflylivechat/tools"
	"log"
	"time"
)

func NewKefuServer(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	kefuInfo := models.FindUser(kefuName.(string))
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
	kefu.Conn = conn
	AddKefuToList(&kefu)

	for {
		//接受消息
		var receive []byte
		messageType, receive, err := conn.ReadMessage()
		if err != nil {
			log.Println("ws/user.go ", err)
			RemoveKefuFromList(kefuInfo.Name)
			conn.Close()
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
	Mux.Lock()
	defer Mux.Unlock()
	
	oldUser, ok := KefuList[kefu.Id]
	if oldUser != nil || ok {
		msg := TypeMessage{
			Type: "close",
			Data: kefu.Id,
		}
		str, _ := json.Marshal(msg)
		if err := oldUser.Conn.WriteMessage(websocket.TextMessage, str); err != nil {
			oldUser.Conn.Close()
		}
	}
	KefuList[kefu.Id] = kefu
	
	// 广播客服上线状态给所有该客服的访客
	broadcastKefuOnlineStatusLocked(kefu.Id, true)
}

func RemoveKefuFromList(kefuId string) {
	Mux.Lock()
	defer Mux.Unlock()
	
	_, ok := KefuList[kefuId]
	if ok {
		delete(KefuList, kefuId)
		// 广播客服离线状态给所有该客服的访客
		broadcastKefuOnlineStatusLocked(kefuId, false)
	}
}

// 给指定客服发消息
func OneKefuMessage(toId string, str []byte) {
	Mux.RLock()
	kefu, ok := KefuList[toId]
	Mux.RUnlock()
	
	if ok {
		log.Println("OneKefuMessage lock")
		kefu.Mux.Lock()
		defer kefu.Mux.Unlock()
		log.Println("OneKefuMessage unlock")
		error := kefu.Conn.WriteMessage(websocket.TextMessage, str)
		tools.Logger().Println("send_kefu_message", error, string(str))
	}
}

// 内部函数：在持有锁的情况下广播客服状态
func broadcastKefuOnlineStatusLocked(kefuId string, isOnline bool) {
	kefuInfo := models.FindUser(kefuId)
	msg := TypeMessage{
		Type: "kefu_status_change",
		Data: map[string]interface{}{
			"kefu_id":   kefuId,
			"is_online": isOnline,
			"nickname":  kefuInfo.Nickname,
			"avator":    kefuInfo.Avator,
		},
	}
	str, _ := json.Marshal(msg)
	
	for _, visitor := range ClientList {
		if visitor.To_id == kefuId && visitor.Conn != nil {
			visitor.Conn.WriteMessage(websocket.TextMessage, str)
		}
	}
}
func KefuMessage(visitorId, content string, kefuInfo models.User) {
	msg := TypeMessage{
		Type: "message",
		Data: ClientMessage{
			Name:    kefuInfo.Nickname,
			Avator:  kefuInfo.Avator,
			Id:      visitorId,
			Time:    time.Now().Format("2006-01-02 15:04:05"),
			ToId:    visitorId,
			Content: content,
			IsKefu:  "yes",
		},
	}
	str, _ := json.Marshal(msg)
	OneKefuMessage(kefuInfo.Name, str)
}

// 给客服客户端发送消息判断客户端是否在线
func SendPingToKefuClient() {
	msg := TypeMessage{
		Type: "many pong",
	}
	str, _ := json.Marshal(msg)
	
	// 收集需要删除的客服ID
	var toRemove []string
	
	Mux.RLock()
	// 先遍历获取所有需要检查的客服ID
	var kefuIds []string
	for kefuId := range KefuList {
		kefuIds = append(kefuIds, kefuId)
	}
	Mux.RUnlock()
	
	// 遍历检查每个客服连接（遍历过程中不持有锁，避免死锁）
	for _, kefuId := range kefuIds {
		Mux.RLock()
		kefu, ok := KefuList[kefuId]
		Mux.RUnlock()
		
		if !ok || kefu == nil {
			toRemove = append(toRemove, kefuId)
			continue
		}
		
		kefu.Mux.Lock()
		err := kefu.Conn.WriteMessage(websocket.TextMessage, str)
		kefu.Mux.Unlock()
		if err != nil {
			log.Println("定时发送ping给客服，失败", err.Error())
			toRemove = append(toRemove, kefuId)
		}
	}
	
	// 从列表中移除并广播离线状态
	for _, kefuId := range toRemove {
		RemoveKefuFromList(kefuId)
	}
}
