package tools

import "testing"

func TestImportSql(t *testing.T) {
	tool:=&ImportSqlTool{
		SqlPath:  "../import.sql",
		Username: "go-fly",
		Password: "go-fly",
		Server:   "127.0.0.1",
		Port:     "3306",
		Database: "go-fly",
	}
	tool.ImportSql()
}