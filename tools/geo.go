package tools

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

func GetCity(path, ipAddress string) (string, string) {
	db, err := geoip2.Open(path)
	if err != nil {
		return "", ""
	}
	defer db.Close()
	record, err := db.City(net.ParseIP(ipAddress))
	fmt.Println(record.City.Names["en"])
	return record.Country.Names["en"], record.City.Names["en"]
}
