package tmpl

import (
	"github.com/gin-gonic/gin"
	"goflylivechat/tools"
	"net/http"
)

// 登陆界面
func PageLogin(c *gin.Context) {
	if noExist, _ := tools.IsFileNotExist("./install.lock"); noExist {
		c.Redirect(302, "/install")
	}
	c.HTML(http.StatusOK, "login.html", nil)
}

// 绑定界面
func PageBind(c *gin.Context) {
	c.HTML(http.StatusOK, "bind.html", gin.H{})
}
