package controller

import (
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"github.com/taoshihan1991/imaptool/ws"
	"strconv"
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
	name := c.PostForm("name")
	password := c.PostForm("password")
	rePassword := c.PostForm("rePassword")
	avator := "/static/images/4.jpg"
	nickname := c.PostForm("nickname")
	captchaCode := c.PostForm("captcha")
	roleId := 1
	if name == "" || password == "" || rePassword == "" || nickname == "" || captchaCode == "" {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "参数不能为空",
			"result": "",
		})
		return
	}
	if password != rePassword {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "密码不一致",
			"result": "",
		})
		return
	}
	oldUser := models.FindUser(name)
	if oldUser.Name != "" {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "用户名已经存在",
			"result": "",
		})
		return
	}
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if !captcha.VerifyString(captchaId.(string), captchaCode) {
			c.JSON(200, gin.H{
				"code":   400,
				"msg":    "验证码验证失败",
				"result": "",
			})
			return
		}
	} else {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "验证码失效",
			"result": "",
		})
		return
	}
	//插入新用户
	uid := models.CreateUser(name, tools.Md5(password), avator, nickname)
	if uid == 0 {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "增加用户失败",
			"result": "",
		})
		return
	}
	models.CreateUserRole(uid, uint(roleId))

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "注册完成",
		"result": "",
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
