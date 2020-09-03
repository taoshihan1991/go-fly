package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"strconv"
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
func GetKefuInfoSetting(c *gin.Context){
	kefuId := c.Query("kefu_id")
	user:=models.FindUserById(kefuId)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result":user,
	})
}
func PostKefuInfo(c *gin.Context){
	id:=c.PostForm("id")
	name:=c.PostForm("name")
	password:=c.PostForm("password")
	avator:=c.PostForm("avator")
	nickname:=c.PostForm("nickname")
	roleId:=c.PostForm("role_id")
	if roleId==""{
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "请选择角色!",
		})
		return
	}
	//插入新用户
	if id==""{
		uid:=models.CreateUser(name,tools.Md5(password),avator,nickname)
		roleIdInt,_:=strconv.Atoi(roleId)
		models.CreateUserRole(uid,uint(roleIdInt))
	}else{
		//更新用户
		if password!=""{
			password=tools.Md5(password)
		}
		models.UpdateUser(id,name,password,avator,nickname)
	}

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
	models.DeleteUserById(kefuId)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
		"result":"",
	})
}
