package tmpl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 咨询界面
func PageChat(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	refer := c.Query("refer")
	if refer == "" {
		refer = c.Request.Referer()
	}
	if refer == "" {
		refer = "​​Direct Link"
	}
	c.HTML(http.StatusOK, "chat_page.html", gin.H{
		"KEFU_ID": kefuId,
		"Refer":   refer,
	})
}
