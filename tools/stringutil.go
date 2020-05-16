// stringutil 包含有用于处理字符串的工具函数。
package tools

import (
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

// Reverse 将其实参字符串以符文为单位左右反转。
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//转换编码
func Encoding(html string,ct string) string {
	e,name:=DetermineEncoding(html)
	if name!="utf-8"{
		html=ConvertToStr(html,"gbk","utf-8")
		e=unicode.UTF8
	}
	r := strings.NewReader(html)

	utf8Reader := transform.NewReader(r, e.NewDecoder())
	//将其他编码的reader转换为常用的utf8reader
	all, _ := ioutil.ReadAll(utf8Reader)
	return string(all)
}
func DetermineEncoding(html string) (encoding.Encoding,string) {
	e, name, _ := charset.DetermineEncoding([]byte(html), "")
	return e,name
}
