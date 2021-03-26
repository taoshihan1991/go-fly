package middleware

import "github.com/gin-gonic/gin"

func CrossSite(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//服务器支持的所有跨域请求的方法
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
	//允许跨域设置可以返回其他子段，可以自定义字段
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
	// 允许浏览器（客户端）可以解析的头部 （重要）
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
	//允许客户端传递校验信息比如 cookie (重要)
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
}
