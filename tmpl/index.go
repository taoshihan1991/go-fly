package tmpl

import (
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"net/http"
)

func RenderList(w http.ResponseWriter, render interface{}) {
	header := tools.FileGetContent("html/header.html")
	html := tools.FileGetContent("html/list.html")
	t, _ := template.New("list").Parse(html)
	render.(*tools.IndexData).Header=template.HTML(header)
	t.Execute(w, render)
}

