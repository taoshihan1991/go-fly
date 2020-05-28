package controller

import (
	"github.com/taoshihan1991/imaptool/tmpl"
	"net/http"
)
func ActionMain(w http.ResponseWriter, r *http.Request){
	render:=tmpl.NewRender(w)
	render.Display("main",render)
}
