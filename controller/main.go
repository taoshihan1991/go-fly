package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/ws"
)

func MainCheckAuth(c *gin.Context) {
	id, _ := c.Get("kefu_id")
	userinfo := models.FindUserRole("user.avator,user.name,user.id, role.name role_name", id)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "验证成功",
		"result": gin.H{
			"avator":    userinfo.Avator,
			"name":      userinfo.Name,
			"role_name": userinfo.RoleName,
		},
	})
}
func GetStatistics(c *gin.Context) {
	visitors := models.CountVisitors()
	message := models.CountMessage()
	session := len(ws.ClientList)
	kefuNum := 0
	for _, kefus := range ws.KefuList {
		kefuNum += len(kefus)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result": gin.H{
			"visitors": visitors,
			"message":  message,
			"session":  session + kefuNum,
		},
	})
}
