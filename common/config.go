package common

import (
	"encoding/json"
	"goflylivechat/tools"
	"io/ioutil"
)

type Mysql struct {
	Server   string
	Port     string
	Database string
	Username string
	Password string
}

func GetMysqlConf() *Mysql {
	var mysql = &Mysql{}
	isExist, _ := tools.IsFileExist(MysqlConf)
	if !isExist {
		return mysql
	}
	info, err := ioutil.ReadFile(MysqlConf)
	if err != nil {
		return mysql
	}
	err = json.Unmarshal(info, mysql)
	return mysql
}
