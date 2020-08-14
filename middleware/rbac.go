package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"strings"
)

func RbacAuth(c *gin.Context){
	roleId, _ :=c.Get("role_id")
	role:=models.FindRole(roleId)
	if role.Method!="*"{
		methods:=strings.Split(role.Method,",")
		for _,m:=range methods{
			if c.Request.Method!=m{
				c.JSON(200, gin.H{
					"code": 403,
					"msg":  "没有权限:"+c.Request.Method+","+c.Request.RequestURI,
				})
				c.Abort()
				return
			}
		}
	}
	if role.Path!="*"{
		paths:=strings.Split(role.Path,",")
		for _,p:=range paths{
			if c.Request.RequestURI!=p{
				c.JSON(200, gin.H{
					"code": 403,
					"msg":  "没有权限:"+c.Request.Method+","+c.Request.RequestURI,
				})
				c.Abort()
			}
		}
	}
}

