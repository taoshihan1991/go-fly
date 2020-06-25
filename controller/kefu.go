package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
)

func GetKefuInfo(c *gin.Context){
	 kefuId, _ := c.Get("kefu_id")
	 user:=models.FindUserById(kefuId)
	 info:=make(map[string]interface{})
	 info["kf_name"]=user.Nickname
	info["kf_id"]=user.Name
	info["avator"]=user.Avator
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"result":info,
	})
}
