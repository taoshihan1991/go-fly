package tmpl

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/config"
	"net/http"
)

//咨询界面
func PageChat(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	lang,_ := c.Get("lang")
	language:=config.CreateLanguage(lang.(string))
	refer := c.Query("refer")
	if refer==""{
		refer=c.Request.Referer()
	}
	c.HTML(http.StatusOK, "chat_page.html", gin.H{
		"KEFU_ID":kefuId,
		"SendBtn":language.Send,
		"Lang":lang.(string),
		"Refer":refer,
	})
}
func PageKfChat(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	visitorId:=c.Query("visitor_id")
	token:=c.Query("token")
	c.HTML(http.StatusOK, "chat_kf_page.html", gin.H{
		"KefuId":kefuId,
		"VisitorId":visitorId,
		"Token":token,
	})
}
