package tools

import (
	"github.com/ipipdotnet/ipdb-go"
)

func ParseIp(myip string) *ipdb.CityInfo {
	db, err := ipdb.NewCity("./config/city.free.ipdb")
	if err != nil {
		return nil
	}
	db.Reload("./config/city.free.ipdb")
	c, err := db.FindInfo(myip, "CN")
	if err != nil {
		return nil
	}
	return c
}
