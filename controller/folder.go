package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/config"
	"github.com/taoshihan1991/imaptool/tmpl"
	"github.com/taoshihan1991/imaptool/tools"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

const PageSize = 20

func GetFolders(c *gin.Context) {
	fid := c.Query("fid")
	currentPage, _ := strconv.Atoi(c.Query("page"))
	if fid == "" {
		fid = "INBOX"
	}
	if currentPage == 0 {
		currentPage = 1
	}

	mailServer := config.CreateMailServer()

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

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": result,
	})
}
func GetFolderList(c *gin.Context) {
	fid := c.Query("fid")
	if fid == "" {
		fid = "INBOX"
	}

	mailServer := config.CreateMailServer()

	result := make(map[string]interface{})
	folders := tools.GetFolders(mailServer.Server, mailServer.Email, mailServer.Password, fid)
	result["folders"] = folders
	result["total"] = folders[fid]
	result["fid"] = fid

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": result,
	})
}

//输出列表
func ActionFolder(w http.ResponseWriter, r *http.Request) {
	fid := tools.GetUrlArg(r, "fid")
	currentPage, _ := strconv.Atoi(tools.GetUrlArg(r, "page"))
	if fid == "" {
		fid = "INBOX"
	}
	if currentPage == 0 {
		currentPage = 1
	}
	render := tmpl.NewFolderHtml(w)
	render.CurrentPage = currentPage
	render.Fid = fid
	render.Display("list", render)
}

//写信界面
func ActionWrite(w http.ResponseWriter, r *http.Request) {
	render := tmpl.NewRender(w)
	render.SetLeft("mail_left")
	render.Display("write", nil)
}

//读信界面
func ActionDetail(w http.ResponseWriter, r *http.Request) {
	fid := tools.GetUrlArg(r, "fid")
	id, _ := strconv.Atoi(tools.GetUrlArg(r, "id"))

	render := tmpl.NewDetailHtml(w)
	render.SetLeft("mail_left")
	render.Fid = fid
	render.Id = uint32(id)
	render.Display("mail_detail", render)
}

//获取邮件夹接口
func FolderDir(w http.ResponseWriter, r *http.Request) {
	fid := tools.GetUrlArg(r, "fid")

	if fid == "" {
		fid = "INBOX"
	}

	mailServer := tools.GetMailServerFromCookie(r)
	w.Header().Set("content-type", "text/json;charset=utf-8;")

	if mailServer == nil {
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "验证失败"})
		w.Write(msg)
		return
	}
	result := make(map[string]interface{})
	folders := tools.GetFolders(mailServer.Server, mailServer.Email, mailServer.Password, fid)
	result["folders"] = folders
	result["total"] = folders[fid]
	result["fid"] = fid
	msg, _ := json.Marshal(tools.JsonListResult{
		JsonResult: tools.JsonResult{Code: 200, Msg: "获取成功"},
		Result:     result,
	})
	w.Write(msg)
}

//邮件夹接口
func FoldersList(w http.ResponseWriter, r *http.Request) {
	fid := tools.GetUrlArg(r, "fid")
	currentPage, _ := strconv.Atoi(tools.GetUrlArg(r, "page"))

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

	msg, _ := json.Marshal(tools.JsonListResult{
		JsonResult: tools.JsonResult{Code: 200, Msg: "获取成功"},
		Result:     result,
	})
	w.Write(msg)
}

//邮件接口
func FolderMail(w http.ResponseWriter, r *http.Request) {
	fid := tools.GetUrlArg(r, "fid")
	id, _ := strconv.Atoi(tools.GetUrlArg(r, "id"))
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

//发送邮件接口
func FolderSend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json;charset=utf-8;")
	mailServer := tools.GetMailServerFromCookie(r)

	if mailServer == nil {
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "验证失败"})
		w.Write(msg)
		return
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "操作失败," + err.Error()})
		w.Write(msg)
		return
	}
	var sendData tools.SmtpBody
	err = json.Unmarshal(bodyBytes, &sendData)
	if err != nil {
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "操作失败," + err.Error()})
		w.Write(msg)
		return
	}

	smtpServer := sendData.Smtp
	smtpFrom := mailServer.Email
	smtpTo := sendData.To
	smtpBody := sendData.Body
	smtpPass := mailServer.Password
	smtpSubject := sendData.Subject
	err = tools.SendSmtp(smtpServer, smtpFrom, smtpPass, smtpTo, smtpSubject, smtpBody)
	if err != nil {
		msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: err.Error()})
		w.Write(msg)
		return
	}
	msg, _ := json.Marshal(tools.JsonResult{Code: 200, Msg: "发送成功!"})
	w.Write(msg)
}
