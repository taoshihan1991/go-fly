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
func SendNoticeEmail(username, msg string) {
	smtp := models.FindConfig("NoticeEmailSmtp")
	email := models.FindConfig("NoticeEmailAddress")
	password := models.FindConfig("NoticeEmailPassword")
	if smtp == "" || email == "" || password == "" {
		return
	}
	err:=tools.SendSmtp(smtp, email, password, []string{email}, "[通知]"+username, msg)
	if err!=nil{
		log.Println(err)
	}
}
