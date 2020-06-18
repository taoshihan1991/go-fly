package tools

import "html/template"

type MailServer struct {
	Server, Email, Password string
}
type ViewHtml struct {
	Header template.HTML
	Nav    template.HTML
}
type IndexData struct {
	ViewHtml
	Folders           map[string]int
	Mails             interface{}
	MailPagelist      []*MailItem
	CurrentPage       int
	Fid               string
	NextPage, PrePage string
	NumPages          template.HTML
}
type ViewData struct {
	Folders  map[string]int
	HtmlBody template.HTML
	MailItem
}
type MailItem struct {
	Subject string
	Fid     string
	Id      uint32
	From    string
	To      string
	Body    string
	Date    string
}
type MailPageList struct {
	MailItems []*MailItem
}
type JsonResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type JsonListResult struct {
	JsonResult
	Result interface{} `json:"result"`
}
type SmtpBody struct {
	Smtp     string
	From     string
	To       []string
	Password string
	Subject  string
	Body     string
}
