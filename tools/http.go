package tools

import (
	"fmt"
	"github.com/taoshihan1991/imaptool/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func Get(url string) string {
	res, err := http.Get(url)
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
func SendServerJiang(content string) string {
	noticeServerJiang, err := strconv.ParseBool(models.FindConfig("NoticeServerJiang"))
	serverJiangAPI := models.FindConfig("ServerJiangAPI")
	if err != nil || !noticeServerJiang || serverJiangAPI == "" {
		log.Println("do not notice serverjiang:", serverJiangAPI, noticeServerJiang)
		return ""
	}
	sendStr := fmt.Sprintf("%s,访客来了", content)
	desp := "[登录](https://gofly.sopans.com/main)"
	url := serverJiangAPI + "?text=" + sendStr + "&desp=" + desp
	//log.Println(url)
	res := Get(url)
	return res
}
