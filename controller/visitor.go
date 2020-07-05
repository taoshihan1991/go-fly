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
func GetVisitorMessage(c *gin.Context) {
	visitorId:=c.Query("visitorId")
	messages:=models.FindMessageByVisitorId(visitorId)
	result:=make([]map[string]interface{},0)
	for _,message:=range messages{
		item:=make(map[string]interface{})
		var visitor models.Visitor
		var kefu models.User
		if visitor.Name=="" || kefu.Name==""{
			kefu=models.FindUser(message.KefuId)
			visitor=models.FindVisitorByVistorId(message.VisitorId)
		}
		item["time"]=message.CreatedAt
		item["content"]=message.Content
		item["mes_type"]=message.MesType
		item["visitor_name"]=visitor.Name
		item["visitor_avator"]=visitor.Avator
		item["kefu_name"]=kefu.Nickname
		item["kefu_avator"]=kefu.Avator
		result=append(result,item)

	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result":result,
	})
}
