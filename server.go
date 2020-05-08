package main

import (
	"fmt"
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type IndexData struct {
	Folders map[string]int
	Mails   interface{}
}

func main() {
	log.Println("listen on 8080...")
	http.HandleFunc("/", index)
	http.HandleFunc("/list", list)
	//登陆界面
	http.HandleFunc("/login", login)
	//监听端口
	http.ListenAndServe(":8080", nil)
}

//首页跳转
func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI()=="/favicon.ico"{
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
	values:=r.URL.Query()
	fid:=""
	currentPage:=0
	if len(values["fid"])!=0{
		fid=values["fid"][0]
	}
	if len(values["page"])!=0{
		currentPage,_=strconv.Atoi(values["page"][0])
	}
	if fid==""{
		fid="INBOX"
	}
	if currentPage==0{
		currentPage=1
	}


	auth := getCookie(r, "auth")
	authStrings := strings.Split(auth, "|")

	render := new(IndexData)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		folders :=tools.GetFolders(authStrings[0], authStrings[1], authStrings[2])
		render.Folders = folders
	}()
	mails := tools.GetFolderMail(authStrings[0], authStrings[1], authStrings[2], fid,currentPage, 20)
	render.Mails = mails
	wg.Wait()
	t, _ := template.ParseFiles("./tmpl/index.html")
	t.Execute(w, render)
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
			t, _ := template.ParseFiles("./tmpl/login.html")
			t.Execute(w, errStr)
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
		t, _ := template.ParseFiles("./tmpl/login.html")
		t.Execute(w, nil)
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
