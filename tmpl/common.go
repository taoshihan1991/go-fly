package tmpl

import (
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"net/http"
)

type CommonHtml struct{
	Header 			  template.HTML
	Nav 			  template.HTML
	Rw				  http.ResponseWriter
}
func NewRender(rw http.ResponseWriter)*CommonHtml{
	obj:=new(CommonHtml)
	obj.Rw=rw
	header := tools.FileGetContent("html/header.html")
	nav := tools.FileGetContent("html/nav.html")
	obj.Header=template.HTML(header)
	obj.Nav=template.HTML(nav)
	return obj
}
func (obj *CommonHtml)Display(file string,data interface{}){
	main := tools.FileGetContent("html/"+file+".html")
	t, _ := template.New(file).Parse(main)
	t.Execute(obj.Rw, data)
}