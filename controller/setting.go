package controller

import (
	"encoding/json"
	"fmt"
	"github.com/taoshihan1991/imaptool/tmpl"
	"github.com/taoshihan1991/imaptool/tools"
	"io/ioutil"
	"net/http"
	"os"
)

const configDir = "config/"
const configFile=configDir+"account.json"
func ActionSetting(w http.ResponseWriter, r *http.Request){
	render:=tmpl.NewSettingHtml(w)
	render.SetLeft("setting_left")
	render.SetBottom("setting_bottom")
	account:=getAccount()
	render.Username=account["Username"]
	render.Password=account["Password"]
	render.Display("setting",render)
}
func SettingAccount(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "text/json;charset=utf-8;")
	mailServer := tools.GetMailServerFromCookie(r)

	if mailServer == nil {
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "验证失败"})
		w.Write(msg)
		return
	}

	username:=r.PostFormValue("username")
	password:=r.PostFormValue("password")

	isExist,_:=tools.IsFileExist(configDir)
	if !isExist{
		os.Mkdir(configDir,os.ModePerm)
	}
	fileConfig:=configFile
	file, _ := os.OpenFile(fileConfig, os.O_RDWR|os.O_CREATE, os.ModePerm)

	format:=`{
	"Username":"%s",
	"Password":"%s"
}
`
	data := fmt.Sprintf(format,username,password)
	file.WriteString(data)

	msg, _ := json.Marshal(tools.JsonResult{Code: 200, Msg: "操作成功!"})
	w.Write(msg)
}
func SettingGetAccount(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "text/json;charset=utf-8;")
	mailServer := tools.GetMailServerFromCookie(r)

	if mailServer == nil {
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "验证失败"})
		w.Write(msg)
		return
	}
	result:=getAccount()
	msg, _ := json.Marshal(tools.JsonListResult{
		JsonResult: tools.JsonResult{Code: 200, Msg: "获取成功"},
		Result:     result,
	})
	w.Write(msg)
}
func getAccount()map[string]string{
	var account map[string]string
	isExist,_:=tools.IsFileExist(configFile)
	if !isExist{
		return account
	}
	info,err:=ioutil.ReadFile(configFile)
	if err!=nil{
		return account
	}

	err=json.Unmarshal(info,&account)
	return account
}