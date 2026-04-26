package models

import (
	"time"
)

type OfflineMessageStatus string

const (
	OfflineMessageStatusUnread    OfflineMessageStatus = "unread"
	OfflineMessageStatusRead      OfflineMessageStatus = "read"
	OfflineMessageStatusReplied   OfflineMessageStatus = "replied"
)

type OfflineMessage struct {
	Model
	KefuId       string                `json:"kefu_id"`
	VisitorId    string                `json:"visitor_id"`
	VisitorName  string                `json:"visitor_name"`
	VisitorEmail string                `json:"visitor_email"`
	VisitorPhone string                `json:"visitor_phone"`
	Subject      string                `json:"subject"`
	Content      string                `json:"content"`
	Status       OfflineMessageStatus  `json:"status"`
	ReplyContent string                `json:"reply_content"`
	ReplyTime    *time.Time            `json:"reply_time"`
	RepliedBy    string                `json:"replied_by"`
	SourceIp     string                `json:"source_ip"`
	Refer        string                `json:"refer"`
	Extra        string                `json:"extra"`
}

type OfflineMessageDetail struct {
	OfflineMessage
	VisitorAvator string `json:"visitor_avator"`
	KefuName      string `json:"kefu_name"`
	KefuAvator    string `json:"kefu_avator"`
	CreateTime    string `json:"create_time"`
}

func CreateOfflineMessage(kefuId, visitorId, visitorName, visitorEmail, visitorPhone, subject, content, sourceIp, refer, extra string) uint {
	msg := &OfflineMessage{
		KefuId:       kefuId,
		VisitorId:    visitorId,
		VisitorName:  visitorName,
		VisitorEmail: visitorEmail,
		VisitorPhone: visitorPhone,
		Subject:      subject,
		Content:      content,
		Status:       OfflineMessageStatusUnread,
		SourceIp:     sourceIp,
		Refer:        refer,
		Extra:        extra,
	}
	msg.UpdatedAt = time.Now()
	DB.Create(msg)
	return msg.ID
}

func FindOfflineMessageById(id uint) OfflineMessage {
	var msg OfflineMessage
	DB.Where("id = ?", id).First(&msg)
	return msg
}

func FindOfflineMessageDetailById(id uint) OfflineMessageDetail {
	var msg OfflineMessageDetail
	DB.Table("offline_message").
		Select("offline_message.*, visitor.avator as visitor_avator, user.nickname as kefu_name, user.avator as kefu_avator").
		Joins("left join user on offline_message.kefu_id = user.name").
		Joins("left join visitor on offline_message.visitor_id = visitor.visitor_id").
		Where("offline_message.id = ?", id).
		First(&msg)
	msg.CreateTime = msg.CreatedAt.Format("2006-01-02 15:04:05")
	return msg
}

func UpdateOfflineMessageStatus(id uint, status OfflineMessageStatus) {
	msg := &OfflineMessage{
		Status: status,
	}
	msg.UpdatedAt = time.Now()
	DB.Model(&OfflineMessage{}).Where("id = ?", id).Update(msg)
}

func UpdateOfflineMessageReply(id uint, replyContent string, repliedBy string) {
	now := time.Now()
	msg := &OfflineMessage{
		Status:       OfflineMessageStatusReplied,
		ReplyContent: replyContent,
		ReplyTime:    &now,
		RepliedBy:    repliedBy,
	}
	msg.UpdatedAt = now
	DB.Model(&OfflineMessage{}).Where("id = ?", id).Updates(msg)
}

func DeleteOfflineMessage(id uint) {
	DB.Where("id = ?", id).Delete(OfflineMessage{})
}

func CountOfflineMessagesByKefuId(kefuId string, status OfflineMessageStatus) uint {
	var count uint
	query := DB.Model(&OfflineMessage{}).Where("kefu_id = ?", kefuId)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	query.Count(&count)
	return count
}

func CountOfflineMessages(query interface{}, args ...interface{}) uint {
	var count uint
	DB.Model(&OfflineMessage{}).Where(query, args...).Count(&count)
	return count
}

func FindOfflineMessagesByPage(page uint, pagesize uint, query interface{}, args ...interface{}) []*OfflineMessageDetail {
	offset := (page - 1) * pagesize
	if offset < 0 {
		offset = 0
	}
	var messages []*OfflineMessageDetail
	baseQuery := DB.Table("offline_message").
		Select("offline_message.*, visitor.avator as visitor_avator, user.nickname as kefu_name, user.avator as kefu_avator").
		Joins("left join user on offline_message.kefu_id = user.name").
		Joins("left join visitor on offline_message.visitor_id = visitor.visitor_id")
	
	if query != nil {
		baseQuery = baseQuery.Where(query, args...)
	}
	
	baseQuery.Offset(offset).Limit(pagesize).Order("offline_message.created_at desc").Find(&messages)
	
	for _, msg := range messages {
		msg.CreateTime = msg.CreatedAt.Format("2006-01-02 15:04:05")
	}
	return messages
}

func FindOfflineMessagesByStatus(kefuId string, status OfflineMessageStatus) []OfflineMessage {
	var messages []OfflineMessage
	query := DB.Where("kefu_id = ?", kefuId)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	query.Order("created_at desc").Find(&messages)
	return messages
}

func FindUnreadOfflineMessagesByKefuId(kefuId string) []OfflineMessage {
	return FindOfflineMessagesByStatus(kefuId, OfflineMessageStatusUnread)
}

func MarkOfflineMessageAsRead(id uint) {
	UpdateOfflineMessageStatus(id, OfflineMessageStatusRead)
}

func IsValidOfflineMessageStatus(status string) bool {
	s := OfflineMessageStatus(status)
	return s == OfflineMessageStatusUnread || 
		   s == OfflineMessageStatusRead || 
		   s == OfflineMessageStatusReplied
}

func (s OfflineMessageStatus) String() string {
	return string(s)
}

func GetOfflineMessageStatusText(status OfflineMessageStatus) string {
	switch status {
	case OfflineMessageStatusUnread:
		return "未读"
	case OfflineMessageStatusRead:
		return "已读"
	case OfflineMessageStatusReplied:
		return "已回复"
	default:
		return "未知"
	}
}

func BatchUpdateOfflineMessageStatus(ids []uint, status OfflineMessageStatus) error {
	if len(ids) == 0 {
		return nil
	}
	now := time.Now()
	return DB.Model(&OfflineMessage{}).
		Where("id IN (?)", ids).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": now,
		}).Error
}

func BatchMarkAsRead(ids []uint) error {
	return BatchUpdateOfflineMessageStatus(ids, OfflineMessageStatusRead)
}

func GetOfflineMessageStatsByKefuId(kefuId string) map[string]int64 {
	stats := make(map[string]int64)
	
	var unreadCount int64
	DB.Model(&OfflineMessage{}).Where("kefu_id = ? AND status = ?", kefuId, OfflineMessageStatusUnread).Count(&unreadCount)
	stats["unread"] = unreadCount
	
	var readCount int64
	DB.Model(&OfflineMessage{}).Where("kefu_id = ? AND status = ?", kefuId, OfflineMessageStatusRead).Count(&readCount)
	stats["read"] = readCount
	
	var repliedCount int64
	DB.Model(&OfflineMessage{}).Where("kefu_id = ? AND status = ?", kefuId, OfflineMessageStatusReplied).Count(&repliedCount)
	stats["replied"] = repliedCount
	
	var totalCount int64
	DB.Model(&OfflineMessage{}).Where("kefu_id = ?", kefuId).Count(&totalCount)
	stats["total"] = totalCount
	
	return stats
}

func FindOfflineMessagesByVisitorId(visitorId string) []OfflineMessage {
	var messages []OfflineMessage
	DB.Where("visitor_id = ?", visitorId).Order("created_at desc").Find(&messages)
	return messages
}

func FindLastOfflineMessageByVisitorId(visitorId string) OfflineMessage {
	var msg OfflineMessage
	DB.Where("visitor_id = ?", visitorId).Order("created_at desc").First(&msg)
	return msg
}
