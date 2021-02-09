package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type User struct {
	Conn    *websocket.Conn
	Name    string
	Id      string
	Avator  string
	To_id   string
	Role_id string
}
type Message struct {
	conn        *websocket.Conn
	context     *gin.Context
	content     []byte
	messageType int
}
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
	IsKefu    string `json:"is_kefu"`
}

var ClientList = make(map[string]*User)
var KefuList = make(map[string][]*User)
var message = make(chan *Message, 10)
var upgrader = websocket.Upgrader{}
var Mux sync.RWMutex

func init() {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	go UpdateVisitorStatusCron()
}
func SendServerJiang(title string, content string, domain string) string {
	noticeServerJiang, err := strconv.ParseBool(models.FindConfig("NoticeServerJiang"))
	serverJiangAPI := models.FindConfig("ServerJiangAPI")
	if err != nil || !noticeServerJiang || serverJiangAPI == "" {
		log.Println("do not notice serverjiang:", serverJiangAPI, noticeServerJiang)
		return ""
	}
	sendStr := fmt.Sprintf("%s%s", title, content)
	desp := title + ":" + content + "[登录](http://" + domain + "/main)"
	url := serverJiangAPI + "?text=" + sendStr + "&desp=" + desp
	//log.Println(url)
	res := tools.Get(url)
	return res
}
func SendFlyServerJiang(title string, content string, domain string) string {
	sendStr := fmt.Sprintf("%s%s", title, content)
	ip := tools.GetExternalIp()
	desp := content + ",内网:" + domain + ",外网:" + ip
	url := "https://sc.ftqq.com/SCU113707T98a3ef409d8a018c98ee7abe3278a7155f5f49e886742.send?text=" + sendStr + "&desp=" + desp
	//log.Println(url)
	res := tools.Get(url)
	return res
}

//定时给更新数据库状态
func UpdateVisitorStatusCron() {
	for {
		visitors := models.FindVisitorsOnline()
		for _, visitor := range visitors {
			if visitor.VisitorId == "" {
				continue
			}
			_, ok := ClientList[visitor.VisitorId]
			if !ok {
				models.UpdateVisitorStatus(visitor.VisitorId, 0)
			}
		}
		SendPingToKefuClient()
		time.Sleep(60 * time.Second)
	}
}

//后端广播发送消息
func WsServerBackend() {
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
		case "inputing":
			//msg := TypeMessage{
			//	Type: "pong",
			//}
			//str, _ := json.Marshal(msg)
			//var data SimpleMessage
			//json.Unmarshal([]byte(typeMsg.Data.(string)),&data)
			//log.Println(data)
			//from:=typeMsg.Data["from"]
			to := typeMsg.Data.(map[string]interface{})["to"].(string)
			//content:=typeMsg.Data["content"]
			log.Println(to)
			Mux.Lock()
			//conn.WriteMessage(websocket.TextMessage, str)
			OneKefuMessage(to, message.content)
			Mux.Unlock()
		}

	}
}
func UpdateVisitorUser(visitorId string, toId string) {
	guest, _ := ClientList[visitorId]
	guest.To_id = toId
}
