package tools

import (
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
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
	c=connect(server,email,password)
	if c==nil{
		return false
	}
	return true
}
//获取连接
func connect(server string, email string, password string)*client.Client{
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
//获取邮件夹
func GetFolders(server string, email string, password string,folder string)map[string]int{
	var c *client.Client
	//defer c.Logout()
	c=connect(server,email,password)
	if c==nil{
		return nil
	}
	// 列邮箱
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()
	// 存储邮件夹
	var folders =make(map[string]int)
	for m := range mailboxes {
		folders[m.Name]=0
	}
	for m,_ := range folders {
		if m==folder {
			mbox, _ := c.Select(m, true)
			if mbox != nil {
				folders[m] = int(mbox.Messages)
			}
			break
		}
	}
	log.Println(folders)
	return folders
}

//获取邮件夹邮件
func GetFolderMail(server string, email string, password string,folder string,currentPage int,pagesize int)[]string{
	var c *client.Client
	//defer c.Logout()
	c=connect(server,email,password)
	if c==nil{
		return nil
	}

	mbox, _ := c.Select(folder, true)
	to := mbox.Messages-uint32((currentPage-1)*pagesize)
	from := to-uint32(pagesize)
	if to <=uint32(pagesize){
		from=1
	}

	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, pagesize)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()
	var res []string

	dec :=new(mime.WordDecoder)
	dec.CharsetReader= func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "gb2312":
			content, err := ioutil.ReadAll(input)
			if err != nil {
				return nil, err
			}
			//ret:=bytes.NewReader(content)
			//ret:=transform.NewReader(bytes.NewReader(content), simplifiedchinese.HZGB2312.NewEncoder())

			utf8str:=ConvertToStr(string(content),"gbk","utf-8")
			t:=bytes.NewReader([]byte(utf8str))
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

			utf8str:=ConvertToStr(string(content),"gbk","utf-8")
			t:=bytes.NewReader([]byte(utf8str))
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

			utf8str:=ConvertToStr(string(content),"gbk","utf-8")
			t:=bytes.NewReader([]byte(utf8str))
			//ret:=utf8.DecodeRune(t)
			//log.Println(ret)
			return t, nil
		default:
			return nil,fmt.Errorf("unhandle charset:%s",charset)

		}
	}
	for msg:=range messages{

		ret,err:=dec.Decode(msg.Envelope.Subject)
		if err!=nil{
			ret,_=dec.DecodeHeader(msg.Envelope.Subject)
		}
		res=append(res,ret)
	}

	log.Println(res)
	return res
}
// 任意编码转特定编码
func ConvertToStr(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}