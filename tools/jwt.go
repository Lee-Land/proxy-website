package tools

import (
	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtSecret = []byte("#liu_proxy#")
)

type MyClamis struct {
	UserName string
	jwt.RegisteredClaims
}

/*
	生成jwt字符串
*/
func GenerateToken(username string) (string, error) {
	clamis := &MyClamis{
		username,
		jwt.RegisteredClaims{
			//设置过期时间
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		},
	}
	// 创建Token结构体
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis)

	// 调用加密方法，发挥Token字符串
	signingString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return signingString, nil
}

/*
	解析jwt字符串
*/
func ParseToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		token, jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
}
