package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)
func CasbinACL(c *gin.Context){
	_, _ := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
}
