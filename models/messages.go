package models
type Message struct {
	Model
	KefuId string `json:"kefu_id"`
	VisitorId string `json:"visitor_id"`
	Content string `json:"content"`
	MesType string `json:"mes_type"`
	Status string `json:"status"`
}
func CreateMessage(kefu_id string,visitor_id string,content string,mes_type string){
	v:=&Message{
		KefuId:   kefu_id,
		VisitorId: visitor_id,
		Content:   content,
		MesType:  mes_type,
		Status:    "unread",
	}
	DB.Create(v)
}
func FindMessageByVisitorId(visitor_id string)[]Message{
	var messages []Message
	DB.Where("visitor_id=?",visitor_id).Order("id asc").Find(&messages)
	return messages
}