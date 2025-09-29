package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/internal/api/common/response"
	"github.com/guothion/xuanyuan/internal/global"
	"os"
)

type AdminController struct{}

func (ac *AdminController) ShowConfig(ctx *gin.Context) {
	response.Success(ctx, global.App.Config)
}

func (ac *AdminController) ShowEnvVars(ctx *gin.Context) {
	response.Success(ctx, os.Environ())
}
