package tmpl

import "net/http"

type SettingHtml struct {
	*CommonHtml
	Username, Password string
}

func NewSettingHtml(w http.ResponseWriter) *SettingHtml {
	obj := new(SettingHtml)
	parent := NewRender(w)
	obj.CommonHtml = parent
	return obj
}
