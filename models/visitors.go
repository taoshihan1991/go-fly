package models

import "github.com/jinzhu/gorm"

type Visitor struct {
	gorm.Model
	Name string `json:"name"`
	Avator string `json:"avator"`
	SourceIp string `json:"source_ip"`
	ToId string `json:"to_id"`
	VisitorId string `json:"visitor_id"`
	Status uint `json:"status"`
	Refer string `json:"refer"`
	City string `json:"city"`
	ClientIp string `json:"client_ip"`
}
func CreateVisitor(name string,avator string,sourceIp string,toId string,visitorId string,refer string,city string,clientIp string){
	old:=FindVisitorByVistorId(visitorId)
	if old.Name!=""{
		return
	}
	v:=&Visitor{
		Name:name,
		Avator: avator,
		SourceIp:sourceIp,
		ToId:toId,
		VisitorId: visitorId,
		Status:1,
		Refer:refer,
		City:city,
		ClientIp:clientIp,
	}
	DB.Create(v)
}
func FindVisitorByVistorId(visitorId string)Visitor{
	var v Visitor
	DB.Where("visitor_id = ?", visitorId).First(&v)
	return v
}
func FindVisitors()[]Visitor{
	var visitors []Visitor
	DB.Find(&visitors)
	return visitors
}