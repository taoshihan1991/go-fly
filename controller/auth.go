package controller

import (
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
)

func CheckKefuPass(username string, password string) (models.User, models.User_role, bool) {
	info := models.FindUser(username)
	var uRole models.User_role
	if info.Name == "" || info.Password != tools.Md5(password) {
		return info, uRole, false
	}
	uRole = models.FindRoleByUserId(info.ID)

	return info, uRole, true
}
