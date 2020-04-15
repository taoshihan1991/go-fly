package main

import (
	"log"

	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-imap"
)

func main() {
	log.Println("正在连接服务器...")

	// 连接到服务器
	c, err := client.DialTLS("imap.sina.net:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("连接成功")

	// 不要忘了退出
	defer c.Logout()

	// 登陆
	if err := c.Login("shihan2@appdev.sinanet.com", "tsh_xxx"); err != nil {
		log.Fatal(err)
	}
	log.Println("成功登陆")

	// 列邮箱
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func () {
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
	log.Println("INBOX的Flags标记:", mbox.Flags)

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