package router

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/middleware"
	"github.com/taoshihan1991/imaptool/tmpl"
)

func InitViewRouter(engine *gin.Engine) {
	engine.GET("/", tmpl.PageIndex)
	engine.GET("/index_:lang", tmpl.PageIndex)
	engine.GET("/install", tmpl.PageInstall)
	engine.GET("/detail_:page",tmpl.PageDetail)
	engine.GET("/login", tmpl.PageLogin)
	engine.GET("/bind", tmpl.PageBind)
	engine.GET("/chatIndex",tmpl.PageChat)
	engine.GET("/main", middleware.JwtPageMiddleware, tmpl.PageMain)
	engine.GET("/chat_main", middleware.JwtPageMiddleware,middleware.DomainLimitMiddleware, tmpl.PageChatMain)
	engine.GET("/setting",middleware.DomainLimitMiddleware, tmpl.PageSetting)
	engine.GET("/setting_statistics", tmpl.PageSettingStatis)
	engine.GET("/setting_indexpage", tmpl.PageSettingIndexPage)
	engine.GET("/setting_indexpages", tmpl.PageSettingIndexPages)
	engine.GET("/setting_mysql", tmpl.PageSettingMysql)
	engine.GET("/setting_welcome", tmpl.PageSettingWelcome)
	engine.GET("/setting_deploy", tmpl.PageSettingDeploy)
	engine.GET("/setting_kefu_list", tmpl.PageKefuList)
	engine.GET("/setting_avator", tmpl.PageAvator)
	engine.GET("/setting_modifypass", tmpl.PageModifypass)
	engine.GET("/setting_ipblack", tmpl.PageIpblack)
	engine.GET("/setting_config", tmpl.PageConfig)
	engine.GET("/mail_list", tmpl.PageMailList)
	engine.GET("/roles_list", tmpl.PageRoleList)
}
