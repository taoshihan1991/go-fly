package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
)

func GetReplys(c *gin.Context) {
	kefuId, _ := c.Get("kefu_id")
	models.FindReplyByUserId(kefuId)
}
