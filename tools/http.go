package tools

import (
	"io/ioutil"
	"net/http"
)

func Get(url string)string{
	res, err :=http.Get(url)
	if err != nil {
		return ""
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return ""
	}
	return string(robots)
}
