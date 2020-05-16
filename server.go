package main

import (
	"fmt"
	"github.com/taoshihan1991/imaptool/tmpl"
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

const PAGE_SIZE = 20

func main() {
	log.Println("listen on 8080...")
	http.HandleFunc("/", index)
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

	auth := getCookie(r, "auth")
	if !strings.Contains(auth, "|") {
		http.Redirect(w, r, "/login", 302)
	} else {
		authStrings := strings.Split(auth, "|")
		res := tools.CheckEmailPassword(authStrings[0], authStrings[1], authStrings[2])
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

	auth := getCookie(r, "auth")
	authStrings := strings.Split(auth, "|")

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
		folders := tools.GetFolders(authStrings[0], authStrings[1], authStrings[2], fid)
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
		mails := tools.GetFolderMail(authStrings[0], authStrings[1], authStrings[2], fid, currentPage, PAGE_SIZE)
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

	auth := getCookie(r, "auth")
	authStrings := strings.Split(auth, "|")
	var wg sync.WaitGroup
	var render = new(tools.ViewData)
	wg.Add(1)
	go func() {
		defer wg.Done()
		folders := tools.GetFolders(authStrings[0], authStrings[1], authStrings[2], fid)
		render.Folders = folders
		render.Fid = fid
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		mail := tools.GetMessage(authStrings[0], authStrings[1], authStrings[2], fid, id)
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
//获取cookie
func getCookie(r *http.Request, name string) string {
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}
