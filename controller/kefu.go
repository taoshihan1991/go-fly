package controller

import (
	"github.com/gin-gonic/gin"
	"goflylivechat/models"
	"goflylivechat/tools"
	"goflylivechat/ws"
	"net/http"
)

func PostKefuAvator(c *gin.Context) {

	avator := c.PostForm("avator")
	if avator == "" {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "不能为空",
			"result": "",
		})
		return
	}
	kefuName, _ := c.Get("kefu_name")
	models.UpdateUserAvator(kefuName.(string), avator)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
func PostKefuPass(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	newPass := c.PostForm("new_pass")
	confirmNewPass := c.PostForm("confirm_new_pass")
	old_pass := c.PostForm("old_pass")
	if newPass != confirmNewPass {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "密码不一致",
			"result": "",
		})
		return
	}
	user := models.FindUser(kefuName.(string))
	if user.Password != tools.Md5(old_pass) {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "旧密码不正确",
			"result": "",
		})
		return
	}
	models.UpdateUserPass(kefuName.(string), tools.Md5(newPass))
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
func PostKefuClient(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	clientId := c.PostForm("client_id")

	if clientId == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "client_id不能为空",
		})
		return
	}
	models.CreateUserClient(kefuName.(string), clientId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
func GetKefuInfo(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	user := models.FindUser(kefuName.(string))
	info := make(map[string]interface{})
	info["avator"] = user.Avator
	info["username"] = user.Name
	info["nickname"] = user.Nickname
	info["uid"] = user.ID
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
	ws.SendPingToKefuClient()
	kefus := models.FindUsers()
	for _, kefu := range kefus {
		if uint(id) == kefu.ID {
			continue
		}

		item := make(map[string]interface{})
		item["name"] = kefu.Name
		item["nickname"] = kefu.Nickname
		item["avator"] = kefu.Avator
		item["status"] = "offline"
		kefu, ok := ws.KefuList[kefu.Name]
		if ok && kefu != nil {
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
func PostTransKefu(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	visitorId := c.Query("visitor_id")
	curKefuId, _ := c.Get("kefu_name")
	user := models.FindUser(kefuId)
	visitor := models.FindVisitorByVistorId(visitorId)
	if user.Name == "" || visitor.Name == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "访客或客服不存在",
		})
		return
	}
	models.UpdateVisitorKefu(visitorId, kefuId)
	ws.UpdateVisitorUser(visitorId, kefuId)
	go ws.VisitorOnline(kefuId, visitor)
	go ws.VisitorOffline(curKefuId.(string), visitor.VisitorId, visitor.Name)
	go ws.VisitorNotice(visitor.VisitorId, "客服转接到"+user.Nickname)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "转移成功",
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
func PostKefuRegister(c *gin.Context) {
	name := c.PostForm("username")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	avatar := "/static/images/4.jpg"

	if name == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":   400,
			"msg":    "All fields are required",
			"result": nil,
		})
		return
	}

	existingUser := models.FindUser(name)
	if existingUser.Name != "" {
		c.JSON(http.StatusOK, gin.H{
			"code":   409,
			"msg":    "Username already exists",
			"result": nil,
		})
		return
	}

	userID := models.CreateUser(name, tools.Md5(password), avatar, nickname)
	if userID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   500,
			"msg":    "Registration Failed",
			"result": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Registration successful",
		"result": gin.H{
			"user_id": userID,
		},
	})
}
func PostKefuInfo(c *gin.Context) {
	name, _ := c.Get("kefu_name")
	password := c.PostForm("password")
	avator := c.PostForm("avator")
	nickname := c.PostForm("nickname")
	if password != "" {
		password = tools.Md5(password)
	}
	if name == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "客服账号不能为空",
		})
		return
	}
	models.UpdateUser(name.(string), password, avator, nickname)

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
