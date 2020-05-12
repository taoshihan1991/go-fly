package tools

import "html/template"

type IndexData struct {
	Folders map[string]int
	Mails   interface{}
	MailPagelist []*MailItem
	CurrentPage int
	Fid string
	NextPage,PrePage string
	NumPages template.HTML

}
type ViewData struct {
	Folders map[string]int
	Fid string
}
type MailItem struct{
	Subject string
	Id uint32
	Fid string
}
type MailPageList struct{
	MailItems []*MailItem
}
