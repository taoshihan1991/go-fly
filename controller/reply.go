package controller

import (
	"github.com/gin-gonic/gin"
	"goflylivechat/models"
	"log"
)

type ReplyForm struct {
	GroupName string `form:"group_name" binding:"required"`
}
type ReplyContentForm struct {
	GroupId  string `form:"group_id" binding:"required"`
	Content  string `form:"content" binding:"required"`
	ItemName string `form:"item_name" binding:"required"`
}

func GetReplys(c *gin.Context) {
	kefuId, _ := c.Get("kefu_name")
	log.Println(kefuId)
	res := models.FindReplyByUserId(kefuId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": res,
	})
}
func GetAutoReplys(c *gin.Context) {
	kefu_id := c.Query("kefu_id")
	res := models.FindReplyTitleByUserId(kefu_id)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": res,
	})
}
func PostReply(c *gin.Context) {
	var replyForm ReplyForm
	kefuId, _ := c.Get("kefu_name")
	err := c.Bind(&replyForm)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error:" + err.Error(),
		})
		return
	}
	models.CreateReplyGroup(replyForm.GroupName, kefuId.(string))
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
func PostReplyContent(c *gin.Context) {
	var replyContentForm ReplyContentForm
	kefuId, _ := c.Get("kefu_name")
	err := c.Bind(&replyContentForm)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 200,
			"msg":  "error:" + err.Error(),
		})
		return
	}
	models.CreateReplyContent(replyContentForm.GroupId, kefuId.(string), replyContentForm.Content, replyContentForm.ItemName)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
func PostReplyContentSave(c *gin.Context) {
	kefuId, _ := c.Get("kefu_name")
	replyId := c.PostForm("reply_id")
	replyTitle := c.PostForm("reply_title")
	replyContent := c.PostForm("reply_content")
	if replyId == "" || replyTitle == "" || replyContent == "" {
		c.JSON(400, gin.H{
			"code": 200,
			"msg":  "参数错误!",
		})
		return
	}
	models.UpdateReplyContent(replyId, kefuId.(string), replyTitle, replyContent)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
func DelReplyContent(c *gin.Context) {
	kefuId, _ := c.Get("kefu_name")
	id := c.Query("id")
	models.DeleteReplyContent(id, kefuId.(string))
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
func DelReplyGroup(c *gin.Context) {
	kefuId, _ := c.Get("kefu_name")
	id := c.Query("id")
	models.DeleteReplyGroup(id, kefuId.(string))
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
func PostReplySearch(c *gin.Context) {
	kefuId, _ := c.Get("kefu_name")
	search := c.PostForm("search")
	if search == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	res := models.FindReplyBySearcch(kefuId, search)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": res,
	})
}
