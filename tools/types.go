package tools

import "html/template"
type MailServer struct{
	Server,Email,Password string
}
type IndexData struct {
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
type JsonResult  struct{
	Code int `json:"code"`
	Msg  string `json:"msg"`
}
type JsonFolders struct{
	JsonResult
	Result interface{} `json:"result"`
}