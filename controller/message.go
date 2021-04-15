package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/common"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"github.com/taoshihan1991/imaptool/ws"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func SendMessageV2(c *gin.Context) {
	fromId := c.PostForm("from_id")
	toId := c.PostForm("to_id")
	content := c.PostForm("content")
	cType := c.PostForm("type")
	if content == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "内容不能为空",
		})
		return
	}
	//限流
	if !tools.LimitFreqSingle("sendmessage:"+c.ClientIP(), 1, 2) {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  c.ClientIP() + "发送频率过快",
		})
		return
	}
	var kefuInfo models.User
	var vistorInfo models.Visitor
	if cType == "kefu" {
		kefuInfo = models.FindUser(fromId)
		vistorInfo = models.FindVisitorByVistorId(toId)
	} else if cType == "visitor" {
		vistorInfo = models.FindVisitorByVistorId(fromId)
		kefuInfo = models.FindUser(toId)
	}

	if kefuInfo.ID == 0 || vistorInfo.ID == 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
		return
	}

	models.CreateMessage(kefuInfo.Name, vistorInfo.VisitorId, content, cType)
	//var msg TypeMessage
	if cType == "kefu" {
		guest, ok := ws.ClientList[vistorInfo.VisitorId]

		if guest != nil && ok {
			ws.VisitorMessage(vistorInfo.VisitorId, content, kefuInfo)
		}
		ws.KefuMessage(vistorInfo.VisitorId, content, kefuInfo)
		//msg = TypeMessage{
		//	Type: "message",
		//	Data: ws.ClientMessage{
		//		Name:    kefuInfo.Nickname,
		//		Avator:  kefuInfo.Avator,
		//		Id:      vistorInfo.VisitorId,
		//		Time:    time.Now().Format("2006-01-02 15:04:05"),
		//		ToId:    vistorInfo.VisitorId,
		//		Content: content,
		//		IsKefu:  "yes",
		//	},
		//}
		//str2, _ := json.Marshal(msg)
		//ws.OneKefuMessage(kefuInfo.Name, str2)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
		})
	}
	if cType == "visitor" {
		guest, ok := ws.ClientList[vistorInfo.VisitorId]
		if ok && guest != nil {
			guest.UpdateTime = time.Now()
		}
		//kefuConns, ok := ws.KefuList[kefuInfo.Name]
		//if kefuConns == nil || !ok {
		//	c.JSON(200, gin.H{
		//		"code": 200,
		//		"msg":  "ok",
		//	})
		//	return
		//}
		msg := ws.TypeMessage{
			Type: "message",
			Data: ws.ClientMessage{
				Avator:  vistorInfo.Avator,
				Id:      vistorInfo.VisitorId,
				Name:    vistorInfo.Name,
				ToId:    kefuInfo.Name,
				Content: content,
				Time:    time.Now().Format("2006-01-02 15:04:05"),
				IsKefu:  "no",
			},
		}
		str, _ := json.Marshal(msg)
		ws.OneKefuMessage(kefuInfo.Name, str)
		//ws.KefuMessage(vistorInfo.VisitorId, content, kefuInfo)
		go ws.SendServerJiang(vistorInfo.Name+"说", content, c.Request.Host)
		go SendAppGetuiPush(kefuInfo.Name, vistorInfo.Name, content)
		kefus, ok := ws.KefuList[kefuInfo.Name]
		if !ok || len(kefus) == 0 {
			go SendNoticeEmail(content+"|"+vistorInfo.Name, content)
		}
		go ws.VisitorAutoReply(vistorInfo, kefuInfo, content)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
		})
	}

}
func SendVisitorNotice(c *gin.Context) {
	notice := c.Query("msg")
	if notice == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "msg不能为空",
		})
		return
	}
	msg := ws.TypeMessage{
		Type: "notice",
		Data: notice,
	}
	str, _ := json.Marshal(msg)
	for _, visitor := range ws.ClientList {
		visitor.Conn.WriteMessage(websocket.TextMessage, str)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
func SendCloseMessageV2(c *gin.Context) {
	visitorId := c.Query("visitor_id")
	if visitorId == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "visitor_id不能为空",
		})
		return
	}

	oldUser, ok := ws.ClientList[visitorId]
	if oldUser != nil || ok {
		msg := ws.TypeMessage{
			Type: "force_close",
			Data: visitorId,
		}
		str, _ := json.Marshal(msg)
		err := oldUser.Conn.WriteMessage(websocket.TextMessage, str)
		oldUser.Conn.Close()
		delete(ws.ClientList, visitorId)
		tools.Logger().Println("close_message", oldUser, err)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
func UploadImg(c *gin.Context) {
	f, err := c.FormFile("imgfile")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "上传失败!",
		})
		return
	} else {

		fileExt := strings.ToLower(path.Ext(f.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
			})
			return
		}
		isMainUploadExist, _ := tools.IsFileExist(common.Upload)
		if !isMainUploadExist {
			os.Mkdir(common.Upload, os.ModePerm)
		}
		fileName := tools.Md5(fmt.Sprintf("%s%s", f.Filename, time.Now().String()))
		fildDir := fmt.Sprintf("%s%d%s/", common.Upload, time.Now().Year(), time.Now().Month().String())
		isExist, _ := tools.IsFileExist(fildDir)
		if !isExist {
			os.Mkdir(fildDir, os.ModePerm)
		}
		filepath := fmt.Sprintf("%s%s%s", fildDir, fileName, fileExt)
		c.SaveUploadedFile(f, filepath)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "上传成功!",
			"result": gin.H{
				"path": filepath,
			},
		})
	}
}
func UploadFile(c *gin.Context) {
	SendAttachment, err := strconv.ParseBool(models.FindConfig("SendAttachment"))
	if !SendAttachment || err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "禁止上传附件!",
		})
		return
	}
	f, err := c.FormFile("realfile")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "上传失败!",
		})
		return
	} else {

		fileExt := strings.ToLower(path.Ext(f.Filename))
		if f.Size >= 90*1024*1024 {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "上传失败!不允许超过90M",
			})
			return
		}

		fileName := tools.Md5(fmt.Sprintf("%s%s", f.Filename, time.Now().String()))
		fildDir := fmt.Sprintf("%s%d%s/", common.Upload, time.Now().Year(), time.Now().Month().String())
		isExist, _ := tools.IsFileExist(fildDir)
		if !isExist {
			os.Mkdir(fildDir, os.ModePerm)
		}
		filepath := fmt.Sprintf("%s%s%s", fildDir, fileName, fileExt)
		c.SaveUploadedFile(f, filepath)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "上传成功!",
			"result": gin.H{
				"path": filepath,
			},
		})
	}
}
func GetMessagesV2(c *gin.Context) {
	visitorId := c.Query("visitor_id")
	messages := models.FindMessageByVisitorId(visitorId)
	//result := make([]map[string]interface{}, 0)
	chatMessages := make([]ChatMessage, 0)
	var visitor models.Visitor
	var kefu models.User
	for _, message := range messages {
		//item := make(map[string]interface{})
		if visitor.Name == "" || kefu.Name == "" {
			kefu = models.FindUser(message.KefuId)
			visitor = models.FindVisitorByVistorId(message.VisitorId)
		}
		var chatMessage ChatMessage
		chatMessage.Time = message.CreatedAt.Format("2006-01-02 15:04:05")
		chatMessage.Content = message.Content
		chatMessage.MesType = message.MesType
		if message.MesType == "kefu" {
			chatMessage.Name = kefu.Nickname
			chatMessage.Avator = kefu.Avator
		} else {
			chatMessage.Name = visitor.Name
			chatMessage.Avator = visitor.Avator
		}
		chatMessages = append(chatMessages, chatMessage)
	}
	models.ReadMessageByVisitorId(visitorId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": chatMessages,
	})
}
