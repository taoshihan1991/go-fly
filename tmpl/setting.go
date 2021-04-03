package tmpl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//设置界面
func PageSetting(c *gin.Context) {
	c.HTML(http.StatusOK, "setting.html", gin.H{
		"tab_index": "1-1",
		"action":    "setting",
	})
}

//设置欢迎
func PageSettingWelcome(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_welcome.html", gin.H{
		"tab_index": "1-2",
		"action":    "setting_welcome",
	})
}

//统计
func PageSettingStatis(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_statistics.html", gin.H{
		"tab_index": "1-3",
		"action":    "setting_statistics",
	})
}

//设置mysql
func PageSettingMysql(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_mysql.html", gin.H{
		"tab_index": "2-4",
		"action":    "setting_mysql",
	})
}

//设置部署
func PageSettingDeploy(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_deploy.html", gin.H{
		"tab_index": "2-5",
		"action":    "setting_deploy",
	})
}

func PageKefuList(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_kefu_list.html", gin.H{
		"tab_index": "3-2",
		"action":    "setting_kefu_list",
	})
}
func PageAvator(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_avator.html", gin.H{
		"tab_index": "3-2",
		"action":    "setting_avator",
	})
}
func PageModifypass(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_modifypass.html", gin.H{
		"tab_index": "3-2",
		"action":    "setting_modifypass",
	})
}

//角色列表
func PageRoleList(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_role_list.html", gin.H{
		"tab_index": "3-1",
		"action":    "roles_list",
	})
}

//角色列表
func PageIpblack(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_ipblack.html", gin.H{
		"tab_index": "4-5",
		"action":    "setting_ipblack",
	})
}

//配置项列表
func PageConfig(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_config.html", gin.H{
		"tab_index": "4-6",
		"action":    "setting_config",
	})
}

//配置项编辑首页
func PageSettingIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_pageindex.html", gin.H{
		"tab_index": "4-7",
		"action":    "setting_pageindex",
	})
}

//配置项编辑首页
func PageSettingIndexPages(c *gin.Context) {
	c.HTML(http.StatusOK, "setting_indexpages.html", gin.H{
		"tab_index": "4-7",
		"action":    "setting_indexpages",
	})
}
