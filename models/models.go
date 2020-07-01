package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/taoshihan1991/imaptool/config"
	"time"
)
var DB *gorm.DB
type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
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