package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/middleware"
	"github.com/guothion/xuanyuan/pkg/global"
	"os"
)

type AdminController struct{}

func (ac *AdminController) ShowConfig(ctx *gin.Context) {
	middleware.ResponseData(ctx, global.App.Config)
}

func (ac *AdminController) ShowEnvVars(ctx *gin.Context) {
	middleware.ResponseData(ctx, os.Environ())
}
