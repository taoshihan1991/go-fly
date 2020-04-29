package main

import (
	"fmt"
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	log.Println("listen on 8080...")
	http.HandleFunc("/", index)
	//登陆界面
	http.HandleFunc("/login", login)
	//监听端口
	http.ListenAndServe(":8080", nil)
}

//输出首页
func index(w http.ResponseWriter, r *http.Request) {
	auth := getCookie(r, "auth")
	if !strings.Contains(auth,"|"){
		http.Redirect(w, r, "/login", 302)
	}else {
		authStrings := strings.Split(auth, "|")
		res := tools.CheckEmailPassword(authStrings[0], authStrings[1], authStrings[2])
		if res{
			folders := tools.GetFolders(authStrings[0], authStrings[1], authStrings[2])
			t, _ := template.ParseFiles("./tmpl/index.html")
			t.Execute(w, folders)
		}else{
			http.Redirect(w, r, "/login", 302)
		}
	}
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
