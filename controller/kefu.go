package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
)

func GetKefuInfo(c *gin.Context){
	 kefuId, _ := c.Get("kefu_id")
	 user:=models.FindUserById(kefuId)
	 info:=make(map[string]interface{})
	 info["name"]=user.Nickname
	info["id"]=user.Name
	info["avator"]=user.Avator
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result":info,
	})
}
func PostKefuInfo(c *gin.Context){
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result":"aaa",
	})
}
func GetKefuList(c *gin.Context){
	users:=models.FindUsers()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"result":users,
	})
}
