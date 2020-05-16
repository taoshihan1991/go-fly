package main

import (
	"fmt"
	"github.com/taoshihan1991/imaptool/tmpl"
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const PageSize = 20

func main() {
	log.Println("listen on 8080...\r\n访问：http://127.0.0.1:8080")
	//根路径
	http.HandleFunc("/", index)
	//邮件夹
	http.HandleFunc("/list", list)
	//登陆界面
	http.HandleFunc("/login", login)
	//详情界面
	http.HandleFunc("/view", view)
	//监听端口
	http.ListenAndServe(":8080", nil)
}

//首页跳转
func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() == "/favicon.ico" {
		return
	}

	mailServer:=tools.GetMailServerFromCookie(r)
	if mailServer==nil {
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
	values := r.URL.Query()
	fid := ""
	currentPage := 0
	if len(values["fid"]) != 0 {
		fid = values["fid"][0]
	}
	if len(values["page"]) != 0 {
		currentPage, _ = strconv.Atoi(values["page"][0])
	}
	if fid == "" {
		fid = "INBOX"
	}
	if currentPage == 0 {
		currentPage = 1
	}

	mailServer:=tools.GetMailServerFromCookie(r)

	render := new(tools.IndexData)
	render.CurrentPage = currentPage
	var prePage int
	if (currentPage - 1) <= 0 {
		prePage = 1
	} else {
		prePage = currentPage - 1
	}
	render.PrePage = fmt.Sprintf("/list?fid=%s&page=%d", fid, prePage)
	render.NextPage = fmt.Sprintf("/list?fid=%s&page=%d", fid, currentPage+1)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		folders := tools.GetFolders(mailServer.Server, mailServer.Email,mailServer.Password, fid)
		render.Folders = folders
		render.Fid = fid

		//PageCount:= render.Folders[fid]/PAGE_SIZE
		numPages := ""
		start := currentPage - 5
		if start <= 0 {
			start = 1
		}
		end := start + 11
		//if end>=PageCount{
		//	end=PageCount
		//}

		for i := start; i < end; i++ {
			active := ""
			if currentPage == i {
				active = "active"
			}
			numPages += fmt.Sprintf("<li class=\"page-item %s\"><a class=\"page-link\" href=\"/list?fid=%s&page=%d\">%d</a></li>", active, fid, i, i)
		}
		render.NumPages = template.HTML(numPages)
	}()
	go func() {
		defer wg.Done()
		mails := tools.GetFolderMail(mailServer.Server, mailServer.Email, mailServer.Password, fid, currentPage, PageSize)
		render.MailPagelist = mails
	}()

	wg.Wait()
	tmpl.RenderList(w, render)
}

//详情界面
func view(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fid := ""
	if len(values["fid"]) != 0 {
		fid = values["fid"][0]
	} else {
		fid = "INBOX"
	}
	var id uint32
	if len(values["id"]) != 0 {
		i, _ := strconv.Atoi(values["id"][0])
		id = uint32(i)
	} else {
		id = 0
	}

	mailServer:=tools.GetMailServerFromCookie(r)
	var wg sync.WaitGroup
	var render = new(tools.ViewData)
	wg.Add(1)
	go func() {
		defer wg.Done()
		folders := tools.GetFolders(mailServer.Server, mailServer.Email, mailServer.Password, fid)
		render.Folders = folders
		render.Fid = fid
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		mail := tools.GetMessage(mailServer.Server, mailServer.Email, mailServer.Password, fid, id)
		render.From = mail.From
		render.To = mail.To
		render.Subject = mail.Subject
		render.Date = mail.Date
		render.HtmlBody = template.HTML(mail.Body)
	}()
	wg.Wait()
	tmpl.RenderView(w, render)
}

//登陆界面
func login(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	server := r.PostFormValue("server")
	password := r.PostFormValue("password")
	var errStr string
	if email != "" && server != "" && password != "" {
		res := tools.CheckEmailPassword(server, email, password)
		if !res {
			errStr = "连接或验证失败"
			tmpl.RenderLogin(w, errStr)
		} else {
			auth := fmt.Sprintf("%s|%s|%s", server, email, password)
			cookie := http.Cookie{
				Name:  "auth",
				Value: auth,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", 302)
		}
	} else {
		tmpl.RenderLogin(w, errStr)
	}
}

//加密cookie
//func authCookie(){
//
//}

