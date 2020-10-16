package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
)

func Index(c *gin.Context) {
	jump := models.FindConfig("JumpLang")
	if jump != "cn" {
		jump = "en"
	}
	c.Redirect(302, "/index_"+jump)
}
