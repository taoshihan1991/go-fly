package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetLanguage(c *gin.Context) {
	var lang string
	if lang = c.Param("lang"); lang == "" {
		lang = c.Query("lang")
	}
	if lang == "" || lang != "cn" {
		lang = "en"
	}
	c.Set("lang", lang)
}
