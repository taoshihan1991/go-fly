package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetLanguage(c *gin.Context){
	lang := c.Query("lang")
	if lang == "" ||lang!="cn"{
		lang = "en"
	}
	c.Set("lang",lang)
}

