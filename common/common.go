package common

var (
	PageSize          uint    = 10
	VisitorPageSize   uint    = 8
	Version           string  = "0.3.9"
	VisitorExpire     float64 = 600
	Upload            string  = "static/upload/"
	Dir               string  = "config/"
	MysqlConf         string  = Dir + "mysql.json"
	IsCompireTemplate bool    = false //是否编译静态模板到二进制
)
