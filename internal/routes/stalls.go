package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/internal/api/controller"
	"github.com/guothion/xuanyuan/internal/middleware"
	"github.com/guothion/xuanyuan/internal/service/account"
)

func init() {
	c := &StallRoute{}
	routes[c.BasePath()] = c
}

type StallRoute struct{}

func (u *StallRoute) BasePath() string { return "/stall" }

func (u *StallRoute) RegisterRouter(router *gin.RouterGroup) {
	stallController := new(controller.StallsController)

	stallRouter := router.Group(u.BasePath()).Use(middleware.JWTAuth(account.AppGuardName))
	{
		stallRouter.GET("", stallController.GetStalls)
		stallRouter.POST("/create", stallController.Create)
	}
}
