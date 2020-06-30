package tmpl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//设置界面
func PageSetting(c *gin.Context) {
	c.HTML(http.StatusOK, "setting.html", gin.H{
		"tab_index":"2-3",
		"action":"setting",
	})
}
//设置mysql
func PageSettingMysql(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_mysql.html", gin.H{
		"tab_index":"2-4",
		"action":"setting_mysql",
	})
}
func PageKefuList(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_kefu_list.html", gin.H{
		"tab_index":"1-2",
		"action":"setting_kefu_list",
	})
}
type SettingHtml struct {
	*CommonHtml
	Username, Password string
}

func NewSettingHtml(w http.ResponseWriter) *SettingHtml {
	obj := new(SettingHtml)
	parent := NewRender(w)
	obj.CommonHtml = parent
	return obj
}
