package config

type Language struct {
	WebCopyRight                                                             string
	MainIntro                                                                string
	Send                                                                     string
	Notice, Maintech, NowAsk, LaterAsk                                       string
	IndexSubIntro, IndexVisitors, IndexAgent, IndexDocument, IndexOnlineChat string
}

func CreateLanguage(lang string) *Language {
	var language *Language

	if lang == "en" {
		language = &Language{
			WebCopyRight:    "TaoShihan",
			MainIntro:       "Simple and Powerful Go language online customer chat system",
			IndexSubIntro:   "GO-FLY, a Vue 2.0-based online customer service instant messaging system for PHP engineers and Golang engineers",
			IndexDocument:   "API Documents",
			IndexVisitors:   "Visitors Here",
			IndexAgent:      "Agents Here",
			IndexOnlineChat: "Let’s chat. - We're online",
			Send:            "Send",
			Notice:          "Hello and welcome to go-fly - how can we help?",
			Maintech:        "Main technical architecture",
			NowAsk:          "Start Chat",
			LaterAsk:        "Chat Later",
		}
	}
	if lang == "cn" {
		language = &Language{
			WebCopyRight:    "陶士涵的菜地版权所有",
			MainIntro:       "极简强大的Go语言在线客服系统",
			IndexSubIntro:   "GO-FLY，一套为PHP工程师、Golang工程师准备的基于 Vue 2.0的在线客服即时通讯系统",
			IndexVisitors:   "访客入口",
			IndexAgent:      "客服入口",
			IndexDocument:   "接口文档",
			IndexOnlineChat: "真人在线-点击与我交流",
			Send:            "发送",
			Notice:          "欢迎您访问go-fly！有什么我能帮助您的？",
			Maintech:        "主要技术架构",
			NowAsk:          "现在咨询",
			LaterAsk:        "稍后再说",
		}
	}
	return language
}
