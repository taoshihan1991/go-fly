package router

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/middleware"
	"github.com/taoshihan1991/imaptool/tmpl"
)

func InitViewRouter(engine *gin.Engine){
	engine.GET("/index", tmpl.PageIndex)
	engine.GET("/login", tmpl.PageLogin)
	engine.GET("/chat_page",middleware.SetLanguage, tmpl.PageChat)
	engine.GET("/chatIndex",middleware.SetLanguage, tmpl.PageChat)
	engine.GET("/main",middleware.JwtPageMiddleware,tmpl.PageMain)
	engine.GET("/chat_main",middleware.JwtPageMiddleware,tmpl.PageChatMain)
	engine.GET("/setting", tmpl.PageSetting)
	engine.GET("/setting_mysql", tmpl.PageSettingMysql)
	engine.GET("/setting_welcome", tmpl.PageSettingWelcome)
	engine.GET("/setting_deploy", tmpl.PageSettingDeploy)
	engine.GET("/setting_kefu_list",tmpl.PageKefuList)
	engine.GET("/mail_list", tmpl.PageMailList)
	engine.GET("/roles_list", tmpl.PageRoleList)
	engine.GET("/webjs", tmpl.PageWebJs)
	engine.GET("/webcss", tmpl.PageWebCss)
}
