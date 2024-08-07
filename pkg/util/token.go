package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenOptions struct {
	SecretKey string
	Duration  int64   // x 秒，间隔
	Fields    map[string]any
}

type Token struct {
	AccessToken  string
	AccessExpire int64  // 时间戳
}


func GenerateAuth2Token(opt TokenOptions) (Token, error) {

	now := time.Now().Add(-time.Minute).Unix()  // 损耗

	accessToken, err := genToken(now, opt.SecretKey, opt.Fields, opt.Duration)
	if err != nil {
		return Token{}, nil 
	}

	return Token{
		AccessToken:  accessToken,
		AccessExpire: (now+opt.Duration),
	}, nil
}

func genToken(iat int64, secretKey string, payloads map[string]any, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds  // 过期时间戳
	claims["iat"] = iat            // 签发时间戳

	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}