package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
	"net/http"
	"strconv"
	"sync"
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
}

var ClientList = make(map[string]*User)
var KefuList = make(map[string][]*User)
var message = make(chan *Message)
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

}
func SendServerJiang(content string) string {
	noticeServerJiang, err := strconv.ParseBool(models.FindConfig("NoticeServerJiang"))
	serverJiangAPI := models.FindConfig("ServerJiangAPI")
	if err != nil || !noticeServerJiang || serverJiangAPI == "" {
		log.Println("do not notice serverjiang:", serverJiangAPI, noticeServerJiang)
		return ""
	}
	sendStr := fmt.Sprintf("%s,访客来了", content)
	desp := "[登录](https://gofly.sopans.com/main)"
	url := serverJiangAPI + "?text=" + sendStr + "&desp=" + desp
	//log.Println(url)
	res := tools.Get(url)
	return res
}
