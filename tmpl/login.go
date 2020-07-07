package tmpl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//登陆界面
func PageLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

//咨询界面
func PageChat(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	c.HTML(http.StatusOK, "chat_page.html", gin.H{
		"KEFU_ID":kefuId,
	})
}

