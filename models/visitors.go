package models

import "github.com/jinzhu/gorm"

type Visitor struct {
	gorm.Model
	Name string
	Avator string
	SourceIp string
	ToId string
	VisitorId string
}
func CreateVisitor(name string,avator string,sourceIp string,toId string,visitorId string){
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
	}
	DB.Create(v)
}
func FindVisitorByVistorId(visitorId string)Visitor{
	var v Visitor
	DB.Where("visitor_id = ?", visitorId).First(&v)
	return v
}