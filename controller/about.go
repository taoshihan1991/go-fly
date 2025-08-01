package controller

import (
	"github.com/gin-gonic/gin"
	"goflylivechat/models"
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
func GetAbouts(c *gin.Context) {
	about := models.FindAbouts()
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": about,
	})
}
func PostAbout(c *gin.Context) {
	title_cn := c.PostForm("title_cn")
	title_en := c.PostForm("title_en")
	keywords_cn := c.PostForm("keywords_cn")
	keywords_en := c.PostForm("keywords_en")
	desc_cn := c.PostForm("desc_cn")
	desc_en := c.PostForm("desc_en")
	css_js := c.PostForm("css_js")
	html_cn := c.PostForm("html_cn")
	html_en := c.PostForm("html_en")
	if title_cn == "" || title_en == "" || html_cn == "" || html_en == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error",
		})
		return
	}
	models.UpdateAbout("index", title_cn, title_en, keywords_cn, keywords_en, desc_cn, desc_en, css_js, html_cn, html_en)

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
