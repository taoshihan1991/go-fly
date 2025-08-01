package controller

import (
	"github.com/gin-gonic/gin"
	"goflylivechat/models"
)

func GetConfigs(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	configs := models.FindConfigsByUserId(kefuName)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": configs,
	})
}
func GetConfig(c *gin.Context) {
	key := c.Query("key")
	config := models.FindConfig(key)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": config,
	})
}
func PostConfig(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")
	kefuName, _ := c.Get("kefu_name")
	if key == "" || value == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error",
		})
		return
	}
	models.UpdateConfig(kefuName, key, value)

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
