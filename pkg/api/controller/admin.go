package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/middleware"
	"github.com/guothion/xuanyuan/pkg/config"
	"os"
)

func ShowConfig(ctx *gin.Context) {
	middleware.ResponseData(ctx, config.Config)
}

func ShowEnvVars(ctx *gin.Context) {
	middleware.ResponseData(ctx, os.Environ())
}
