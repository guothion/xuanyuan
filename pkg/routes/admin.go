package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/controller"
)

func init() {
	c := &AdminRoute{}
	routes[c.BasePath()] = c
}

type AdminRoute struct{}

func (ar *AdminRoute) BasePath() string { return "/conf" }

func (ar *AdminRoute) RegisterRouter(router *gin.RouterGroup) {
	adminController := &controller.AdminController{}
	router.GET("/config", adminController.ShowConfig)
	router.GET("/env-vars", adminController.ShowEnvVars)
}
