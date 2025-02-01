package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// MyClain 自定义claim
type MyClaim struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(userId int64) (string, error) {
	// 创建一个自定义的claim
	claim := MyClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*24, // 过期时间
			Issuer:    "ai-hiresphere",              // 签发人
		},
	}
	// 使用HS256算法进行加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 硬编码进行加密 应该存储在配置文件中
	return token.SignedString([]byte("huiphfawoiegfuoizHsciuLHofuwehaif;h"))
}
