package tools

import (
	"errors"
	"github.com/ipipdotnet/ipdb-go"
	"net"
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
func GetServerIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}
func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
func GetExternalIp() string {
	return Get("http://myexternalip.com/raw")
}
