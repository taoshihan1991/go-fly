package controller

import (
	"encoding/json"
	"fmt"
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"net/http"
)
//登陆界面
func ActionLogin(w http.ResponseWriter, r *http.Request){
	html := tools.FileGetContent("html/login.html")
	t, _ := template.New("login").Parse(html)
	t.Execute(w, nil)
}
//验证接口
func LoginCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json;charset=utf-8;")
	msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "验证失败"})
	authType := r.PostFormValue("type")
	password := r.PostFormValue("password")
	switch authType {
	case "local":
		username := r.PostFormValue("username")
		if AuthLocal(username,password){
			msg, _ = json.Marshal(tools.JsonResult{Code: 200, Msg: "验证成功,正在跳转..."})
			w.Write(msg)
			return
		}
	default:
		email := r.PostFormValue("email")
		server := r.PostFormValue("server")
		if email != "" && server != "" && password != "" {
			res := tools.CheckEmailPassword(server, email, password)
			if res {
				msg, _ = json.Marshal(tools.JsonResult{Code: 200, Msg: "验证成功,正在跳转..."})
				auth := fmt.Sprintf("%s|%s|%s", server, email, password)
				tools.SetCookie("auth", auth, &w)
				w.Write(msg)
				return
			}
		}
	}
	w.Write(msg)
}
