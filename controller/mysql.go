package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/tmpl"
	"net/http"
)
func MysqlGetConf(c *gin.Context) {

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "验证成功",
	})
}
func ActionMysqlSet(w http.ResponseWriter, r *http.Request) {
	render := tmpl.NewSettingHtml(w)
	render.SetLeft("setting_left")
	render.SetBottom("setting_bottom")
	render.Display("mysql_setting", render)
}
