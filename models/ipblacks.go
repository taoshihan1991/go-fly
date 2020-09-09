package models

import "time"

type Ipblack struct{
	ID     uint `gorm:"primary_key" json:"id"`
	IP string `json:"ip"`
	KefuId string `json:"kefu_id"`
	CreateAt time.Time `json:"create_at"`
}
func CreateIpblack(ip string,kefuId string)uint{
	black:=&Ipblack{
		IP:ip,
		KefuId: kefuId,
		CreateAt: time.Now(),
	}
	DB.Create(black)
	return black.ID
}