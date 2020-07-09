package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"log"
)
func CasbinACL(c *gin.Context){
	roleId, _ :=c.Get("role_id")
	sub:=fmt.Sprintf("%s_%d","role",int(roleId.(float64)))
	obj:=c.Request.RequestURI
	act:=c.Request.Method
	e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	log.Println(sub,obj,act,err)
	ok,err:=e.Enforce(sub,obj,act)
	if err!=nil{
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "没有权限:"+err.Error(),
		})
		c.Abort()
	}
	if !ok{
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  fmt.Sprintf("没有权限:%s,%s,%s",sub,obj,act),
		})
		c.Abort()
	}
}
