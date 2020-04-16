package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

//全局变量
//imap服务地址,邮箱,密码
var (
	server, email, password string
)

func main() {
	//获取参数中的数据
	flag.StringVar(&server, "server", "", "imap服务地址(包含端口)")
	flag.StringVar(&email, "email", "", "邮箱名")
	flag.StringVar(&password, "password", "", "密码")
	flag.Parse()
	if flag.NFlag() < 3 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if server == "" || email == "" || password == "" {
		log.Fatal("服务器地址,用户名,密码,参数必填")
	}
	log.Println("正在连接服务器...")
	//支持加密和非加密端口
	serverSlice := strings.Split(server, ":")
	uri := serverSlice[0]
	port, _ := strconv.Atoi(serverSlice[1])
	if port != 993 && port != 143 {
		log.Fatal("服务器地址端口号错误",port)
	}
	// 连接到服务器
	var c *client.Client
	var err error
	if port == 993 {
		c, err = client.DialTLS(fmt.Sprintf("%s:%d", uri, port), nil)
	} else {
		c, err = client.Dial(fmt.Sprintf("%s:%d", uri, port))
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Println("连接成功")

	// 不要忘了退出
	defer c.Logout()

	// 登陆
	if err := c.Login(email, password); err != nil {
		log.Fatal(err)
	}
	log.Println("成功登陆")

	// 列邮箱
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("邮箱:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("INBOX的邮件个数:", mbox.Messages)

	// 获取最新的 4 封信
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// 我们在这使用无符号整型, 这是再获取from的id
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	log.Println("最新的 4 封信:")
	for msg := range messages {
		log.Println("* " + msg.Envelope.Subject)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("结束!")
}
