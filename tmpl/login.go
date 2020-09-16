package tmpl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//登陆界面
func PageLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}



