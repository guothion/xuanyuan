package middleware

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/bootstrap"
	"github.com/guothion/xuanyuan/pkg/util"
	"github.com/sirupsen/logrus"
	"time"
)

// 什么是 JWT ？
// https://github.com/appleboy/gin-jwt/blob/master/_example/basic/server.go

const (
	AuthJWTKey         = "VSr#XavWJ^p35Y~h*V"
	AuthJWTRealm       = "xuanyuan-service"
	AuthJWTIdentityKey = "xuanyuan-identity"

	JWTPropUserID   = "id"
	JWTPropUserName = "name"
)

func createJWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       AuthJWTRealm,
		Key:         []byte(AuthJWTKey),
		Timeout:     time.Hour * 7 * 24,
		MaxRefresh:  time.Hour * 12,
		IdentityKey: AuthJWTIdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*bootstrap.LoginIdentity); ok {
				return jwt.MapClaims{
					JWTPropUserID:   v.ID,
					JWTPropUserName: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return claims
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginRequest bootstrap.LoginRequest
			if err := c.ShouldBind(&loginRequest); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			loginIndentity, err := bootstrap.Session.Login(&loginRequest)
			if err != nil {
				logrus.Errorf("Authenticate fail for: %v,error:%v", loginRequest, err)
				return nil, err
			}

			logrus.Infof("New login session from %s for %s", util.ParseRemoteRealIP(c.Request), loginIndentity)
			return loginIndentity, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			_, ok := data.(*jwt.GinJWTMiddleware)
			if !ok {
				logrus.Warningf("Authorize fail to: %v", c.Request.URL)
				return false
			}
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
			parseError := &gin.Error{
				Err: errors.New(message),
			}
			c.Errors = append(c.Errors, parseError)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	return authMiddleware, err
}

var jwtMiddleware *jwt.GinJWTMiddleware

func initJWTMiddleware() {
	var err error
	if jwtMiddleware, err = createJWTMiddleware(); err != nil {
		logrus.Fatalf("Init jwt failed:%v", err)
	}

	if err = jwtMiddleware.MiddlewareInit(); err != nil {
		logrus.Fatalf("jwtMiddlwware.MiddlewareInit() failed:%v", err.Error())
	}
}

func LoginHandler(c *gin.Context) {
	jwtMiddleware.LoginHandler(c)
}

func LoginRequired() gin.HandlerFunc {
	return jwtMiddleware.MiddlewareFunc()
}
