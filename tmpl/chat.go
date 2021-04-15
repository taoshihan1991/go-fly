package tmpl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//咨询界面
func PageChat(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	lang, _ := c.Get("lang")
	refer := c.Query("refer")
	if refer == "" {
		refer = c.Request.Referer()
	}
	if refer == "" {
		refer = "直接访问"
	}
	c.HTML(http.StatusOK, "chat_page.html", gin.H{
		"KEFU_ID": kefuId,
		"Lang":    lang.(string),
		"Refer":   refer,
	})
}
func PageKfChat(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	visitorId := c.Query("visitor_id")
	token := c.Query("token")
	c.HTML(http.StatusOK, "chat_kf_page.html", gin.H{
		"KefuId":    kefuId,
		"VisitorId": visitorId,
		"Token":     token,
	})
}
