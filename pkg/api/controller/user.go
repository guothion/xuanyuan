package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/middleware"
	"github.com/guothion/xuanyuan/pkg/common"
	"github.com/guothion/xuanyuan/pkg/model"
	"github.com/sirupsen/logrus"
)

func init() {

}

type UserController struct{}

func (u *UserController) BasePath() string { return "/v1/user" }

func (u *UserController) ReginsterRouter(engine *gin.Engine) {
	routerGroup := engine.Group(u.BasePath())
	routerGroup.Use(middleware.SessionRequireMiddleware)

	routerGroup.GET("", u.List)
}

func (u *UserController) List(ctx *gin.Context) {
	var (
		context *common.Context
		result  *model.ListUserResponse
		err     error
	)

	logrus.Println(context, result, err)

}
