package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/internal/api/controller"
	"github.com/guothion/xuanyuan/internal/middleware"
	"github.com/guothion/xuanyuan/internal/service/account"
)

func init() {
	c := &AdminLogsRoute{}
	routes[c.BasePath()] = c
}

type AdminLogsRoute struct {
}

func (u *AdminLogsRoute) BasePath() string { return "/admin_logs" }

func (u *AdminLogsRoute) RegisterRouter(router *gin.RouterGroup) {
	adminLogsController := new(controller.AdminLogsController)
	adminLogsRouter := router.Group(u.BasePath()).Use(middleware.JWTAuth(account.AppGuardName))
	{
		adminLogsRouter.GET("/info/:id", adminLogsController.Get)
		adminLogsRouter.POST("/create", adminLogsController.Create)
		adminLogsRouter.GET("", adminLogsController.GetList)
	}
}
