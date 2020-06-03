package controller

import "github.com/taoshihan1991/imaptool/config"

func AuthLocal(username string,password string)bool{
	account:=config.GetAccount()
	if username==account["Username"] && password==account["Password"]{
		return true
	}
	return false
}

