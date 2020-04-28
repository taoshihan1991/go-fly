package tools

import (
	"fmt"
	"github.com/emersion/go-imap/client"
	"strconv"
	"strings"
)

func CheckEmailPassword(server string, email string, password string) bool {
	if !strings.Contains(server, ":") {
		return false
	}
	var c *client.Client
	var err error
	serverSlice := strings.Split(server, ":")
	uri := serverSlice[0]
	port, _ := strconv.Atoi(serverSlice[1])
	if port != 993 && port != 143 {
		return false
	}
	if port == 993 {
		c, err = client.DialTLS(fmt.Sprintf("%s:%d", uri, port), nil)
	} else {
		c, err = client.Dial(fmt.Sprintf("%s:%d", uri, port))
	}
	if err != nil {
		return false
	}

	// 不要忘了退出
	defer c.Logout()

	// 登陆
	if err := c.Login(email, password); err != nil {
		return false
	}
	return true
}
