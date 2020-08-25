package config

import (
	"encoding/json"
	"fmt"
	"github.com/taoshihan1991/imaptool/tools"
	"io/ioutil"
	"os"
)
var(
	PageSize uint=10
	VisitorPageSize uint=8
)
const Dir = "config/"
const AccountConf = Dir + "account.json"
const MysqlConf = Dir + "mysql.json"
const MailConf = Dir + "mail.json"
const LangConf=Dir+"language.json"
const MainConf = Dir + "config.json"
type Mysql struct{
	Server string
	Port string
	Database string
	Username string
	Password string
}
type MailServer struct {
	Server, Email, Password string
}
type Config struct {
	Upload string
}
func CreateConfig()*Config{
	var configObj Config
	c:=&Config{
		Upload: "static/upload/",
	}
	isExist, _ := tools.IsFileExist(MainConf)
	if !isExist {
		return c
	}
	info, err := ioutil.ReadFile(MainConf)
	if err != nil {
		return c
	}
	err = json.Unmarshal(info, &configObj)
	return &configObj
}
func CreateMailServer() *MailServer {
	var imap MailServer
	isExist, _ := tools.IsFileExist(MailConf)
	if !isExist {
		return &imap
	}
	info, err := ioutil.ReadFile(MailConf)
	if err != nil {
		return &imap
	}

	err = json.Unmarshal(info, &imap)
	return &imap
}
func CreateMysql() *Mysql {
	var mysql Mysql
	isExist, _ := tools.IsFileExist(MysqlConf)
	if !isExist {
		return &mysql
	}
	info, err := ioutil.ReadFile(MysqlConf)
	if err != nil {
		return &mysql
	}

	err = json.Unmarshal(info, &mysql)
	return &mysql
}
func GetMysql() map[string]string {
	var mysql map[string]string
	isExist, _ := tools.IsFileExist(MysqlConf)
	if !isExist {
		return mysql
	}
	info, err := ioutil.ReadFile(MysqlConf)
	if err != nil {
		return mysql
	}

	err = json.Unmarshal(info, &mysql)
	return mysql
}
func GetAccount() map[string]string {
	var account map[string]string
	isExist, _ := tools.IsFileExist(AccountConf)
	if !isExist {
		return account
	}
	info, err := ioutil.ReadFile(AccountConf)
	if err != nil {
		return account
	}

	err = json.Unmarshal(info, &account)
	return account
}
func GetUserInfo(uid string) map[string]string {
	var userInfo map[string]string
	userFile := Dir + "sess_" + uid + ".json"
	isExist, _ := tools.IsFileExist(userFile)
	if !isExist {
		return userInfo
	}
	info, err := ioutil.ReadFile(userFile)
	if err != nil {
		return userInfo
	}

	err = json.Unmarshal(info, &userInfo)
	return userInfo
}
func SetUserInfo(uid string, info map[string]string) {
	userFile := Dir + "sess_" + uid + ".json"
	isExist, _ := tools.IsFileExist(Dir)
	if !isExist {
		os.Mkdir(Dir, os.ModePerm)
	}
	file, _ := os.OpenFile(userFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	str := "{\r\n"
	for k, v := range info {
		str += fmt.Sprintf(`"%s":"%s",`, k, v)
	}
	str += fmt.Sprintf(`"session_id":"%s"%s}`, uid, "\r\n")
	file.WriteString(str)
}
