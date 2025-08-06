package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/middleware"
	"github.com/guothion/xuanyuan/pkg/common"
	"github.com/guothion/xuanyuan/pkg/model"
	"github.com/guothion/xuanyuan/pkg/service/userService"
	"github.com/sirupsen/logrus"
	"net/http"
)

// 这里我们给controller设置key
func init() {
	c := &UserController{}
	controllers[c.BasePath()] = c
}

type UserController struct{}

func (u *UserController) BasePath() string { return "/v1/user" }

func (u *UserController) RegisterRouter(engine *gin.Engine) {
	routerGroup := engine.Group(u.BasePath())
	routerGroup.Use(middleware.SessionRequireMiddleware)
	routerGroup.GET("", u.List)

	routerGroup.POST("/create", u.Create)
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

// create user api
func (u *UserController) Create(ctx *gin.Context) {
	var (
		context *common.Context
		req     model.User
		err     error
	)

	if context, err = parseContext(ctx); err != nil {
		middleware.RespondForbidden(ctx)
		return
	}
	defer context.Cancel()

	if err = ctx.ShouldBindJSON(&req); err != nil {
		middleware.RespondBadRequest(ctx, err)
		return
	}

	if err = userService.UserService.Create(context, &req); err != nil {
		middleware.RespondFailure(ctx, err)
		return
	}

	middleware.RespondCreated(ctx)
}
