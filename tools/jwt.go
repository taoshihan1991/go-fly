package tools

import (
	"github.com/golang-jwt/jwt"
)

const SECRET = "golivechat"

func MakeToken(obj map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(obj))
	tokenString, err := token.SignedString([]byte(SECRET))
	return tokenString, err
}
func ParseToken(tokenStr string) map[string]interface{} {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return nil
	}
	finToken := token.Claims.(jwt.MapClaims)
	return finToken
}
