package tools

import (
	"log"
	"testing"
)

func TestGetOutboundIP(t *testing.T) {
	ip, err := GetOutboundIP()
	log.Println(ip, err)
}
