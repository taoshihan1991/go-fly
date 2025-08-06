package tools

import "testing"

func TestGetCity(t *testing.T) {
	GetCity("../config/GeoLite2-City.mmdb", "")
}
