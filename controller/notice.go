package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/ws"
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
	status := "online"
	if kefus, ok := ws.KefuList[kefuId]; !ok {
		if len(kefus) <= 0 {
			status = "offline"
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result": gin.H{
			"welcome":  result,
			"username": user.Nickname,
			"avatar":   user.Avator,
			"status":   status,
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
