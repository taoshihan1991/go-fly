package tools

import (
	"encoding/base64"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"strings"
)

func SendSmtp(server string, from string, password string, to []string, subject string, body string) error {
	auth := sasl.NewPlainClient("", from, password)
	subjectBase := base64.StdEncoding.EncodeToString([]byte(subject))
	msg := strings.NewReader(
		"From: " + from + "\r\n" +
			"To: " + strings.Join(to, ",") + "\r\n" +
			"Subject: =?UTF-8?B?" + subjectBase + "?=\r\n" +
			"Content-Type: text/html; charset=UTF-8" +
			"\r\n\r\n" +
			body + "\r\n")
	err := smtp.SendMail(server, auth, from, to, msg)
	if err != nil {
		return err
	}
	return nil
}
