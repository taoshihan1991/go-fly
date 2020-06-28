package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
)

func GetVisitor(c *gin.Context) {
	visitorId:=c.Query("visitorId")
	vistor:=models.FindVisitorByVistorId(visitorId)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result":vistor,
	})
}
func GetVisitors(c *gin.Context) {
	vistors:=models.FindVisitors()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result":vistors,
	})
}
