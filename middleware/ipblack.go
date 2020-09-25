package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
)

func Ipblack(c *gin.Context) {
	ip := c.ClientIP()
	ipblack := models.FindIp(ip)
	if ipblack.IP != "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "IP已被加入黑名单",
		})
		c.Abort()
		return
	}
}
