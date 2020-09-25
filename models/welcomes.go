package models

import "time"

type Welcome struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserId    string    `json:"user_id"`
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
	}
	DB.Create(w)
	return w.ID
}
func FindWelcomeByUserId(userId interface{}) Welcome {
	var w Welcome
	DB.Where("user_id = ? and is_default=?", userId, 1).First(&w)
	return w
}
func FindWelcomesByUserId(userId interface{}) []Welcome {
	var w []Welcome
	DB.Where("user_id = ?", userId).Find(&w)
	return w
}
func DeleteWelcome(userId interface{}, id string) {
	DB.Where("user_id = ? and id = ?", userId, id).Delete(Welcome{})
}
