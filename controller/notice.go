package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/config"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
	"net/http"
	"time"
)
func GetNotice(c *gin.Context) {
	kefuId:=c.Query("kefu_id")
	lang,_:=c.Get("lang")
	language:=config.CreateLanguage(lang.(string))
	welcome:=models.FindWelcomeByUserId(kefuId)
	user:=models.FindUser(kefuId)
	var content string
	log.Println(welcome)
	if welcome.Content!=""{
		content=welcome.Content
	}else {
		content=language.Notice
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result":gin.H{
			"nickname":user.Nickname,
			"avator":user.Avator,
			"name":user.Name,
			"content":content,
			"time":time.Now().Format("2006-01-02 15:04:05"),
		},
	})
}
var upgrader = websocket.Upgrader{}
var oldFolders map[string]int

//推送新邮件到达
func PushMailServer(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		mailServer := tools.GetMailServerFromCookie(r)
		var msg []byte
		if mailServer == nil {
			msg, _ = json.Marshal(tools.JsonResult{Code: 400, Msg: "验证失败"})
			err = c.WriteMessage(mt, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		} else {
			folders := tools.GetMailNum(mailServer.Server, mailServer.Email, mailServer.Password)
			for name, num := range folders {
				if oldFolders[name] != num {
					result := make(map[string]interface{})
					result["folder_name"] = name
					result["new_num"] = num - oldFolders[name]
					msg, _ := json.Marshal(tools.JsonListResult{
						JsonResult: tools.JsonResult{Code: 200, Msg: "获取成功"},
						Result:     result,
					})
					c.WriteMessage(mt, msg)
				}
			}
			oldFolders = folders
		}
	}
}
