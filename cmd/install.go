package cmd

import (
	"github.com/spf13/cobra"
	"github.com/taoshihan1991/imaptool/common"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "安装导入数据",
	Run: func(cmd *cobra.Command, args []string) {
		install()
	},
}

func install() {
	if ok, _ := tools.IsFileNotExist("./install.lock"); !ok {
		log.Println("请先删除./install.lock")
		os.Exit(1)
	}
	sqlFile := common.Dir + "go-fly.sql"
	isExit, _ := tools.IsFileExist(common.MysqlConf)
	dataExit, _ := tools.IsFileExist(sqlFile)
	if !isExit || !dataExit {
		log.Println("config/mysql.json 数据库配置文件或者数据库文件go-fly.sql不存在")
		os.Exit(1)
	}
	sqls, _ := ioutil.ReadFile(sqlFile)
	sqlArr := strings.Split(string(sqls), "|")
	for _, sql := range sqlArr {
		if sql == "" {
			continue
		}
		err := models.Execute(sql)
		if err == nil {
			log.Println(sql, "\t success!")
		} else {
			log.Println(sql, err, "\t failed!")
			os.Exit(1)
		}
	}
	installFile, _ := os.OpenFile("./install.lock", os.O_RDWR|os.O_CREATE, os.ModePerm)
	installFile.WriteString("gofly live chat")
}
