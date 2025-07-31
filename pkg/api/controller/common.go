package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/common"
	"github.com/guothion/xuanyuan/pkg/config"
)

// 解析 Context
func parseContext(ctx *gin.Context) (context *common.Context, err error) {
	rawValue, exists := ctx.Get(config.KeyRequestContext)
	if !exists {
		err = common.NewForbiddenError("request context not present")
		return
	}
	// 这里用到了类型断言
	context, exists = rawValue.(*common.Context)
	if !exists {
		err = common.NewForbiddenError("request context not present")
		return
	}
	return
}
