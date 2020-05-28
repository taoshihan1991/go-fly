package controller

import (
	"github.com/taoshihan1991/imaptool/tmpl"
	"net/http"
)

func ActionSetting(w http.ResponseWriter, r *http.Request){
	render:=tmpl.NewRender(w)
	render.Display("setting",render)
}
