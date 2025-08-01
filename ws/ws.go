package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"goflylivechat/models"
	"goflylivechat/tools"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type User struct {
	Conn       *websocket.Conn
	Name       string
	Id         string
	Avator     string
	To_id      string
	Role_id    string
	Mux        sync.Mutex
	UpdateTime time.Time
}
type Message struct {
	conn        *websocket.Conn
	context     *gin.Context
	content     []byte
	messageType int
	Mux         sync.Mutex
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
var KefuList = make(map[string]*User)
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
	return ""
}

// 定时给更新数据库状态
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

// 后端广播发送消息
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
			message.Mux.Lock()
			defer message.Mux.Unlock()
			conn.WriteMessage(websocket.TextMessage, str)
		case "inputing":
			data := typeMsg.Data.(map[string]interface{})
			from := data["from"].(string)
			to := data["to"].(string)
			//限流
			if tools.LimitFreqSingle("inputing:"+from, 1, 2) {
				OneKefuMessage(to, message.content)
			}
		}

	}
}
func UpdateVisitorUser(visitorId string, toId string) {
	guest, ok := ClientList[visitorId]
	if ok {
		guest.To_id = toId
	}
}
