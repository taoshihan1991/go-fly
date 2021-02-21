package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
	"net/http"
	"time"
)

func GetNotice(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	welcomes := models.FindWelcomesByKeyword(kefuId, "welcome")
	user := models.FindUser(kefuId)
	result := make([]gin.H, 0)
	for _, welcome := range welcomes {
		h := gin.H{
			"name":    user.Nickname,
			"avator":  user.Avator,
			"is_kefu": false,
			"content": welcome.Content,
			"time":    time.Now().Format("2006-01-02 15:04:05"),
		}
		result = append(result, h)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result": gin.H{
			"welcome":  result,
			"username": user.Nickname,
			"avatar":   user.Avator,
		},
	})
}
func GetNotices(c *gin.Context) {
	kefuId, _ := c.Get("kefu_name")
	welcomes := models.FindWelcomesByUserId(kefuId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": welcomes,
	})
}
func PostNotice(c *gin.Context) {
	kefuId, _ := c.Get("kefu_name")
	content := c.PostForm("content")
	models.CreateWelcome(fmt.Sprintf("%s", kefuId), content)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
func PostNoticeSave(c *gin.Context) {
	kefuId, _ := c.Get("kefu_name")
	content := c.PostForm("content")
	id := c.PostForm("id")
	models.UpdateWelcome(fmt.Sprintf("%s", kefuId), id, content)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
func DelNotice(c *gin.Context) {
	kefuId, _ := c.Get("kefu_name")
	id := c.Query("id")
	models.DeleteWelcome(kefuId, id)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
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
