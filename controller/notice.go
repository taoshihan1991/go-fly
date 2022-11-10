package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
)

func GetNotice(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	user := models.FindUser(kefuId)
	if user.ID==0{
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "user not found",
		})
		return
	}
	welcomeMessage:=models.FindConfig("WelcomeMessage")
	offlineMessage:=models.FindConfig("OfflineMessage")
	allNotice:=models.FindConfig("AllNotice")
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result": gin.H{
			"welcome":welcomeMessage,
			"offline":offlineMessage,
			"avatar":user.Avator,
			"nickname":user.Nickname,
			"allNotice":allNotice,
		},
	})
}
