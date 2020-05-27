package main

import (
	"encoding/json"
	"fmt"
	"github.com/taoshihan1991/imaptool/tmpl"
	"github.com/taoshihan1991/imaptool/tools"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const PageSize = 20

func main() {
	log.Println("listen on 8080...\r\ngo：http://127.0.0.1:8080")
	//根路径
	http.HandleFunc("/", index)
	//邮件夹
	http.HandleFunc("/list", list)
	//登陆界面
	http.HandleFunc("/login", login)
	//验证接口
	http.HandleFunc("/check", check)
	//邮件夹接口
	http.HandleFunc("/folders", folders)
	//邮件接口
	http.HandleFunc("/mail", mail)
	//详情界面
	http.HandleFunc("/view", view)
	//写信界面
	http.HandleFunc("/write", write)
	//发送邮件接口
	http.HandleFunc("/send", send)
	//监听端口
	http.ListenAndServe(":8080", nil)
}

//首页跳转
func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() == "/favicon.ico" {
		return
	}

	mailServer := tools.GetMailServerFromCookie(r)
	if mailServer == nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		res := tools.CheckEmailPassword(mailServer.Server, mailServer.Email, mailServer.Password)
		if res {
			http.Redirect(w, r, "/list", 302)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	}
}

//输出列表
func list(w http.ResponseWriter, r *http.Request) {
	fid:=tools.GetUrlArg(r,"fid")
	currentPage, _ :=strconv.Atoi(tools.GetUrlArg(r,"page"))
	if fid == "" {
		fid = "INBOX"
	}
	if currentPage == 0 {
		currentPage = 1
	}
	render := new(tools.IndexData)
	render.CurrentPage = currentPage
	render.Fid = fid
	tmpl.RenderList(w, render)
}

//详情界面
func view(w http.ResponseWriter, r *http.Request) {
	fid:=tools.GetUrlArg(r,"fid")
	id, _ :=strconv.Atoi(tools.GetUrlArg(r,"id"))
	//
	//mailServer:=tools.GetMailServerFromCookie(r)
	//var wg sync.WaitGroup
	var render = new(tmpl.ViewHtml)
	render.Fid = fid
	render.Id = uint32(id)
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	folders := tools.GetFolders(mailServer.Server, mailServer.Email, mailServer.Password, fid)
	//	render.Folders = folders
	//	render.Fid = fid
	//}()
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	mail := tools.GetMessage(mailServer.Server, mailServer.Email, mailServer.Password, fid, id)
	//	render.From = mail.From
	//	render.To = mail.To
	//	render.Subject = mail.Subject
	//	render.Date = mail.Date
	//	render.HtmlBody = template.HTML(mail.Body)
	//}()
	//wg.Wait()
	tmpl.RenderView(w, render)
}

//登陆界面
func login(w http.ResponseWriter, r *http.Request) {
	tmpl.RenderLogin(w, nil)
}
//写信界面
func write(w http.ResponseWriter, r *http.Request) {
	render:=new(tmpl.CommonHtml)
	tmpl.RenderWrite(w, render)
}
//验证接口
func check(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	server := r.PostFormValue("server")
	password := r.PostFormValue("password")
	msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "验证失败"})

	w.Header().Set("content-type", "text/json;charset=utf-8;")
	if email != "" && server != "" && password != "" {
		res := tools.CheckEmailPassword(server, email, password)
		if res {
			msg, _ = json.Marshal(tools.JsonResult{Code: 200, Msg: "验证成功,正在跳转..."})
			auth := fmt.Sprintf("%s|%s|%s", server, email, password)
			tools.SetCookie("auth", auth, &w)
			w.Write(msg)
		} else {
			w.Write(msg)
		}
	} else {
		w.Write(msg)
	}
}
//发送邮件接口
func send(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "text/json;charset=utf-8;")
	mailServer := tools.GetMailServerFromCookie(r)

	if mailServer == nil {
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "验证失败"})
		w.Write(msg)
		return
	}

	bodyBytes,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "操作失败,"+err.Error()})
		w.Write(msg)
		return
	}
	var sendData tools.SmtpBody
	err = json.Unmarshal(bodyBytes, &sendData)
	if err!=nil{
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "操作失败,"+err.Error()})
		w.Write(msg)
		return
	}

	smtpServer:=sendData.Smtp
	smtpFrom:=sendData.From
	smtpTo:=sendData.To
	smtpBody:=sendData.Body
	smtpPass:=sendData.Password
	smtpSubject:=sendData.Subject
	err=tools.Send(smtpServer,smtpFrom,smtpPass,smtpTo,smtpSubject,smtpBody)
	if err!=nil{
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: err.Error()})
		w.Write(msg)
		return
	}
	msg, _ := json.Marshal(tools.JsonResult{Code: 200, Msg: "发送成功!"})
	w.Write(msg)
}
//邮件夹接口
func folders(w http.ResponseWriter, r *http.Request) {
	fid:=tools.GetUrlArg(r,"fid")
	currentPage, _ :=strconv.Atoi(tools.GetUrlArg(r,"page"))

	if fid == "" {
		fid = "INBOX"
	}
	if currentPage == 0 {
		currentPage = 1
	}

	mailServer := tools.GetMailServerFromCookie(r)
	w.Header().Set("content-type", "text/json;charset=utf-8;")

	if mailServer == nil {
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "验证失败"})
		w.Write(msg)
		return
	}
	var wg sync.WaitGroup
	wg.Add(2)
	result := make(map[string]interface{})
	go func() {
		defer wg.Done()
		folders := tools.GetFolders(mailServer.Server, mailServer.Email, mailServer.Password, fid)
		result["folders"] = folders
		result["total"] = folders[fid]
	}()
	go func() {
		defer wg.Done()
		mails := tools.GetFolderMail(mailServer.Server, mailServer.Email, mailServer.Password, fid, currentPage, PageSize)
		result["mails"] = mails
	}()
	wg.Wait()
	result["pagesize"] = PageSize
	result["fid"] = fid

	msg, _ := json.Marshal(tools.JsonFolders{
		JsonResult: tools.JsonResult{Code: 200, Msg: "获取成功"},
		Result:     result,
	})
	w.Write(msg)
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

	msg, _ := json.Marshal(tools.JsonFolders{
		JsonResult: tools.JsonResult{Code: 200, Msg: "获取成功"},
		Result:     result,
	})
	w.Write(msg)
}
