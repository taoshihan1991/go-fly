package tmpl

import (
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"net/http"
)

type CommonHtml struct{
	Header 			  template.HTML
	Nav 			  template.HTML
	Left 			  template.HTML
	Bottom			  template.HTML
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
func (obj *CommonHtml)SetLeft(file string){
	leftStr := tools.FileGetContent("html/"+file+".html")
	obj.Left=template.HTML(leftStr)
}
func (obj *CommonHtml)SetBottom(file string){
	str := tools.FileGetContent("html/"+file+".html")
	obj.Bottom=template.HTML(str)
}
func (obj *CommonHtml)Display(file string,data interface{}){
	if data==nil{
		data=obj
	}
	main := tools.FileGetContent("html/"+file+".html")
	t, _ := template.New(file).Parse(main)
	t.Execute(obj.Rw, data)
}