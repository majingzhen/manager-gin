package utils

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func GenUID() string {
	uid := uuid.NewV4()
	uidStr := strings.ReplaceAll(uid.String(), "-", "")
	return uidStr
}
