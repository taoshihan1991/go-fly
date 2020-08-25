package models

import "time"

type Welcome struct {
	ID        uint `gorm:"primary_key" json:"id"`
	UserId string `json:"user_id"`
	Content string `json:"content"`
	IsDefault uint `json:"is_default"`
	Ctime time.Time `json:"ctime"`
}
func FindWelcomeByUserId(userId interface{})Welcome{
	var w Welcome
	DB.Where("user_id = ? and is_default=?", userId,1).First(&w)
	return w
}