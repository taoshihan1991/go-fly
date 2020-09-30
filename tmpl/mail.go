package tmpl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//邮箱列表界面
func PageMailList(c *gin.Context) {
	return
	c.HTML(http.StatusOK, "list.html", gin.H{})
}
