package tmpl

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"net/http"
)
//登陆界面
func PageLogin(c *gin.Context){
	html := tools.FileGetContent("html/login.html")
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, html)
}
//咨询界面
func PageChat(c *gin.Context){
	html := tools.FileGetContent("html/chat_page.html")
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, html)
}
func RenderLogin(w http.ResponseWriter, render interface{}) {
	html := tools.FileGetContent("html/login.html")
	t, _ := template.New("login").Parse(html)
	t.Execute(w, render)
}
