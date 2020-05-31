package controller

import (
	"github.com/taoshihan1991/imaptool/tmpl"
	"net/http"
)
//聊天主界面
func ActionChatMain(w http.ResponseWriter, r *http.Request){
	render:=tmpl.NewRender(w)
	render.Display("chat_main",nil)
}
