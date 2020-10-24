package ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/models"
	"log"
)

func NewVisitorServer(c *gin.Context) {
	go kefuServerBackend()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	//获取GET参数,创建WS
	vistorInfo := models.FindVisitorByVistorId(c.Query("visitor_id"))
	if vistorInfo.VisitorId == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "访客不存在",
		})
		return
	}
	user := &User{
		Conn:   conn,
		Name:   vistorInfo.Name,
		Avator: vistorInfo.Avator,
		Id:     vistorInfo.VisitorId,
		To_id:  vistorInfo.ToId,
	}
	go SendServerJiang(vistorInfo.Name)
	AddVisitorToList(user)

	for {
		//接受消息
		var receive []byte
		messageType, receive, err := conn.ReadMessage()
		if err != nil {
			for uid, visitor := range ClientList {
				if visitor.Conn == conn {
					log.Println("删除用户", uid)
					delete(ClientList, uid)
					models.UpdateVisitorStatus(uid, 0)
					userInfo := make(map[string]string)
					userInfo["uid"] = uid
					userInfo["name"] = visitor.Name
					msg := TypeMessage{
						Type: "userOffline",
						Data: userInfo,
					}
					str, _ := json.Marshal(msg)
					kefuConns := KefuList[visitor.To_id]
					if kefuConns != nil {
						for _, kefuConn := range kefuConns {
							kefuConn.Conn.WriteMessage(websocket.TextMessage, str)
						}
					}
				}
			}
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
func AddVisitorToList(user *User) {
	//用户id对应的连接
	ClientList[user.Id] = user
	lastMessage := models.FindLastMessageByVisitorId(user.Id)
	userInfo := make(map[string]string)
	userInfo["uid"] = user.Id
	userInfo["username"] = user.Name
	userInfo["avator"] = user.Avator
	userInfo["last_message"] = lastMessage.Content
	if userInfo["last_message"] == "" {
		userInfo["last_message"] = "新访客"
	}
	msg := TypeMessage{
		Type: "userOnline",
		Data: userInfo,
	}
	str, _ := json.Marshal(msg)

	//新版
	mKefuConns := KefuList[user.To_id]
	if mKefuConns != nil {
		for _, kefu := range mKefuConns {
			kefu.Conn.WriteMessage(websocket.TextMessage, str)
		}
	}
}
