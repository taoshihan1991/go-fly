package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/config"
	"github.com/taoshihan1991/imaptool/database"
	"github.com/taoshihan1991/imaptool/tools"
	"os"
)

func MysqlGetConf(c *gin.Context) {
	mysqlInfo := config.GetMysql()
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "验证成功",
		"result": mysqlInfo,
	})
}
func MysqlSetConf(c *gin.Context) {

	mysqlServer := c.PostForm("server")
	mysqlPort := c.PostForm("port")
	mysqlDb := c.PostForm("database")
	mysqlUsername := c.PostForm("username")
	mysqlPassword := c.PostForm("password")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqlUsername, mysqlPassword, mysqlServer, mysqlPort, mysqlDb)
	mysql := database.NewMysql()
	mysql.Dsn = dsn
	err := mysql.Ping()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "数据库连接失败：" + err.Error(),
		})
		return
	}
	isExist, _ := tools.IsFileExist(config.Dir)
	if !isExist {
		os.Mkdir(config.Dir, os.ModePerm)
	}
	fileConfig := config.MysqlConf
	file, _ := os.OpenFile(fileConfig, os.O_RDWR|os.O_CREATE, os.ModePerm)

	format := `{
	"Server":"%s",
	"Port":"%s",
	"Database":"%s",
	"Username":"%s",
	"Password":"%s"
}
`
	data := fmt.Sprintf(format, mysqlServer, mysqlPort, mysqlDb, mysqlUsername, mysqlPassword)
	file.WriteString(data)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "操作成功",
	})
}
