package tools

import (
	"github.com/satori/go.uuid"
)

func Uuid() string {
	u2 := uuid.NewV4()
	return u2.String()
}
