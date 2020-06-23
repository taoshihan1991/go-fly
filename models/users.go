package models
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type User struct {
	gorm.Model
	Id int64
	Name string
	Password string
}

func CreateUser(name string,password string){
	user:=&User{
		Name:name,
		Password: password,
	}
	DB.Create(user)
}
func FindUsers()[]User{
	var users []User
	DB.Find(&users)
	return users
}