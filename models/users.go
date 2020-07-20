package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type User struct {
	Model
	Name string `json:"name"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avator string `json:"avator"`
	RoleName string `json:"role_name"`
}
func CreateUser(name string,password string,avator string,nickname string){
	user:=&User{
		Name:name,
		Password: password,
		Avator:avator,
		Nickname: nickname,
	}
	DB.Create(user)
}
func UpdateUser(id string,name string,password string,avator string,nickname string){
	user:=&User{
		Name:name,
		Avator:avator,
		Nickname: nickname,
	}
	if password!=""{
		user.Password=password
	}
	DB.Model(&User{}).Where("id = ?",id).Update(user)
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
func DeleteUserById(id string){
	DB.Where("id = ?",id).Delete(User{})
}
func FindUsers()[]User{
	var users []User
	DB.Order("id desc").Find(&users)
	return users
}
func FindUserRole(query interface{},id interface{})User{
	var user User
	DB.Select(query).Where("user.id = ?", id).Joins("join user_role on user.id=user_role.user_id").Joins("join role on user_role.role_id=role.id").First(&user)
	return user
}