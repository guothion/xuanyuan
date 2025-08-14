package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/common/request"
	"github.com/guothion/xuanyuan/pkg/api/common/response"
	"github.com/guothion/xuanyuan/pkg/api/middleware"
	"github.com/guothion/xuanyuan/pkg/common"
	"github.com/guothion/xuanyuan/pkg/model"
	"github.com/guothion/xuanyuan/pkg/service/account"
)

type UserController struct{}

func (u *UserController) Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := account.UserService.Register(form); err != nil {
		response.ValidateFail(c, err.Error())
	} else {
		tokenData, err, _ := account.JwtService.CreateToken(account.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

func (u *UserController) Login(ctx *gin.Context) {
	var form request.Login
	if err := ctx.ShouldBind(&form); err != nil {
		response.ValidateFail(ctx, request.GetErrorMsg(form, err))
		return
	}

	if err, user := account.UserService.Login(form); err != nil {
		response.BusinessFail(ctx, err.Error())
	} else {
		tokenData, err, _ := account.JwtService.CreateToken(account.AppGuardName, user)
		if err != nil {
			response.BusinessFail(ctx, err.Error())
			return
		}
		response.Success(ctx, tokenData)
	}
}

func (u *UserController) Update(ctx *gin.Context) {
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

	if err = account.UserService.Update(context, &req); err != nil {
		middleware.RespondFailure(ctx, err)
		return
	}
	middleware.RespondUpdated(ctx)
}
