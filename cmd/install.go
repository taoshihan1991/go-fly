package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/taoshihan1991/imaptool/config"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"io/ioutil"
	"os"
	"strings"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "example:go-fly install",
	Run: func(cmd *cobra.Command, args []string) {
		install()
	},
}
func install(){
	sqlFile:=config.Dir+"go-fly.sql"
	isExit,_:=tools.IsFileExist(config.MysqlConf)
	dataExit,_:=tools.IsFileExist(sqlFile)
	if !isExit||!dataExit{
		fmt.Println("config/mysql.json 数据库配置文件或者数据库文件go-fly.sql不存在")
		os.Exit(1)
	}
	sqls,_:=ioutil.ReadFile(sqlFile)
	sqlArr:=strings.Split(string(sqls),";")
	for _,sql:=range sqlArr{
		if sql==""{
			continue
		}
		models.Execute(sql)
	}
}