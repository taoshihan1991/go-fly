// stringutil 包含有用于处理字符串的工具函数。
package tools

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"net/http"
)

//获取URL的GET参数
func GetUrlArg(r *http.Request, name string) string {
	var arg string
	values := r.URL.Query()
	arg = values.Get(name)
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

// Reverse2 将其实参字符串以符文为单位左右反转。
func Reverse2(s string) string {
	r := []rune(s)
	left := 0
	right := len(r) - 1
	for left < right {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
	return string(r)
}

//获取文件内容，可以打包到二进制
func FileGetContent(file string) string {
	str := ""
	box := packr.New("tmpl", "../static")
	content, err := box.FindString(file)
	if err != nil {
		return str
	}
	return content
}

func ShowStringByte(str string) {
	s := []byte(str)
	for i, c := range s {
		fmt.Println(i, c)
	}
}
func NilChannel() {
	var ch chan int
	ch <- 1
}
func Int2Str(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
