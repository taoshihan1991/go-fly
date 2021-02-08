package models

import "time"

type User_client struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Kefu       string `json:"kefu"`
	Client_id  string `json:"client_id"`
	Created_at string `json:"created_at"`
}

func CreateUserClient(kefu, clientId string) uint {
	u := &User_client{
		Kefu:       kefu,
		Client_id:  clientId,
		Created_at: time.Now().Format("2006-01-02 15:04:05"),
	}
	DB.Create(u)
	return u.ID
}
func FindClients(kefu string) []User_client {
	var arr []User_client
	DB.Where("kefu = ?", kefu).Find(&arr)
	return arr
}
