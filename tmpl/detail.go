package tmpl

import "net/http"

type DetailHtml struct {
	*CommonHtml
	Fid string
	Id  uint32
}

func NewDetailHtml(w http.ResponseWriter) *DetailHtml {
	obj := new(DetailHtml)
	parent := NewRender(w)
	obj.CommonHtml = parent
	return obj
}
