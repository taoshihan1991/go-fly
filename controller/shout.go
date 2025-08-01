package controller

import (
	"encoding/json"
	"fmt"
	"goflylivechat/models"
	"goflylivechat/tools"
	"goflylivechat/ws"
	"log"
	"strconv"
	"time"
)

func SendServerJiang(title string, content string, domain string) string {
	noticeServerJiang, err := strconv.ParseBool(models.FindConfig("NoticeServerJiang"))
	serverJiangAPI := models.FindConfig("ServerJiangAPI")
	if err != nil || !noticeServerJiang || serverJiangAPI == "" {
		log.Println("do not notice serverjiang:", serverJiangAPI, noticeServerJiang)
		return ""
	}
	sendStr := fmt.Sprintf("%s%s", title, content)
	desp := title + ":" + content + "[登录](http://" + domain + "/main)"
	url := serverJiangAPI + "?text=" + sendStr + "&desp=" + desp
	//log.Println(url)
	res := tools.Get(url)
	return res
}
func SendVisitorLoginNotice(kefuName, visitorName, avator, content, visitorId string) {
	if !tools.LimitFreqSingle("sendnotice:"+visitorId, 1, 120) {
		log.Println("SendVisitorLoginNotice limit")
		return
	}
	userInfo := make(map[string]string)
	userInfo["username"] = visitorName
	userInfo["avator"] = avator
	userInfo["content"] = content
	msg := ws.TypeMessage{
		Type: "notice",
		Data: userInfo,
	}
	str, _ := json.Marshal(msg)
	ws.OneKefuMessage(kefuName, str)
}
func SendNoticeEmail(username, msg string) {
	smtp := models.FindConfig("NoticeEmailSmtp")
	email := models.FindConfig("NoticeEmailAddress")
	password := models.FindConfig("NoticeEmailPassword")
	if smtp == "" || email == "" || password == "" {
		return
	}
	err := tools.SendSmtp(smtp, email, password, []string{email}, "[通知]"+username, msg)
	if err != nil {
		log.Println(err)
	}
}
func SendAppGetuiPush(kefu string, title, content string) {
	token := models.FindConfig("GetuiToken")
	if token == "" {
		token = getGetuiToken()
		if token == "" {
			return
		}
	}
	format := `
{
    "request_id":"%s",
    "settings":{
        "ttl":3600000
    },
    "audience":{
        "cid":[
            "%s"
        ]
    },
    "push_message":{
        "notification":{
            "title":"%s",
            "body":"%s",
            "click_type":"url",
            "url":"https//:xxx"
        }
    }
}
`
	clients := models.FindClients(kefu)
	if len(clients) == 0 {
		return
	}
	//clientIds := make([]string, 0)
	for _, client := range clients {
		//clientIds = append(clientIds, client.Client_id)
		req := fmt.Sprintf(format, tools.Md5(tools.Uuid()), client.Client_id, title, content)
		num := sendPushApi(token, req)
		if num == 10001 {
			token = getGetuiToken()
			sendPushApi(token, req)
		}
	}

}
func sendPushApi(token string, req string) int {
	appid := models.FindConfig("GetuiAppID")
	if appid == "" {
		return 0
	}
	url := "https://restapi.getui.com/v2/" + appid + "/push/single/cid"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json;charset=utf-8"
	headers["token"] = token
	res, err := tools.PostHeader(url, []byte(req), headers)
	tools.Logger().Infoln(url, req, err, res)

	if err == nil && res != "" {
		var pushRes GetuiResponse
		json.Unmarshal([]byte(res), &pushRes)
		if pushRes.Code == 10001 {
			return 10001
		}
	}
	return 200
}
func getGetuiToken() string {
	appid := models.FindConfig("GetuiAppID")
	appkey := models.FindConfig("GetuiAppKey")
	//appsecret := models.FindConfig("GetuiAppSecret")
	appmastersecret := models.FindConfig("GetuiMasterSecret")
	if appid == "" {
		return ""
	}
	type req struct {
		Sign      string `json:"sign"`
		Timestamp string `json:"timestamp"`
		Appkey    string `json:"appkey"`
	}
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	reqJson := req{
		Sign:      tools.Sha256(appkey + timestamp + appmastersecret),
		Timestamp: timestamp,
		Appkey:    appkey,
	}
	reqStr, _ := json.Marshal(reqJson)
	url := "https://restapi.getui.com/v2/" + appid + "/auth"
	res, err := tools.Post(url, "application/json;charset=utf-8", reqStr)
	log.Println(url, string(reqStr), err, res)
	if err == nil && res != "" {
		var pushRes GetuiResponse
		json.Unmarshal([]byte(res), &pushRes)
		if pushRes.Code == 0 {
			token := pushRes.Data["token"].(string)
			//models.UpdateConfig("GetuiToken", token)
			return token
		}
	}
	return ""
}
