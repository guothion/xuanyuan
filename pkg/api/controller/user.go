package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/middleware"
	"github.com/guothion/xuanyuan/pkg/common"
	"github.com/guothion/xuanyuan/pkg/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

// 这里我们给controller设置key
func init() {
	c := &UserController{}
	controllers[c.BasePath()] = c
}

type UserController struct{}

func (u *UserController) BasePath() string { return "/v1/userService" }

func (u *UserController) RegisterRouter(engine *gin.Engine) {
	routerGroup := engine.Group(u.BasePath())
	routerGroup.Use(middleware.SessionRequireMiddleware)
	routerGroup.GET("", u.List)

	routerGroup.GET("/:name", u.UserName)
}

func (u *UserController) List(ctx *gin.Context) {
	var (
		context *common.Context
		result  *model.ListUserResponse
		err     error
	)

	logrus.Println(context, result, err)
	ctx.JSON(http.StatusOK, result)
}

func (u *UserController) UserName(ctx *gin.Context) {
	name := ctx.Param("name")

	ctx.String(http.StatusOK, "name="+name)
}
