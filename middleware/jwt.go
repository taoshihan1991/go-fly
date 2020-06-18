package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
)
func JwtPageMiddleware(c *gin.Context){
	token := c.Query("token")
	userinfo := tools.ParseToken(token)
	log.Println(userinfo)
	if userinfo == nil {
		c.Redirect(302,"/login")
		c.Abort()
	}
}
func JwtApiMiddleware(c *gin.Context){
	log.Println("路由中间件")
	token := c.Query("token")
	userinfo := tools.ParseToken(token)
	log.Println(userinfo)
	if userinfo == nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "验证失败",
		})
		c.Abort()
	}
}
