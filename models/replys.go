package models

type ReplyItem struct {
	Id       string `json:"item_id"`
	Content  string `json:"item_content"`
	GroupId  string `json:"group_id"`
	ItemName string `json:"item_name"`
	UserId   string `json:"user_id"`
}
type ReplyGroup struct {
	Id        string       `json:"group_id"`
	GroupName string       `json:"group_name"`
	UserId    string       `json:"user_id"`
	Items     []*ReplyItem `json:"items";"`
}

func FindReplyItemByUserIdTitle(userId interface{}, title string) ReplyItem {
	var reply ReplyItem
	DB.Where("user_id = ? and item_name = ?", userId, title).Find(&reply)
	return reply
}
func FindReplyByUserId(userId interface{}) []*ReplyGroup {
	var replyGroups []*ReplyGroup
	//DB.Raw("select a.*,b.* from reply_group a left join reply_item b on a.id=b.group_id where a.user_id=? ", userId).Scan(&replyGroups)
	var replyItems []*ReplyItem
	DB.Where("user_id = ?", userId).Find(&replyGroups)
	DB.Where("user_id = ?", userId).Find(&replyItems)
	temp := make(map[string]*ReplyGroup)
	for _, replyGroup := range replyGroups {
		replyGroup.Items = make([]*ReplyItem, 0)
		temp[replyGroup.Id] = replyGroup
	}
	for _, replyItem := range replyItems {
		temp[replyItem.GroupId].Items = append(temp[replyItem.GroupId].Items, replyItem)
	}
	return replyGroups
}
func FindReplyTitleByUserId(userId interface{}) []*ReplyGroup {
	var replyGroups []*ReplyGroup
	//DB.Raw("select a.*,b.* from reply_group a left join reply_item b on a.id=b.group_id where a.user_id=? ", userId).Scan(&replyGroups)
	var replyItems []*ReplyItem
	DB.Where("user_id = ?", userId).Find(&replyGroups)
	DB.Select("item_name,group_id").Where("user_id = ?", userId).Find(&replyItems)
	temp := make(map[string]*ReplyGroup)
	for _, replyGroup := range replyGroups {
		replyGroup.Items = make([]*ReplyItem, 0)
		temp[replyGroup.Id] = replyGroup
	}
	for _, replyItem := range replyItems {
		temp[replyItem.GroupId].Items = append(temp[replyItem.GroupId].Items, replyItem)
	}
	return replyGroups
}
func CreateReplyGroup(groupName string, userId string) {
	g := &ReplyGroup{
		GroupName: groupName,
		UserId:    userId,
	}
	DB.Create(g)
}
func CreateReplyContent(groupId string, userId string, content, itemName string) {
	g := &ReplyItem{
		GroupId:  groupId,
		UserId:   userId,
		Content:  content,
		ItemName: itemName,
	}
	DB.Create(g)
}
func UpdateReplyContent(id, userId, title, content string) {
	r := &ReplyItem{
		ItemName: title,
		Content:  content,
	}
	DB.Model(&ReplyItem{}).Where("user_id = ? and id = ?", userId, id).Update(r)
}
func DeleteReplyContent(id string, userId string) {
	DB.Where("user_id = ? and id = ?", userId, id).Delete(ReplyItem{})
}
func DeleteReplyGroup(id string, userId string) {
	DB.Where("user_id = ? and id = ?", userId, id).Delete(ReplyGroup{})
	DB.Where("user_id = ? and group_id = ?", userId, id).Delete(ReplyItem{})
}
func FindReplyBySearcch(userId interface{}, search string) []*ReplyGroup {
	var replyGroups []*ReplyGroup
	var replyItems []*ReplyItem
	DB.Where("user_id = ?", userId).Find(&replyGroups)
	DB.Where("user_id = ? and content like ?", userId, "%"+search+"%").Find(&replyItems)
	temp := make(map[string]*ReplyGroup)
	for _, replyGroup := range replyGroups {
		replyGroup.Items = make([]*ReplyItem, 0)
		temp[replyGroup.Id] = replyGroup
	}
	for _, replyItem := range replyItems {
		temp[replyItem.GroupId].Items = append(temp[replyItem.GroupId].Items, replyItem)
	}
	var newReplyGroups []*ReplyGroup = make([]*ReplyGroup, 0)
	for _, replyGroup := range replyGroups {
		if len(replyGroup.Items) != 0 {
			newReplyGroups = append(newReplyGroups, replyGroup)
		}
	}
	return newReplyGroups
}
