package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
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
	name:=c.PostForm("name")
	password:=c.PostForm("password")
	avator:=c.PostForm("avator")
	models.CreateUser(name,tools.Md5(password),avator)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result":"",
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
func DeleteKefuInfo(c *gin.Context){
	kefuId := c.Query("id")
	log.Println(kefuId)
	models.DeleteUserById(kefuId)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
		"result":"",
	})
}
