package tmpl

import (
	"net/http"
)
type ViewHtml struct{
	CommonHtml
	Fid string
	Id uint32
}
func RenderView(w http.ResponseWriter, data interface{}) {
	render:=NewRender(w)
	data.(*ViewHtml).Nav=render.Nav
	data.(*ViewHtml).Header=render.Header
	render.Display("view",data)
}
