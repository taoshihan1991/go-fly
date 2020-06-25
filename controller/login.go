package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
	"time"
)

//验证接口
func LoginCheckPass(c *gin.Context) {
	authType := c.PostForm("type")
	password := c.PostForm("password")
	username := c.PostForm("username")
	switch authType {
	case "local":
		sessionId := CheckPass(username, password)
		userinfo := make(map[string]interface{})
		userinfo["name"] = username
		userinfo["create_time"] = time.Now().Unix()
		token, err := tools.MakeToken(userinfo)
		userinfo["ref_token"]=true
		refToken, _ := tools.MakeToken(userinfo)
		log.Println(err)
		if sessionId != "" {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "验证成功,正在跳转",
				"result": gin.H{
					"token": token,
					"ref_token":refToken,
					"create_time":userinfo["create_time"],
				},
			})
			return
		}
	case "kefulogin":
		info,ok:=CheckKefuPass(username, password)
		userinfo:= make(map[string]interface{})
		if !ok{
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "验证失败",
			})
			return
		}
		userinfo["name"] = info.Name
		userinfo["type"] = "kefu"
		userinfo["create_time"] = time.Now().Unix()
		token, _ := tools.MakeToken(userinfo)
		userinfo["ref_token"]=true
		refToken, _ := tools.MakeToken(userinfo)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "验证成功,正在跳转",
			"result": gin.H{
				"token": token,
				"ref_token":refToken,
				"create_time":userinfo["create_time"],
			},
		})
		return

	}
	c.JSON(200, gin.H{
		"code": 400,
		"msg":  "验证失败",
	})
}
