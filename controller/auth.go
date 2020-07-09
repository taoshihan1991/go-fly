package controller

import (
	"github.com/taoshihan1991/imaptool/config"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
)

func CheckPass(username string, password string) string {
	account := config.GetAccount()
	if account == nil {
		account = make(map[string]string)
	}
	if account["Username"] == "" && account["Password"] == "" {
		account["Username"] = "admin"
		account["Password"] = "admin123"
	}
	if username == account["Username"] && password == account["Password"] {

		sessionId := tools.Md5(username)
		info := make(map[string]string)
		info["username"] = username
		config.SetUserInfo(sessionId, info)
		return sessionId
	}
	return ""
}
func CheckKefuPass(username string, password string) (models.User,models.User_role,bool) {
	info:=models.FindUser(username)
	var uRole models.User_role
	if info.Name==""||info.Password!=tools.Md5(password){
		return info,uRole,false
	}
	uRole=models.FindRoleByUserId(info.ID)

	return info,uRole,true
}
func AuthLocal(username string, password string) string {
	account := config.GetAccount()
	if account == nil {
		account = make(map[string]string)
	}
	if account["Username"] == "" && account["Password"] == "" {
		account["Username"] = "admin"
		account["Password"] = "admin123"
	}
	if username == account["Username"] && password == account["Password"] {

		sessionId := tools.Md5(username)
		info := make(map[string]string)
		info["username"] = username
		config.SetUserInfo(sessionId, info)
		return sessionId
	}
	return ""
}

//验证是否已经登录
func AuthCheck(uid string) map[string]string {
	info := config.GetUserInfo(uid)

	return info
}
