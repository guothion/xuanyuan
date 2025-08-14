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
	userController := new(controller.UserController)
	userRouter := router.Group(u.BasePath()).
		Use(middleware.SessionRequireMiddleware)
	{
		// 注册
		userRouter.POST("/register", userController.Register)
		// 登录
		userRouter.POST("/login", userController.Login)
	}
}
