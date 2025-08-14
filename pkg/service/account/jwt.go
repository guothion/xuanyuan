package account

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/guothion/xuanyuan/pkg/global"
)

type jwtService struct{}

type JwtUser interface {
	GetUid() string
}

type CustomClaims struct {
	jwt.RegisteredClaims
}

const (
	TokenTypeBearer = "bearer"
	AppGuardName    = "app"
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (jwtService *jwtService) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutPut, err error, token *jwt.Token) {
	// 使用HS256算法替代ES256,因为ES256需要使用ECDSA密钥对
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.App.Config.Jwt.JwtTtl))),
				ID:        user.GetUid(),
				Issuer:    GuardName,
				NotBefore: jwt.NewNumericDate(time.Now().Add(-1000 * time.Second)),
			},
		},
	)

	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret))
	if err != nil {
		return
	}

	tokenData = TokenOutPut{
		tokenStr,
		int(global.App.Config.Jwt.JwtTtl),
		TokenTypeBearer,
	}
	return
}
