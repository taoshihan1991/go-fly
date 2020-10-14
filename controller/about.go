package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
)

func GetAbout(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		page = "index"
	}
	about := models.FindAboutByPage(page)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": about,
	})
}
