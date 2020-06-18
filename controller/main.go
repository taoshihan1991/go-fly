package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/tmpl"
	"github.com/taoshihan1991/imaptool/tools"
	"net/http"
)

func ActionMain(w http.ResponseWriter, r *http.Request) {
	sessionId := tools.GetCookie(r, "session_id")
	info := AuthCheck(sessionId)
	if len(info) == 0 {
		http.Redirect(w, r, "/login", 302)
		return
	}
	render := tmpl.NewRender(w)
	render.Display("main", render)
}
func MainCheckAuth(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "验证成功",
	})
}
