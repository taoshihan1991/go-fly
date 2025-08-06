package tools

import (
	"github.com/oschwald/geoip2-golang/v2"
	"net/netip"
)

func GetCity(path, ipAddress string) (string, string) {
	db, err := geoip2.Open(path)
	if err != nil {
		return "", ""
	}
	defer db.Close()
	ip, err := netip.ParseAddr(ipAddress)
	if err != nil {
		return "", ""
	}
	record, err := db.City(ip)
	if err != nil {
		return "", ""
	}
	return record.Country.Names.English, record.City.Names.English
}
