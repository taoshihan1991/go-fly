package tmpl

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"net/http"
)

//登陆界面
func PageLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

//咨询界面
func PageChat(c *gin.Context) {
	c.HTML(http.StatusOK, "chat_page.html", nil)
}
func RenderLogin(w http.ResponseWriter, render interface{}) {
	html := tools.FileGetContent("html/login.html")
	t, _ := template.New("login").Parse(html)
	t.Execute(w, render)
}
