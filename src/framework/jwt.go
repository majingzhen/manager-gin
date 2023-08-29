package framework

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("manager_gin_token")

type MyClaims struct {
	UserId   string
	UserName string
	jwt.StandardClaims
}

// GenerateToken 生成Token
func GenerateToken(userId, userName string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userId,
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "manager_gin",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(MySecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {

	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		log.Printf("ParseToken, parseToken is error: %v", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
