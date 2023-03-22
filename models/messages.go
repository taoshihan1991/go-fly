package models

import (
	"fmt"
	"time"
)

type Message struct {
	Model
	KefuId    string `json:"kefu_id"`
	VisitorId string `json:"visitor_id"`
	Content   string `json:"content"`
	MesType   string `json:"mes_type"`
	Status    string `json:"status"`
}
type MessageKefu struct {
	Model
	KefuId        string `json:"kefu_id"`
	VisitorId     string `json:"visitor_id"`
	Content       string `json:"content"`
	MesType       string `json:"mes_type"`
	Status        string `json:"status"`
	VisitorName   string `json:"visitor_name"`
	VisitorAvator string `json:"visitor_avator"`
	KefuName      string `json:"kefu_name"`
	KefuAvator    string `json:"kefu_avator"`
	CreateTime    string `json:"create_time"`
}

func CreateMessage(kefu_id string, visitor_id string, content string, mes_type string) {
	DB.Exec("set names utf8mb4")
	v := &Message{
		KefuId:    kefu_id,
		VisitorId: visitor_id,
		Content:   content,
		MesType:   mes_type,
		Status:    "unread",
	}
	v.UpdatedAt = time.Now()
	DB.Create(v)
}
func FindMessageByVisitorId(visitor_id string) []Message {
	var messages []Message
	DB.Where("visitor_id=?", visitor_id).Order("id asc").Find(&messages)
	return messages
}

//修改消息状态
func ReadMessageByVisitorId(visitor_id string) {
	message := &Message{
		Status: "read",
	}
	DB.Model(&message).Where("visitor_id=?", visitor_id).Update(message)
}

//获取未读数
func FindUnreadMessageNumByVisitorId(visitor_id string) uint {
	var count uint
	DB.Where("visitor_id=? and status=?", visitor_id, "unread").Count(&count)
	return count
}

//查询最后一条消息
func FindLastMessage(visitorIds []string) []Message {
	var messages []Message
	if len(visitorIds) <= 0 {
		return messages
	}
	var ids []Message
	DB.Select("MAX(id) id").Where(" visitor_id in (? )", visitorIds).Group("visitor_id").Find(&ids)
	if len(ids) <= 0 {
		return messages
	}
	var idStr = make([]string, 0, 0)
	for _, mes := range ids {
		idStr = append(idStr, fmt.Sprintf("%d", mes.ID))
	}
	DB.Select("visitor_id,id,content").Where(" id in (? )", idStr).Find(&messages)
	//subQuery := DB.
	//	Table("message").
	//	Where(" visitor_id in (? )", visitorIds).
	//	Order("id desc").
	//	Limit(1024).
	//	SubQuery()
	//DB.Raw("SELECT ANY_VALUE(visitor_id) visitor_id,ANY_VALUE(id) id,ANY_VALUE(content) content FROM ? message_alia GROUP BY visitor_id", subQuery).Scan(&messages)
	//DB.Select("ANY_VALUE(visitor_id) visitor_id,MAX(ANY_VALUE(id)) id,ANY_VALUE(content) content").Group("visitor_id").Find(&messages)
	return messages
}

//查询最后一条消息
func FindLastMessageByVisitorId(visitorId string) Message {
	var m Message
	DB.Select("content").Where("visitor_id=?", visitorId).Order("id desc").First(&m)
	return m
}
func FindMessageByWhere(query interface{}, args ...interface{}) []MessageKefu {
	var messages []MessageKefu
	DB.Table("message").Where(query, args...).Select("message.*,visitor.avator visitor_avator,visitor.name visitor_name,user.avator kefu_avator,user.nickname kefu_name").Joins("left join user on message.kefu_id=user.name").Joins("left join visitor on visitor.visitor_id=message.visitor_id").Order("message.id asc").Find(&messages)
	return messages
}

//查询条数
func CountMessage(query interface{}, args ...interface{}) uint {
	var count uint
	DB.Model(&Message{}).Where(query, args...).Count(&count)
	return count
}
//分页查询
func FindMessageByPage(page uint, pagesize uint, query interface{}, args ...interface{}) []*MessageKefu {
	offset := (page - 1) * pagesize
	if offset < 0 {
		offset = 0
	}
	var messages []*MessageKefu
	DB.Table("message").Select("message.*,visitor.avator visitor_avator,visitor.name visitor_name,user.avator kefu_avator,user.nickname kefu_name").Offset(offset).Joins("left join user on message.kefu_id=user.name").Joins("left join visitor on visitor.visitor_id=message.visitor_id").Where(query, args...).Limit(pagesize).Order("message.id desc").Find(&messages)
	for _, mes := range messages {
		mes.CreateTime = mes.CreatedAt.Format("2006-01-02 15:04:05")
	}
	return messages
}