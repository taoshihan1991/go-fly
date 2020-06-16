// stringutil 包含有用于处理字符串的工具函数。
package tools

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gobuffalo/packr/v2"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"strings"
)
//获取URL的GET参数
func GetUrlArg(r *http.Request,name string)string{
	var arg string
	values := r.URL.Query()
	arg=values.Get(name)
	return arg
}
// Reverse 将其实参字符串以符文为单位左右反转。
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//转换编码
func Encoding(html string, ct string) string {
	e, name := DetermineEncoding(html)
	if name != "utf-8" {
		html = ConvertToStr(html, "gbk", "utf-8")
		e = unicode.UTF8
	}
	r := strings.NewReader(html)

	utf8Reader := transform.NewReader(r, e.NewDecoder())
	//将其他编码的reader转换为常用的utf8reader
	all, _ := ioutil.ReadAll(utf8Reader)
	return string(all)
}
func DetermineEncoding(html string) (encoding.Encoding, string) {
	e, name, _ := charset.DetermineEncoding([]byte(html), "")
	return e, name
}
//获取文件内容，可以打包到二进制
func FileGetContent(file string) string {
	str := ""
	box := packr.New("tmpl","../static")
	content, err := box.FindString(file)
	if err != nil {
		return str
	}
	return content
}
//md5加密
func Md5(src string)string{
	m:=md5.New()
	m.Write([]byte(src))
	res:=hex.EncodeToString(m.Sum(nil))
	return res
}
