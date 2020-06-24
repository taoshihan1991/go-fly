package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/taoshihan1991/imaptool/config"
)
var DB *gorm.DB

func init(){
	mysqlInfo:=config.GetMysql()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlInfo["Username"], mysqlInfo["Password"], mysqlInfo["Server"], mysqlInfo["Port"], mysqlInfo["Database"])
	DB,_=gorm.Open("mysql",dsn)
	DB.SingularTable(true)
	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}
func CloseDB() {
	defer DB.Close()
}