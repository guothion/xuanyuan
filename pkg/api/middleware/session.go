package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/common"
	"github.com/guothion/xuanyuan/pkg/config"
	"net/http"
	"time"
)

func SessionRequireMiddleware(ctx *gin.Context) {
	account := ctx.GetHeader(config.KeyHeaderAccount)
	accessToken := ctx.GetHeader(config.KeyHeaderAccessToken)
	if err := SessionInstance.ValidateAccess(accessToken, account); err != nil {
		response := common.Status{
			Code:    http.StatusForbidden,
			Message: "invalid access token and account",
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, response)
		return
	}

	serviceProfile := ctx.GetHeader(config.KeyHeaderServiceProfile)
	requestContext := common.NewContext(time.Minute * 2).
		WithAccountName(account).
		WithAccessToken(accessToken).
		WithServiceProfile(serviceProfile)

	ctx.Set(config.KeyRequestContext, requestContext)
}
