package main

import (
	"encoding/json"
	"github.com/taoshihan1991/imaptool/controller"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)
func main() {
	log.Println("listen on 8080...\r\ngo：http://127.0.0.1:8080")
	mux:=&http.ServeMux{}
	//根路径
	mux.HandleFunc("/", controller.ActionIndex)
	//邮件夹
	mux.HandleFunc("/list", controller.ActionFolder)
	//登陆界面
	mux.HandleFunc("/login", controller.ActionLogin)
	//验证接口
	mux.HandleFunc("/check", controller.LoginCheck)
	//邮件夹接口
	mux.HandleFunc("/folders", controller.FoldersList)
	//新邮件夹接口
	mux.HandleFunc("/folder_dirs", controller.FolderDir)
	//邮件接口
	mux.HandleFunc("/mail", mail)
	//详情界面
	mux.HandleFunc("/view", controller.ActionDetail)
	//写信界面
	mux.HandleFunc("/write", controller.ActionWrite)
	//框架界面
	mux.HandleFunc("/main", controller.ActionMain)
	//设置界面
	mux.HandleFunc("/setting", controller.ActionSetting)
	//设置账户接口
	mux.HandleFunc("/setting_account", controller.SettingAccount)
	//发送邮件接口
	mux.HandleFunc("/send", controller.FolderSend)
	//聊天界面
	mux.HandleFunc("/chat_main", controller.ActionMain)
	//监听端口
	//http.ListenAndServe(":8080", nil)
	//var myHandler http.Handler
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

//邮件接口
func mail(w http.ResponseWriter, r *http.Request) {
	fid:=tools.GetUrlArg(r,"fid")
	id, _ :=strconv.Atoi(tools.GetUrlArg(r,"id"))
	mailServer := tools.GetMailServerFromCookie(r)
	w.Header().Set("content-type", "text/json;charset=utf-8;")

	if mailServer == nil {
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "验证失败"})
		w.Write(msg)
		return
	}
	var wg sync.WaitGroup
	result := make(map[string]interface{})
	wg.Add(2)
	go func() {
		defer wg.Done()
		folders := tools.GetFolders(mailServer.Server, mailServer.Email, mailServer.Password, fid)
		result["folders"] = folders
		result["total"] = folders[fid]
	}()
	go func() {
		defer wg.Done()
		mail := tools.GetMessage(mailServer.Server, mailServer.Email, mailServer.Password, fid, uint32(id))
		result["from"] = mail.From
		result["to"] = mail.To
		result["subject"] = mail.Subject
		result["date"] = mail.Date
		result["html"] = mail.Body
	}()
	wg.Wait()
	result["fid"] = fid

	msg, _ := json.Marshal(tools.JsonListResult{
		JsonResult: tools.JsonResult{Code: 200, Msg: "获取成功"},
		Result:     result,
	})
	w.Write(msg)
}
