package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"github.com/taoshihan1991/imaptool/ws"
	"strconv"
)

func GetKefuInfo(c *gin.Context) {
	kefuId, _ := c.Get("kefu_id")
	user := models.FindUserById(kefuId)
	info := make(map[string]interface{})
	info["name"] = user.Nickname
	info["id"] = user.Name
	info["avator"] = user.Avator
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": info,
	})
}
func GetKefuInfoAll(c *gin.Context) {
	id, _ := c.Get("kefu_id")
	userinfo := models.FindUserRole("user.avator,user.name,user.id, role.name role_name", id)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "验证成功",
		"result": userinfo,
	})
}
func GetOtherKefuList(c *gin.Context) {
	idStr, _ := c.Get("kefu_id")
	id := idStr.(float64)
	result := make([]interface{}, 0)
	kefus := models.FindUsers()
	for _, kefu := range kefus {
		if uint(id) == kefu.ID {
			continue
		}

		item := make(map[string]interface{})
		item["name"] = kefu.Name
		item["avator"] = kefu.Avator
		item["status"] = "offline"
		kefus, ok := ws.KefuList[kefu.Name]
		if ok && len(kefus) != 0 {
			item["status"] = "online"
		}
		result = append(result, item)
	}
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": result,
	})
}
func GetKefuInfoSetting(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	user := models.FindUserById(kefuId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": user,
	})
}
func PostKefuInfo(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	password := c.PostForm("password")
	avator := c.PostForm("avator")
	nickname := c.PostForm("nickname")
	roleId := c.PostForm("role_id")
	if roleId == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "请选择角色!",
		})
		return
	}
	//插入新用户
	if id == "" {
		uid := models.CreateUser(name, tools.Md5(password), avator, nickname)
		if uid == 0 {
			c.JSON(200, gin.H{
				"code":   400,
				"msg":    "增加用户失败",
				"result": "",
			})
			return
		}
		roleIdInt, _ := strconv.Atoi(roleId)
		models.CreateUserRole(uid, uint(roleIdInt))
	} else {
		//更新用户
		if password != "" {
			password = tools.Md5(password)
		}
		models.UpdateUser(id, name, password, avator, nickname)
		roleIdInt, _ := strconv.Atoi(roleId)
		uid, _ := strconv.Atoi(id)
		models.DeleteRoleByUserId(uid)
		models.CreateUserRole(uint(uid), uint(roleIdInt))
	}

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
func GetKefuList(c *gin.Context) {
	users := models.FindUsers()
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "获取成功",
		"result": users,
	})
}
func DeleteKefuInfo(c *gin.Context) {
	kefuId := c.Query("id")
	models.DeleteUserById(kefuId)
	models.DeleteRoleByUserId(kefuId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "删除成功",
		"result": "",
	})
}
