package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/middleware"
	"github.com/guothion/xuanyuan/pkg/global"
	"os"
)

func ShowConfig(ctx *gin.Context) {
	middleware.ResponseData(ctx, global.App.Config)
}

func ShowEnvVars(ctx *gin.Context) {
	middleware.ResponseData(ctx, os.Environ())
}
