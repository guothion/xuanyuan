package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/common"
	"net/http"
)

func ResponseData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func RespondCreated(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, common.StatusOk)
}
