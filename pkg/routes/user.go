package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/controller"
	"github.com/guothion/xuanyuan/pkg/api/middleware"
	"github.com/guothion/xuanyuan/pkg/service/account"
)

func init() {
	c := &UserRoute{}
	routes[c.BasePath()] = c
}

type UserRoute struct{}

func (u *UserRoute) BasePath() string { return "/user" }

func (u *UserRoute) RegisterRouter(router *gin.RouterGroup) {
	userController := new(controller.UserController)
	// 注册
	router.POST("/register", userController.Register)
	// 登录
	router.POST("/login", userController.Login)

	userRouter := router.Group(u.BasePath()).Use(middleware.JWTAuth(account.AppGuardName))
	{
		userRouter.GET("/info", userController.Info)
	}
}
