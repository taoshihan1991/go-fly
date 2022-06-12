package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
)

/**
域名中间件
*/
func DomainLimitMiddleware(c *gin.Context) {
	//离线或者远程
	if !CheckBindOffcial(c) {
		c.Abort()
		return
	}

}


//绑定官网账户
func CheckBindOffcial(c *gin.Context) bool {
	res, err := tools.HTTPGet("https://gofly.v1kf.com/2/isBindOfficial")
	if err != nil {
		log.Println("离线授权码失败,认证连接失败")
		c.Redirect(302, "/bind")
		c.Abort()
	}
	if string(res) != "success" {
		c.Redirect(302, "/bind")
		c.Abort()
	}
	return true
}
