package tools

import (
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"strings"
)

func Send(server string,from string,password string,to []string,subject string,body string)error{
	auth := sasl.NewPlainClient("", from, password)
	msg := strings.NewReader(
		"From: "+from+"\r\n"+
			"To: "+strings.Join(to,",")+"\r\n"+
		"Subject: "+subject+"\r\n" +
		"\r\n" +
		body+"\r\n")
	err := smtp.SendMail(server, auth, from, to, msg)
	if err != nil {
		return err
	}
	return nil
}