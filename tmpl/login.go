package tmpl

import (
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"net/http"
)

func RenderLogin(w http.ResponseWriter, render interface{}) {
	html:=tools.FileGetContent("html/login.html")
	t, _ := template.New("login").Parse(html)
	t.Execute(w, render)
}
