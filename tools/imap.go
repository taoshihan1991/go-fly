package tools

import (
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"strconv"
	"strings"
)

//验证邮箱密码
func CheckEmailPassword(server string, email string, password string) bool {
	if !strings.Contains(server, ":") {
		return false
	}
	var c *client.Client
	serverSlice := strings.Split(server, ":")
	port, _ := strconv.Atoi(serverSlice[1])
	if port != 993 && port != 143 {
		return false
	}

	// 不要忘了退出
	//defer c.Logout()

	// 登陆
	c = connect(server, email, password)
	if c == nil {
		return false
	}
	return true
}

//获取连接
func connect(server string, email string, password string) *client.Client {
	var c *client.Client
	var err error
	serverSlice := strings.Split(server, ":")
	uri := serverSlice[0]
	port, _ := strconv.Atoi(serverSlice[1])
	if port != 993 && port != 143 {
		return nil
	}
	if port == 993 {
		c, err = client.DialTLS(fmt.Sprintf("%s:%d", uri, port), nil)
	} else {
		c, err = client.Dial(fmt.Sprintf("%s:%d", uri, port))
	}
	if err != nil {
		return nil
	}

	// 登陆
	if err := c.Login(email, password); err != nil {
		return nil
	}
	return c
}
//获取邮件总数
func GetMailNum(server string, email string, password string) map[string]int {
	var c *client.Client
	//defer c.Logout()
	c = connect(server, email, password)
	if c == nil {
		return nil
	}
	// 列邮箱
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()
	//// 存储邮件夹
	var folders = make(map[string]int)
	for m := range mailboxes {
		folders[m.Name] = 0
	}
	for m, _ := range folders {
		mbox, _ := c.Select(m, true)
		if mbox != nil {
			folders[m] = int(mbox.Messages)
		}
	}
	return folders
}
//获取邮件夹
func GetFolders(server string, email string, password string, folder string) map[string]int {
	var c *client.Client
	//defer c.Logout()
	c = connect(server, email, password)
	if c == nil {
		return nil
	}
	// 列邮箱
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()
	// 存储邮件夹
	var folders = make(map[string]int)
	for m := range mailboxes {
		folders[m.Name] = 0
	}
	for m, _ := range folders {
		if m == folder {
			mbox, _ := c.Select(m, true)
			if mbox != nil {
				folders[m] = int(mbox.Messages)
			}
			break
		}
	}
	//log.Println(folders)
	return folders
}

//获取邮件夹邮件
func GetFolderMail(server string, email string, password string, folder string, currentPage int, pagesize int) []*MailItem {
	var c *client.Client
	//defer c.Logout()
	c = connect(server, email, password)
	if c == nil {
		return nil
	}

	mbox, _ := c.Select(folder, true)
	to := mbox.Messages - uint32((currentPage-1)*pagesize)
	from := to - uint32(pagesize)
	if to <= uint32(pagesize) {
		from = 1
	}

	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, pagesize)
	done := make(chan error, 1)
	fetchItem:=imap.FetchItem(imap.FetchEnvelope)
	items := make([]imap.FetchItem,0)
	items=append(items,fetchItem)
	go func() {
		done <- c.Fetch(seqset, items, messages)
	}()
	var mailPagelist = new(MailPageList)

	dec := GetDecoder()

	for msg := range messages {
		log.Println(msg.Envelope.Date)

		ret, err := dec.Decode(msg.Envelope.Subject)
		if err != nil {
			ret, _ = dec.DecodeHeader(msg.Envelope.Subject)
		}
		var mailitem = new(MailItem)

		mailitem.Subject = ret
		mailitem.Id = msg.SeqNum
		mailitem.Fid = folder
		mailitem.Date = msg.Envelope.Date.String()
		from := ""
		for _, s := range msg.Envelope.Sender {
			from += s.Address()
		}
		mailitem.From = from
		mailPagelist.MailItems = append(mailPagelist.MailItems, mailitem)
	}
	return mailPagelist.MailItems
}
func GetMessage(server string, email string, password string, folder string, id uint32) *MailItem {
	var c *client.Client
	//defer c.Logout()
	c = connect(server, email, password)
	if c == nil {
		//return nil
	}
	// Select INBOX
	mbox, err := c.Select(folder, false)
	if err != nil {
		log.Fatal(err)
	}

	// Get the last message
	if mbox.Messages == 0 {
		log.Fatal("No message in mailbox")
	}
	seqSet := new(imap.SeqSet)
	seqSet.AddNum(id)

	// Get the whole message body
	section:= &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 1)
	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			log.Fatal(err)
		}
	}()

	msg := <-messages
	if msg == nil {
		log.Fatal("Server didn't returned message")
	}

	r := msg.GetBody(section)

	if r == nil {
		log.Fatal("Server didn't returned message body")
	}
	var mailitem = new(MailItem)

	// Create a new mail reader
	mr, _ := mail.CreateReader(r)


	// Print some info about the message
	header := mr.Header
	date, _ := header.Date()

	mailitem.Date = date.String()

	var f string
	dec := GetDecoder()

	if from, err := header.AddressList("From"); err == nil {
		for _, address := range from {
			fromStr := address.String()
			temp, _ := dec.DecodeHeader(fromStr)
			f += " " + temp
		}
	}
	mailitem.From = f
	log.Println("From:", mailitem.From)

	var t string
	if to, err := header.AddressList("To"); err == nil {
		log.Println("To:", to)
		for _, address := range to {
			toStr := address.String()
			temp, _ := dec.DecodeHeader(toStr)
			t += " " + temp
		}
	}
	mailitem.To = t

	subject, _ := header.Subject()
	s, err := dec.Decode(subject)
	if err != nil {
		s, _ = dec.DecodeHeader(subject)
	}
	log.Println("Subject:", s)
	mailitem.Subject = s
	// Process each message's part
	var bodyMap = make(map[string]string)
	bodyMap["text/plain"] = ""
	bodyMap["text/html"] = ""

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			//log.Fatal(err)
		}
		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			// This is the message's text (can be plain-text or HTML)

			b, _ := ioutil.ReadAll(p.Body)
			ct := p.Header.Get("Content-Type")
			if strings.Contains(ct, "text/plain") {
				bodyMap["text/plain"] += Encoding(string(b), ct)
			} else {
				bodyMap["text/html"] += Encoding(string(b), ct)
			}
			//body,_:=dec.Decode(string(b))
		case *mail.AttachmentHeader:
			// This is an attachment
			filename, _ := h.Filename()
			log.Println("Got attachment: ", filename)
		}

	}
	if bodyMap["text/html"] != "" {
		mailitem.Body = bodyMap["text/html"]
	} else {
		mailitem.Body = bodyMap["text/plain"]
	}
	//log.Println(mailitem.Body)
	return mailitem
}
func GetDecoder() *mime.WordDecoder {
	dec := new(mime.WordDecoder)
	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		charset = strings.ToLower(charset)
		switch charset {
		case "gb2312":
			content, err := ioutil.ReadAll(input)
			if err != nil {
				return nil, err
			}
			//ret:=bytes.NewReader(content)
			//ret:=transform.NewReader(bytes.NewReader(content), simplifiedchinese.HZGB2312.NewEncoder())

			utf8str := ConvertToStr(string(content), "gbk", "utf-8")
			t := bytes.NewReader([]byte(utf8str))
			//ret:=utf8.DecodeRune(t)
			//log.Println(ret)
			return t, nil
		case "gbk":
			content, err := ioutil.ReadAll(input)
			if err != nil {
				return nil, err
			}
			//ret:=bytes.NewReader(content)
			//ret:=transform.NewReader(bytes.NewReader(content), simplifiedchinese.HZGB2312.NewEncoder())

			utf8str := ConvertToStr(string(content), "gbk", "utf-8")
			t := bytes.NewReader([]byte(utf8str))
			//ret:=utf8.DecodeRune(t)
			//log.Println(ret)
			return t, nil
		case "gb18030":
			content, err := ioutil.ReadAll(input)
			if err != nil {
				return nil, err
			}
			//ret:=bytes.NewReader(content)
			//ret:=transform.NewReader(bytes.NewReader(content), simplifiedchinese.HZGB2312.NewEncoder())

			utf8str := ConvertToStr(string(content), "gbk", "utf-8")
			t := bytes.NewReader([]byte(utf8str))
			//ret:=utf8.DecodeRune(t)
			//log.Println(ret)
			return t, nil
		default:
			return nil, fmt.Errorf("unhandle charset:%s", charset)

		}
	}
	return dec
}

// 任意编码转特定编码
func ConvertToStr(src string, srcCode string, tagCode string) string {
	result := mahonia.NewDecoder(srcCode).ConvertString(src)
	//srcCoder := mahonia.NewDecoder(srcCode)
	//srcResult := srcCoder.ConvertString(src)
	//tagCoder := mahonia.NewDecoder(tagCode)
	//_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	//result := string(cdata)
	return result
}
