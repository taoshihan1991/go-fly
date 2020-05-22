package tools

import (
	"net/http"
	"strings"
)

func SetCookie(name string, value string, w *http.ResponseWriter) {
	cookie := http.Cookie{
		Name:  name,
		Value: value,
	}
	http.SetCookie(*w, &cookie)
}
func GetCookie(r *http.Request, name string) string {
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}
func GetMailServerFromCookie(r *http.Request) *MailServer {
	auth := GetCookie(r, "auth")
	if !strings.Contains(auth, "|") {
		return nil
	}
	authStrings := strings.Split(auth, "|")
	mailServer := &MailServer{
		Server:   authStrings[0],
		Email:    authStrings[1],
		Password: authStrings[2],
	}
	return mailServer
}
