package tmpl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//设置界面
func PageSetting(c *gin.Context) {
	c.HTML(http.StatusOK, "setting.html", gin.H{
		"tab_index":"1-1",
		"action":"setting",
	})
}
//设置欢迎
func PageSettingWelcome(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_welcome.html", gin.H{
		"tab_index":"1-2",
		"action":"setting_welcome",
	})
}
//设置mysql
func PageSettingMysql(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_mysql.html", gin.H{
		"tab_index":"2-4",
		"action":"setting_mysql",
	})
}
//设置部署
func PageSettingDeploy(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_deploy.html", gin.H{
		"tab_index":"2-5",
		"action":"setting_deploy",
	})
}
//前台js部署
func PageWebJs(c *gin.Context){
	c.HTML(http.StatusOK, "chat_web.js",nil)
}
//前台css部署
func PageWebCss(c *gin.Context){
	c.HTML(http.StatusOK, "chat_web.css",nil)
}
func PageKefuList(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_kefu_list.html", gin.H{
		"tab_index":"3-2",
		"action":"setting_kefu_list",
	})
}
//角色列表
func PageRoleList(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_role_list.html", gin.H{
		"tab_index":"3-1",
		"action":"roles_list",
	})
}
//角色列表
func PageIpblack(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_ipblack.html", gin.H{
		"tab_index":"4-5",
		"action":"setting_ipblack",
	})
}
//配置项列表
func PageConfig(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_config.html", gin.H{
		"tab_index":"4-6",
		"action":"setting_config",
	})
}


