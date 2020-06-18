package tmpl

import "net/http"

type FolderHtml struct {
	*CommonHtml
	CurrentPage int
	Fid         string
}

func NewFolderHtml(w http.ResponseWriter) *FolderHtml {
	obj := new(FolderHtml)
	parent := NewRender(w)
	obj.CommonHtml = parent
	return obj
}
