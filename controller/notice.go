package controller

import (
	"github.com/gin-gonic/gin"
	"goflylivechat/models"
)

func GetNotice(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	user := models.FindUser(kefuId)
	if user.ID == 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "user not found",
		})
		return
	}
	welcomeMessage := models.FindConfigByUserId(user.Name, "WelcomeMessage")
	offlineMessage := models.FindConfigByUserId(user.Name, "OfflineMessage")
	allNotice := models.FindConfigByUserId(user.Name, "AllNotice")
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result": gin.H{
			"welcome":   welcomeMessage.ConfValue,
			"offline":   offlineMessage.ConfValue,
			"avatar":    user.Avator,
			"nickname":  user.Nickname,
			"allNotice": allNotice.ConfValue,
		},
	})
}
