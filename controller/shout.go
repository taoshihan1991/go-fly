package controller

import (
	"fmt"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
	"strconv"
)

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
	res := tools.Get(url)
	return res
}
