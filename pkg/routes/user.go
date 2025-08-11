package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/controller"
	"github.com/guothion/xuanyuan/pkg/api/middleware"
)

func init() {
	c := &UserRoute{}
	routes[c.BasePath()] = c
}

type UserRoute struct{}

func (u *UserRoute) BasePath() string { return "/user" }

func (u *UserRoute) RegisterRouter(router *gin.RouterGroup) {
	UserController := new(controller.UserController)
	userRouter := router.Group(u.BasePath()).
		Use(middleware.SessionRequireMiddleware)
	{
		// 创建用户接口
		userRouter.POST("/create", UserController.Create)
	}
}
