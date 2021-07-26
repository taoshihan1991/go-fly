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
	mysql:=config.CreateMysql()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysql.Username, mysql.Password, mysql.Server, mysql.Port, mysql.Database)
	DB,_=gorm.Open("mysql",dsn)
	DB.SingularTable(true)
	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	if err := DB.DB().Ping(); err != nil {
		panic(err)
	}

	InitConfig()
}
func Execute(sql string){
	DB.Exec(sql)
}
func CloseDB() {
	defer DB.Close()
}