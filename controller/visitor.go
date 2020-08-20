package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/config"
	"github.com/taoshihan1991/imaptool/models"
	"log"
	"strconv"
)
func PostVisitor(c *gin.Context) {
	name := c.PostForm("name")
	avator := c.PostForm("avator")
	toId := c.PostForm("to_id")
	id := c.PostForm("id")
	refer := c.PostForm("refer")
	city := c.PostForm("city")
	client_ip := c.PostForm("client_ip")
	if name==""||avator==""||toId==""||id==""||refer==""||city==""||client_ip==""{
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error",
		})
		return
	}
	kefuInfo:=models.FindUser(toId)
	if kefuInfo.ID==0{
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
		return
	}
	models.CreateVisitor(name,avator,c.ClientIP(),toId,id,refer,city,client_ip)

	userInfo := make(map[string]string)
	userInfo["uid"] = id
	userInfo["username"] = name
	userInfo["avator"] = avator
	msg := TypeMessage{
		Type: "userOnline",
		Data: userInfo,
	}
	str, _ := json.Marshal(msg)
	kefuConns:=kefuList[toId]
	if kefuConns!=nil{
		for k,kefuConn:=range kefuConns{
			log.Println(k,"xxxxxxxx")
			kefuConn.WriteMessage(websocket.TextMessage,str)
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
func GetVisitor(c *gin.Context) {
	visitorId:=c.Query("visitorId")
	vistor:=models.FindVisitorByVistorId(visitorId)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result":vistor,
	})
}
// @Summary 获取访客列表接口
// @Produce  json
// @Accept multipart/form-data
// @Param page query   string true "分页"
// @Param token header string true "认证token"
// @Success 200 {object} controller.Response
// @Failure 200 {object} controller.Response
// @Router /visitors [get]
func GetVisitors(c *gin.Context) {
	page,_:=strconv.Atoi(c.Query("page"))
	kefuId,_:=c.Get("kefu_name")
	vistors:=models.FindVisitorsByKefuId(uint(page),config.VisitorPageSize,kefuId.(string))
	count:=models.CountVisitorsByKefuId(kefuId.(string))
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result":gin.H{
			"list":vistors,
			"count":count,
			"pagesize":config.PageSize,
		},
	})
}
// @Summary 获取访客聊天信息接口
// @Produce  json
// @Accept multipart/form-data
// @Param visitorId query   string true "访客ID"
// @Param token header string true "认证token"
// @Success 200 {object} controller.Response
// @Failure 200 {object} controller.Response
// @Router /messages [get]
func GetVisitorMessage(c *gin.Context) {
	visitorId:=c.Query("visitorId")
	messages:=models.FindMessageByVisitorId(visitorId)
	result:=make([]map[string]interface{},0)
	for _,message:=range messages{
		item:=make(map[string]interface{})
		var visitor models.Visitor
		var kefu models.User
		if visitor.Name=="" || kefu.Name==""{
			kefu=models.FindUser(message.KefuId)
			visitor=models.FindVisitorByVistorId(message.VisitorId)
		}
		item["time"]=message.CreatedAt.Format("2006-01-02 15:04:05")
		item["content"]=message.Content
		item["mes_type"]=message.MesType
		item["visitor_name"]=visitor.Name
		item["visitor_avator"]=visitor.Avator
		item["kefu_name"]=kefu.Nickname
		item["kefu_avator"]=kefu.Avator
		result=append(result,item)

	}
	models.ReadMessageByVisitorId(visitorId)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result":result,
	})
}
func GetVisitorOnlines(c *gin.Context) {
	users:=make([]map[string]string,0)
	for uid,visitor :=range clientList{
		userInfo := make(map[string]string)
		userInfo["uid"] = uid
		userInfo["name"] = visitor.name
		userInfo["avator"] = visitor.avator
		users=append(users,userInfo)
	}
	tcps:=make([]string,0)
	for ip,_ :=range clientTcpList{
		tcps=append(tcps,ip)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result":gin.H{
			"ws":users,
			"tcp":tcps,
		},
	})
}