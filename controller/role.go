package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
)

func GetRoleList(c *gin.Context){
	roles:=models.FindRoles()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"result":roles,
	})
}
