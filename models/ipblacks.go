package models

import "time"

type Ipblack struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	IP       string    `json:"ip"`
	KefuId   string    `json:"kefu_id"`
	CreateAt time.Time `json:"create_at"`
}

func CreateIpblack(ip string, kefuId string) uint {
	black := &Ipblack{
		IP:       ip,
		KefuId:   kefuId,
		CreateAt: time.Now(),
	}
	DB.Create(black)
	return black.ID
}
func DeleteIpblackByIp(ip string) {
	DB.Where("ip = ?", ip).Delete(Ipblack{})
}
func FindIp(ip string) Ipblack {
	var ipblack Ipblack
	DB.Where("ip = ?", ip).First(&ipblack)
	return ipblack
}
func FindIpsByKefuId(id string) []Ipblack {
	var ipblack []Ipblack
	DB.Where("kefu_id = ?", id).Find(&ipblack)
	return ipblack
}
func FindIps(query interface{}, args []interface{}, page uint, pagesize uint) []Ipblack {
	offset := (page - 1) * pagesize
	if offset < 0 {
		offset = 0
	}
	var ipblacks []Ipblack
	if query != nil {
		DB.Where(query, args...).Offset(offset).Limit(pagesize).Find(&ipblacks)
	} else {
		DB.Offset(offset).Limit(pagesize).Find(&ipblacks)
	}
	return ipblacks
}

//查询条数
func CountIps(query interface{}, args []interface{}) uint {
	var count uint
	if query != nil {
		DB.Model(&Visitor{}).Where(query, args...).Count(&count)
	} else {
		DB.Model(&Visitor{}).Count(&count)
	}
	return count
}
