package controller

import (
	"github.com/taoshihan1991/imaptool/tools"
	"net/http"
)

//首页跳转
func ActionIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() == "/favicon.ico" {
		return
	}

	mailServer := tools.GetMailServerFromCookie(r)
	if mailServer == nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		res := tools.CheckEmailPassword(mailServer.Server, mailServer.Email, mailServer.Password)
		if res {
			http.Redirect(w, r, "/main", 302)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	}
}
