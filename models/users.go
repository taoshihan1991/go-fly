package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type User struct {
	gorm.Model
	Name string
	Password string
	Nickname string
	Avator string
}
func CreateUser(name string,password string){
	user:=&User{
		Name:name,
		Password: password,
	}
	DB.Create(user)
}
func FindUser(username string)User{
	var user User
	DB.Where("name = ?", username).First(&user)
	return user
}
func FindUserById(id interface{})User{
	var user User
	DB.Where("id = ?", id).First(&user)
	return user
}