package tools

import "testing"

func TestSendSmtp(t *testing.T) {
	body := "<a href=''>hello</a>"
	SendSmtp("smtp.sina.cn:25", "taoshihan1@sina.com", "382e8a5e11cfae8c", []string{"taoshihan1@sina.com"}, "123456", body)
}
