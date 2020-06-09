package config

import (
	"encoding/json"
	"fmt"
	"github.com/taoshihan1991/imaptool/tools"
	"io/ioutil"
	"os"
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
func GetUserInfo(uid string)map[string]string{
	var userInfo map[string]string
	userFile:=Dir+"sess_"+uid+".json"
	isExist,_:=tools.IsFileExist(userFile)
	if !isExist{
		return userInfo
	}
	info,err:=ioutil.ReadFile(userFile)
	if err!=nil{
		return userInfo
	}

	err=json.Unmarshal(info,&userInfo)
	return userInfo
}
func SetUserInfo(uid string,info map[string]string){
	userFile:=Dir+"sess_"+uid+".json"
	isExist,_:=tools.IsFileExist(Dir)
	if !isExist{
		os.Mkdir(Dir,os.ModePerm)
	}
	file, _ := os.OpenFile(userFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	str:="{\r\n"
	for k,v:=range info{
		str+=fmt.Sprintf(`"%s":"%s",`,k,v)
	}
	str+=fmt.Sprintf(`"session_id":"%s"%s}`,uid,"\r\n")
	file.WriteString(str)
}