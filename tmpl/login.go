package tmpl

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/config"
	"net/http"
)

//登陆界面
func PageLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

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

