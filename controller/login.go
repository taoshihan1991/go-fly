package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"log"
	"net/http"
)

//验证接口
func LoginCheckPass(c *gin.Context) {
	authType := c.PostForm("type")
	password := c.PostForm("password")
	username := c.PostForm("username")
	switch authType {
	case "local":
		sessionId := CheckPass(username, password)
		userinfo := make(map[string]interface{})
		userinfo["name"] = username
		token, err := tools.MakeToken(userinfo)
		log.Println(err)
		if sessionId != "" {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "验证成功,正在跳转",
				"result": gin.H{
					"token": token,
				},
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "验证失败",
		})
	}
}
func ActionLogin(w http.ResponseWriter, r *http.Request) {
	html := tools.FileGetContent("html/login.html")
	t, _ := template.New("login").Parse(html)
	t.Execute(w, nil)
}

//验证接口
func LoginCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json;charset=utf-8;")
	msg, _ := json.Marshal(tools.JsonResult{Code: 400, Msg: "验证失败"})
	authType := r.PostFormValue("type")
	password := r.PostFormValue("password")
	switch authType {
	case "local":
		username := r.PostFormValue("username")
		sessionId := AuthLocal(username, password)
		if sessionId != "" {
			tools.SetCookie("session_id", sessionId, &w)
			msg, _ = json.Marshal(tools.JsonResult{Code: 200, Msg: "验证成功,正在跳转..."})
			w.Write(msg)
			return
		}
	default:
		email := r.PostFormValue("email")
		server := r.PostFormValue("server")
		if email != "" && server != "" && password != "" {
			res := tools.CheckEmailPassword(server, email, password)
			if res {
				msg, _ = json.Marshal(tools.JsonResult{Code: 200, Msg: "验证成功,正在跳转..."})
				auth := fmt.Sprintf("%s|%s|%s", server, email, password)
				tools.SetCookie("auth", auth, &w)
				w.Write(msg)
				return
			}
		}
	}
	w.Write(msg)
}
