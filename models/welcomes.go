package models

import "time"

type Welcome struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserId    string    `json:"user_id"`
	Keyword   string    `json:"keyword"`
	Content   string    `json:"content"`
	IsDefault uint      `json:"is_default"`
	Ctime     time.Time `json:"ctime"`
}

func CreateWelcome(userId string, content string) uint {
	if userId == "" || content == "" {
		return 0
	}
	w := &Welcome{
		UserId:  userId,
		Content: content,
		Ctime:   time.Now(),
		Keyword: "welcome",
	}
	DB.Create(w)
	return w.ID
}
func UpdateWelcome(userId string, id string, content string) uint {
	if userId == "" || content == "" {
		return 0
	}
	w := &Welcome{
		Content: content,
	}
	DB.Model(w).Where("user_id = ? and id = ?", userId, id).Update(w)
	return w.ID
}
func FindWelcomeByUserIdKey(userId interface{}, keyword interface{}) Welcome {
	var w Welcome
	DB.Where("user_id = ? and keyword=?", userId, keyword).First(&w)
	return w
}
func FindWelcomesByUserId(userId interface{}) []Welcome {
	var w []Welcome
	DB.Where("user_id = ?", userId).Find(&w)
	return w
}
func FindWelcomesByKeyword(userId interface{}, keyword interface{}) []Welcome {
	var w []Welcome
	DB.Where("user_id = ? and keyword=?", userId, keyword).Find(&w)
	return w
}
func DeleteWelcome(userId interface{}, id string) {
	DB.Where("user_id = ? and id = ?", userId, id).Delete(Welcome{})
}
