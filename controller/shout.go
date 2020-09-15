package controller

import (
	"fmt"
	"github.com/taoshihan1991/imaptool/config"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
)

func SendServerJiang(content string)string{
	conf:=config.CreateConfig()
	if config.ServerJiang=="" || !conf.NoticeServerJiang{
		log.Println("do not notice serverjiang:",config.ServerJiang,conf.NoticeServerJiang)
		return ""
	}
	sendStr:=fmt.Sprintf("%s,访客来了",content)
	desp:="[登录](https://gofly.sopans.com/main)";
	res:=tools.Get(config.ServerJiang+"?text="+sendStr+"&desp="+desp)
	return res
}
