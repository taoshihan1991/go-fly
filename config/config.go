package config

import (
	"encoding/json"
	"github.com/taoshihan1991/imaptool/tools"
	"io/ioutil"
)

const Dir = "config/"
const AccountConf = Dir +"account.json"

func GetAccount()map[string]string{
	var account map[string]string
	isExist,_:=tools.IsFileExist(AccountConf)
	if !isExist{
		return account
	}
	info,err:=ioutil.ReadFile(AccountConf)
	if err!=nil{
		return account
	}

	err=json.Unmarshal(info,&account)
	return account
}
