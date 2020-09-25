package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/taoshihan1991/imaptool/config"
)

type Mysql struct {
	SqlDB *sql.DB
	Dsn   string
}

func NewMysql() *Mysql {
	mysql := config.CreateMysql()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysql.Username, mysql.Password, mysql.Server, mysql.Port, mysql.Database)
	return &Mysql{
		Dsn: dsn,
	}
}

func (db *Mysql) Ping() error {
	sqlDb, _ := sql.Open("mysql", db.Dsn)
	db.SqlDB = sqlDb
	return db.SqlDB.Ping()
}
