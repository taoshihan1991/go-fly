package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/common"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"github.com/taoshihan1991/imaptool/ws"
	"math/rand"
	"strconv"
)

//func PostVisitor(c *gin.Context) {
//	name := c.PostForm("name")
//	avator := c.PostForm("avator")
//	toId := c.PostForm("to_id")
//	id := c.PostForm("id")
//	refer := c.PostForm("refer")
//	city := c.PostForm("city")
//	client_ip := c.PostForm("client_ip")
//	if name == "" || avator == "" || toId == "" || id == "" || refer == "" || city == "" || client_ip == "" {
//		c.JSON(200, gin.H{
//			"code": 400,
//			"msg":  "error",
//		})
//		return
//	}
//	kefuInfo := models.FindUser(toId)
//	if kefuInfo.ID == 0 {
//		c.JSON(200, gin.H{
//			"code": 400,
//			"msg":  "用户不存在",
//		})
//		return
//	}
//	models.CreateVisitor(name, avator, c.ClientIP(), toId, id, refer, city, client_ip)
//
//	userInfo := make(map[string]string)
//	userInfo["uid"] = id
//	userInfo["username"] = name
//	userInfo["avator"] = avator
//	msg := TypeMessage{
//		Type: "userOnline",
//		Data: userInfo,
//	}
//	str, _ := json.Marshal(msg)
//	kefuConns := kefuList[toId]
//	if kefuConns != nil {
//		for k, kefuConn := range kefuConns {
//			log.Println(k, "xxxxxxxx")
//			kefuConn.WriteMessage(websocket.TextMessage, str)
//		}
//	}
//	c.JSON(200, gin.H{
//		"code": 200,
//		"msg":  "ok",
//	})
//}
func PostVisitorLogin(c *gin.Context) {
	ipcity := tools.ParseIp(c.ClientIP())
	avator := fmt.Sprintf("/static/images/%d.jpg", rand.Intn(14))
	toId := c.PostForm("to_id")
	id := c.PostForm("visitor_id")

	if id == "" {
		id = tools.Uuid()
	}
	refer := c.PostForm("refer")
	var (
		city string
		name string
	)
	if ipcity != nil {
		city = ipcity.CountryName + ipcity.RegionName + ipcity.CityName
		name = ipcity.CountryName + ipcity.RegionName + ipcity.CityName + "网友"
	} else {
		city = "未识别地区"
		name = "匿名网友"
	}
	client_ip := c.ClientIP()
	extra := c.PostForm("extra")
	extraJson := tools.Base64Decode(extra)
	if extraJson != "" {
		var extraObj VisitorExtra
		err := json.Unmarshal([]byte(extraJson), &extraObj)
		if err == nil {
			if extraObj.VisitorName != "" {
				name = extraObj.VisitorName
			}
			if extraObj.VisitorAvatar != "" {
				avator = extraObj.VisitorAvatar
			}
		}
	}
	//log.Println(name,avator,c.ClientIP(),toId,id,refer,city,client_ip)
	if name == "" || avator == "" || toId == "" || id == "" || refer == "" || city == "" || client_ip == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error",
		})
		return
	}
	kefuInfo := models.FindUser(toId)
	if kefuInfo.ID == 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "客服不存在",
		})
		return
	}
	visitor := models.FindVisitorByVistorId(id)
	if visitor.Name != "" {
		avator = visitor.Avator
		//更新状态上线
		models.UpdateVisitor(name, visitor.Avator, id, 1, c.ClientIP(), c.ClientIP(), refer, extra)
	} else {
		models.CreateVisitor(name, avator, c.ClientIP(), toId, id, refer, city, client_ip, extra)
	}
	visitor.Name = name
	visitor.Avator = avator
	visitor.ToId = toId
	visitor.ClientIp = c.ClientIP()
	visitor.VisitorId = id

	//各种通知
	go SendNoticeEmail(visitor.Name, "来了")
	go SendAppGetuiPush(kefuInfo.Name, visitor.Name, visitor.Name+"来了")
	go SendVisitorLoginNotice(kefuInfo.Name, visitor.Name, visitor.Avator, visitor.Name+"来了", visitor.VisitorId)
	go ws.VisitorOnline(kefuInfo.Name, visitor)
	//go SendServerJiang(visitor.Name, "来了", c.Request.Host)

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": visitor,
	})
}
func GetVisitor(c *gin.Context) {
	visitorId := c.Query("visitorId")
	vistor := models.FindVisitorByVistorId(visitorId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": vistor,
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
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	if pagesize == 0 {
		pagesize = int(common.VisitorPageSize)
	}
	kefuId, _ := c.Get("kefu_name")
	vistors := models.FindVisitorsByKefuId(uint(page), uint(pagesize), kefuId.(string))
	count := models.CountVisitorsByKefuId(kefuId.(string))
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result": gin.H{
			"list":     vistors,
			"count":    count,
			"pagesize": common.PageSize,
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
	visitorId := c.Query("visitorId")

	query := "message.visitor_id= ?"
	messages := models.FindMessageByWhere(query, visitorId)
	result := make([]map[string]interface{}, 0)
	for _, message := range messages {
		item := make(map[string]interface{})

		item["time"] = message.CreatedAt.Format("2006-01-02 15:04:05")
		item["content"] = message.Content
		item["mes_type"] = message.MesType
		item["visitor_name"] = message.VisitorName
		item["visitor_avator"] = message.VisitorAvator
		item["kefu_name"] = message.KefuName
		item["kefu_avator"] = message.KefuAvator
		result = append(result, item)

	}
	go models.ReadMessageByVisitorId(visitorId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": result,
	})
}

// @Summary 获取在线访客列表接口
// @Produce  json
// @Success 200 {object} controller.Response
// @Failure 200 {object} controller.Response
// @Router /visitors_online [get]
func GetVisitorOnlines(c *gin.Context) {
	users := make([]map[string]string, 0)
	visitorIds := make([]string, 0)
	for uid, visitor := range ws.ClientList {
		userInfo := make(map[string]string)
		userInfo["uid"] = uid
		userInfo["name"] = visitor.Name
		userInfo["avator"] = visitor.Avator
		users = append(users, userInfo)
		visitorIds = append(visitorIds, visitor.Id)
	}

	//查询最新消息
	messages := models.FindLastMessage(visitorIds)
	temp := make(map[string]string, 0)
	for _, mes := range messages {
		temp[mes.VisitorId] = mes.Content
	}
	for _, user := range users {
		user["last_message"] = temp[user["uid"]]
	}

	tcps := make([]string, 0)
	for ip, _ := range clientTcpList {
		tcps = append(tcps, ip)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result": gin.H{
			"ws":  users,
			"tcp": tcps,
		},
	})
}

// @Summary 获取客服的在线访客列表接口
// @Produce  json
// @Success 200 {object} controller.Response
// @Failure 200 {object} controller.Response
// @Router /visitors_kefu_online [get]
func GetKefusVisitorOnlines(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	users := make([]*VisitorOnline, 0)
	visitorIds := make([]string, 0)
	for uid, visitor := range ws.ClientList {
		if visitor.To_id != kefuName {
			continue
		}
		userInfo := new(VisitorOnline)
		userInfo.Uid = uid
		userInfo.Username = visitor.Name
		userInfo.Avator = visitor.Avator
		users = append(users, userInfo)
		visitorIds = append(visitorIds, visitor.Id)
	}

	//查询最新消息
	messages := models.FindLastMessage(visitorIds)
	temp := make(map[string]string, 0)
	for _, mes := range messages {
		temp[mes.VisitorId] = mes.Content
	}
	for _, user := range users {
		user.LastMessage = temp[user.Uid]
		if user.LastMessage == "" {
			user.LastMessage = "新访客"
		}
	}

	tcps := make([]string, 0)
	for ip, _ := range clientTcpList {
		tcps = append(tcps, ip)
	}
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": users,
	})
}
