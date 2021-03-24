package models

import (
	"time"
)

type Visitor struct {
	Model
	Name      string `json:"name"`
	Avator    string `json:"avator"`
	SourceIp  string `json:"source_ip"`
	ToId      string `json:"to_id"`
	VisitorId string `json:"visitor_id"`
	Status    uint   `json:"status"`
	Refer     string `json:"refer"`
	City      string `json:"city"`
	ClientIp  string `json:"client_ip"`
	Extra     string `json:"extra"`
}

func CreateVisitor(name, avator, sourceIp, toId, visitorId, refer, city, clientIp, extra string) {
	v := &Visitor{
		Name:      name,
		Avator:    avator,
		SourceIp:  sourceIp,
		ToId:      toId,
		VisitorId: visitorId,
		Status:    1,
		Refer:     refer,
		City:      city,
		ClientIp:  clientIp,
		Extra:     extra,
	}
	v.UpdatedAt = time.Now()
	DB.Create(v)
}
func FindVisitorByVistorId(visitorId string) Visitor {
	var v Visitor
	DB.Where("visitor_id = ?", visitorId).First(&v)
	return v
}
func FindVisitors(page uint, pagesize uint) []Visitor {
	offset := (page - 1) * pagesize
	if offset < 0 {
		offset = 0
	}
	var visitors []Visitor
	DB.Offset(offset).Limit(pagesize).Order("status desc, updated_at desc").Find(&visitors)
	return visitors
}
func FindVisitorsByKefuId(page uint, pagesize uint, kefuId string) []Visitor {
	offset := (page - 1) * pagesize
	if offset <= 0 {
		offset = 0
	}
	var visitors []Visitor
	//sql := fmt.Sprintf("select * from visitor where id>=(select id from visitor where  to_id='%s' order by updated_at desc limit %d,1) and to_id='%s' order by updated_at desc limit %d ", kefuId, offset, kefuId, pagesize)
	//DB.Raw(sql).Scan(&visitors)
	DB.Where("to_id=?", kefuId).Offset(offset).Limit(pagesize).Order("updated_at desc").Find(&visitors)
	return visitors
}
func FindVisitorsOnline() []Visitor {
	var visitors []Visitor
	DB.Where("status = ?", 1).Find(&visitors)
	return visitors
}
func UpdateVisitorStatus(visitorId string, status uint) {
	visitor := Visitor{}
	DB.Model(&visitor).Where("visitor_id = ?", visitorId).Update("status", status)
}
func UpdateVisitor(name, avator, visitorId string, status uint, clientIp string, sourceIp string, refer, extra string) {
	visitor := &Visitor{
		Status:   status,
		ClientIp: clientIp,
		SourceIp: sourceIp,
		Refer:    refer,
		Extra:    extra,
		Name:     name,
		Avator:   avator,
	}
	visitor.UpdatedAt = time.Now()
	DB.Model(visitor).Where("visitor_id = ?", visitorId).Update(visitor)
}
func UpdateVisitorKefu(visitorId string, kefuId string) {
	visitor := Visitor{}
	DB.Model(&visitor).Where("visitor_id = ?", visitorId).Update("to_id", kefuId)
}

//查询条数
func CountVisitors() uint {
	var count uint
	DB.Model(&Visitor{}).Count(&count)
	return count
}

//查询条数
func CountVisitorsByKefuId(kefuId string) uint {
	var count uint
	DB.Model(&Visitor{}).Where("to_id=?", kefuId).Count(&count)
	return count
}
