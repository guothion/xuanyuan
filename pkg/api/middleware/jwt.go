package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/guothion/xuanyuan/pkg/api/common/response"
	"github.com/guothion/xuanyuan/pkg/global"
	"github.com/guothion/xuanyuan/pkg/service/account"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.TokenFail(c)
			c.Abort()
			return
		}
		tokenStr = tokenStr[len(account.TokenTypeBearer)+1:]

		token, err := jwt.ParseWithClaims(tokenStr, &account.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})

		if err != nil || account.JwtService.IsInBlacklist(tokenStr) {
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*account.CustomClaims)
		if claims.Issuer != GuardName {
			response.TokenFail(c)
			c.Abort()
			return
		}

		if claims.ExpiresAt.Unix()-time.Now().Unix() < global.App.Config.Jwt.RefreshGracePeriod {
			lock := global.Lock("refresh_token_lock", global.App.Config.Jwt.JwtBlacklistGracePeriod)
			if lock.Get() {
				err, user := account.JwtService.GetUserInfo(GuardName, claims.ID)
				if err != nil {
					logrus.Error(err.Error())
					lock.Release()
				} else {
					tokenData, _, _ := account.JwtService.CreateToken(GuardName, user)
					c.Header("new-token", tokenData.AccessToken)
					c.Header("new-expires-in", strconv.Itoa(int(tokenData.ExpiresIn)))
					_ = account.JwtService.JoinBlackList(token)
				}
			}
		}

		c.Set("token", token)
		c.Set("id", claims.ID)
	}
}
