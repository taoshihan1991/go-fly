package controller

import (
	"github.com/taoshihan1991/imaptool/tmpl"
	"net/http"
)

func ActionMysqlSet(w http.ResponseWriter, r *http.Request) {
	render := tmpl.NewSettingHtml(w)
	render.SetLeft("setting_left")
	render.SetBottom("setting_bottom")
	render.Display("mysql_setting", render)
}
