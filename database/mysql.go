package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/taoshihan1991/imaptool/config"
)
type Mysql struct{
	SqlDB *sql.DB
	Dsn  string
}
func NewMysql()*Mysql{
	mysqlInfo:=config.GetMysql()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqlInfo["Username"], mysqlInfo["Password"], mysqlInfo["Server"], mysqlInfo["Port"], mysqlInfo["Database"])
	return &Mysql{
		Dsn:dsn,
	}
}

func (db *Mysql)Ping()error{
	sqlDb, _ := sql.Open("mysql", db.Dsn)
	db.SqlDB=sqlDb
	return db.SqlDB.Ping()
}