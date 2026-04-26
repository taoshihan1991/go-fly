package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goflylivechat/models"
	"goflylivechat/tools"
	"log"
	"strconv"
	"time"
)

func SubmitOfflineMessage(c *gin.Context) {
	kefuId := c.PostForm("kefu_id")
	visitorId := c.PostForm("visitor_id")
	visitorName := c.PostForm("visitor_name")
	visitorEmail := c.PostForm("visitor_email")
	visitorPhone := c.PostForm("visitor_phone")
	subject := c.PostForm("subject")
	content := c.PostForm("content")
	refer := c.PostForm("refer")
	extra := c.PostForm("extra")

	if kefuId == "" || content == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "客服ID和留言内容不能为空",
		})
		return
	}

	if visitorName == "" {
		visitorName = "匿名访客"
	}

	if !tools.LimitFreqSingle("offline_message:"+c.ClientIP(), 1, 60) {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "提交频率过快，请稍后再试",
		})
		return
	}

	kefuInfo := models.FindUser(kefuId)
	if kefuInfo.ID == 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "客服不存在",
		})
		return
	}

	if visitorId == "" {
		visitorId = tools.Uuid()
	}

	sourceIp := c.ClientIP()

	msgId := models.CreateOfflineMessage(
		kefuId,
		visitorId,
		visitorName,
		visitorEmail,
		visitorPhone,
		subject,
		content,
		sourceIp,
		refer,
		extra,
	)

	if visitorEmail != "" {
		go SendOfflineMessageNotice(kefuInfo.Nickname, visitorName, content)
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "留言提交成功，客服将尽快回复您",
		"result": gin.H{
			"message_id": msgId,
		},
	})
}

func GetOfflineMessages(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	status := c.DefaultQuery("status", "")

	if pageSize > 50 {
		pageSize = 50
	}

	var query interface{}
	var args []interface{}

	if status != "" {
		if !models.IsValidOfflineMessageStatus(status) {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "无效的状态值",
			})
			return
		}
		query = "offline_message.kefu_id = ? AND offline_message.status = ?"
		args = append(args, kefuName.(string), status)
	} else {
		query = "offline_message.kefu_id = ?"
		args = append(args, kefuName.(string))
	}

	count := models.CountOfflineMessages(query, args...)
	list := models.FindOfflineMessagesByPage(uint(page), uint(pageSize), query, args...)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result": gin.H{
			"count":    count,
			"page":     page,
			"pagesize": pageSize,
			"list":     list,
		},
	})
}

func GetOfflineMessageDetail(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "无效的留言ID",
		})
		return
	}

	msg := models.FindOfflineMessageDetailById(uint(id))
	if msg.ID == 0 {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "留言不存在",
		})
		return
	}

	if msg.KefuId != kefuName.(string) {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "无权访问此留言",
		})
		return
	}

	if msg.Status == models.OfflineMessageStatusUnread {
		models.MarkOfflineMessageAsRead(uint(id))
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result": gin.H{
			"id":             msg.ID,
			"kefu_id":        msg.KefuId,
			"visitor_id":     msg.VisitorId,
			"visitor_name":   msg.VisitorName,
			"visitor_email":  msg.VisitorEmail,
			"visitor_phone":  msg.VisitorPhone,
			"subject":        msg.Subject,
			"content":        msg.Content,
			"status":         msg.Status,
			"status_text":    models.GetOfflineMessageStatusText(msg.Status),
			"reply_content":  msg.ReplyContent,
			"reply_time":     msg.ReplyTime,
			"replied_by":     msg.RepliedBy,
			"source_ip":      msg.SourceIp,
			"refer":          msg.Refer,
			"visitor_avator": msg.VisitorAvator,
			"kefu_name":      msg.KefuName,
			"kefu_avator":    msg.KefuAvator,
			"create_time":    msg.CreateTime,
		},
	})
}

func ReplyOfflineMessage(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	idStr := c.PostForm("id")
	replyContent := c.PostForm("reply_content")

	if idStr == "" || replyContent == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "留言ID和回复内容不能为空",
		})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "无效的留言ID",
		})
		return
	}

	msg := models.FindOfflineMessageById(uint(id))
	if msg.ID == 0 {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "留言不存在",
		})
		return
	}

	if msg.KefuId != kefuName.(string) {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "无权回复此留言",
		})
		return
	}

	models.UpdateOfflineMessageReply(uint(id), replyContent, kefuName.(string))

	if msg.VisitorEmail != "" {
		go SendOfflineReplyNotice(msg.VisitorName, msg.VisitorEmail, replyContent)
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "回复成功",
	})
}

func UpdateOfflineMessageStatus(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	idStr := c.Param("id")
	status := c.PostForm("status")

	if idStr == "" || status == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "留言ID和状态不能为空",
		})
		return
	}

	if !models.IsValidOfflineMessageStatus(status) {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "无效的状态值",
		})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "无效的留言ID",
		})
		return
	}

	msg := models.FindOfflineMessageById(uint(id))
	if msg.ID == 0 {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "留言不存在",
		})
		return
	}

	if msg.KefuId != kefuName.(string) {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "无权操作此留言",
		})
		return
	}

	models.UpdateOfflineMessageStatus(uint(id), models.OfflineMessageStatus(status))

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "状态更新成功",
	})
}

func DeleteOfflineMessage(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	idStr := c.Param("id")

	if idStr == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "留言ID不能为空",
		})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "无效的留言ID",
		})
		return
	}

	msg := models.FindOfflineMessageById(uint(id))
	if msg.ID == 0 {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "留言不存在",
		})
		return
	}

	if msg.KefuId != kefuName.(string) {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "无权删除此留言",
		})
		return
	}

	models.DeleteOfflineMessage(uint(id))

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func GetOfflineMessageStats(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	stats := models.GetOfflineMessageStatsByKefuId(kefuName.(string))

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result": gin.H{
			"unread":   stats["unread"],
			"read":     stats["read"],
			"replied":  stats["replied"],
			"total":    stats["total"],
		},
	})
}

func BatchMarkOfflineMessagesAsRead(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	idsJson := c.PostForm("ids")

	now := time.Now()
	var updateCount int64

	if idsJson == "" {
		result := models.DB.Model(&models.OfflineMessage{}).
			Where("kefu_id = ? AND status = ?", kefuName.(string), models.OfflineMessageStatusUnread).
			Updates(map[string]interface{}{
				"status":     models.OfflineMessageStatusRead,
				"updated_at": now,
			})
		updateCount = result.RowsAffected
	} else {
		var ids []uint
		err := json.Unmarshal([]byte(idsJson), &ids)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "无效的ID格式",
			})
			return
		}

		if len(ids) == 0 {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "没有有效的留言ID",
			})
			return
		}

		for _, id := range ids {
			msg := models.FindOfflineMessageById(id)
			if msg.ID == 0 || msg.KefuId != kefuName.(string) {
				c.JSON(200, gin.H{
					"code": 403,
					"msg":  fmt.Sprintf("无权操作留言ID: %d", id),
				})
				return
			}
		}

		err = models.BatchMarkAsRead(ids)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "批量标记失败",
			})
			return
		}
		updateCount = int64(len(ids))
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  fmt.Sprintf("已标记 %d 条消息为已读", updateCount),
		"result": gin.H{
			"count": updateCount,
		},
	})
}


func SendOfflineMessageNotice(kefuName, visitorName, content string) {
	smtp := models.FindConfig("NoticeEmailSmtp")
	email := models.FindConfig("NoticeEmailAddress")
	password := models.FindConfig("NoticeEmailPassword")
	if smtp == "" || email == "" || password == "" {
		return
	}
	subject := fmt.Sprintf("[离线留言] %s 发来新消息", visitorName)
	body := fmt.Sprintf("访客: %s\n留言内容: %s", visitorName, content)
	err := tools.SendSmtp(smtp, email, password, []string{email}, subject, body)
	if err != nil {
		log.Println("SendOfflineMessageNotice error:", err)
	}
}

func SendOfflineReplyNotice(visitorName, visitorEmail, replyContent string) {
	smtp := models.FindConfig("NoticeEmailSmtp")
	email := models.FindConfig("NoticeEmailAddress")
	password := models.FindConfig("NoticeEmailPassword")
	if smtp == "" || email == "" || password == "" {
		return
	}
	subject := "[留言回复] 客服已回复您的留言"
	body := fmt.Sprintf("尊敬的 %s:\n\n客服已回复您的留言:\n%s", visitorName, replyContent)
	err := tools.SendSmtp(smtp, email, password, []string{visitorEmail}, subject, body)
	if err != nil {
		log.Println("SendOfflineReplyNotice error:", err)
	}
}
