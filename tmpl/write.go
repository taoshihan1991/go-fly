package tmpl

import (
	"net/http"
)
func RenderWrite(w http.ResponseWriter, data interface{}) {
	render:=NewRender(w)
	render.Display("write",render)
}
