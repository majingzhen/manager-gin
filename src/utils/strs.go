package utils

import (
	"github.com/gogf/gf/crypto/gmd5"
	"strings"
)

func ToTitle(underscoreName string) string {
	words := strings.Split(strings.ToLower(underscoreName), "_")
	for i := 0; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}
func ToCamelCase(underscoreName string) string {
	words := strings.Split(strings.ToLower(underscoreName), "_")
	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

// EncryptionPassword 加密密码
func EncryptionPassword(password, salt string) string {
	// 将盐与密码结合后进行哈希计算
	hashedPassword := gmd5.MustEncryptString(password + salt)
	return hashedPassword
}
