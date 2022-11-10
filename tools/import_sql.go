package tools

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type ImportSqlTool struct {
	SqlPath                                    string
	Username, Password, Server, Port, Database string
}

func (this *ImportSqlTool) ImportSql() error {
	_, err := os.Stat(this.SqlPath)
	if os.IsNotExist(err) {
		log.Println("数据库SQL文件不存在:", err)
		return err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", this.Username, this.Password, this.Server, this.Port, this.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println("数据库连接失败:", err)
		//panic("数据库连接失败!")
		return err
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(59 * time.Second)

	sqls, _ := ioutil.ReadFile(this.SqlPath)
	sqlArr := strings.Split(string(sqls), ";")
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err := db.Exec(sql).Error
		if err != nil {
			log.Println("数据库导入失败:" + err.Error())
			return err
		} else {
			log.Println(sql, "\t success!")
		}
	}
	return nil
}

