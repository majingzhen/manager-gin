package utils

import (
	"github.com/gogf/gf/crypto/gmd5"
	"manager-gin/src/common"
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

func IsHttp(link string) bool {
	prefixes := []string{common.HTTP, common.HTTPS}
	for _, prefix := range prefixes {
		if strings.HasPrefix(link, prefix) {
			return true
		}
	}
	return false
}

func ReplaceEach(str string, searchList []string, replacementList []string) string {
	if str == "" || searchList == nil || replacementList == nil {
		return ""
	}

	if len(searchList) != len(replacementList) {
		panic("Search and Replace array lengths don't match")
	}

	for i := 0; i < len(searchList); i++ {
		str = strings.Replace(str, searchList[i], replacementList[i], -1)
	}

	return str
}
