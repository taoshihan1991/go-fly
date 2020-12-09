package models

type ReplyItem struct {
	Id      string `json:"item_id"`
	Content string `json:"item_content"`
	GroupId string `json:"group_id"`
}
type ReplyGroup struct {
	Id        string      `json:"group_id"`
	GroupName string      `json:"group_name"`
	Items     []ReplyItem `json:"items"`
}

func FindReplyByUserId(userId interface{}) ReplyGroup {
	var replyGroup ReplyGroup
	DB.Raw("select a.*,b.* from reply_group a left join reply_items b on a.id=b.group_id where a.user_id=? ", userId).Scan(&replyGroup)
	return replyGroup
}
