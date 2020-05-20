package tmpl

import (
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"net/http"
)

func RenderList(w http.ResponseWriter, render interface{}) {
	html:=tools.FileGetContent("html/list.html")
	t, _ := template.New("list").Parse(html)
	t.Execute(w,nil)
}
