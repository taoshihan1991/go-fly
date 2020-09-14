package controller

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/config"
	"log"
	"sort"
)

func PostCheckWeixinSign(c *gin.Context){
	token:=config.WeixinToken
	signature:=c.PostForm("signature")
	timestamp:=c.PostForm("timestamp")
	nonce:=c.PostForm("nonce")
	echostr:=c.PostForm("echostr")
	//将token、timestamp、nonce三个参数进行字典序排序
	var tempArray  = []string{token, timestamp, nonce}
	sort.Strings(tempArray)
	//将三个参数字符串拼接成一个字符串进行sha1加密
	var sha1String string = ""
	for _, v := range tempArray {
		sha1String += v
	}
	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))
	//获得加密后的字符串可与signature对比
	if sha1String == signature {
		c.Writer.Write([]byte(echostr))
	} else {
		log.Println("微信API验证失败")
	}
}
