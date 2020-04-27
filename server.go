package main

import (
	"html/template"
	"log"
	"net/http"
)

func main(){
	log.Println("listen on 8080...")
	http.HandleFunc("/",index)
	//登陆界面
	http.HandleFunc("/login",login)
	//监听端口
	http.ListenAndServe(":8080", nil)
}
//输出首页
func index(w http.ResponseWriter, r *http.Request){
	auth:=getCookie(r,"auth")
	if auth=="" {
		http.Redirect(w, r, "/login", 302)
	}
	t,_:=template.ParseFiles("./tmpl/index.html")
	t.Execute(w, nil)
}
//登陆界面
func login(w http.ResponseWriter, r *http.Request){
	t,_:=template.ParseFiles("./tmpl/login.html")
	t.Execute(w, nil)
}
//加密cookie
//func authCookie(){
//
//}
//获取cookie
func getCookie(r *http.Request,name string)string{
	cookies:=r.Cookies()
	for _,cookie:=range cookies{
		if cookie.Name==name{
			return cookie.Value
		}
	}
	return ""
}
