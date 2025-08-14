package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/common/request"
	"github.com/guothion/xuanyuan/pkg/api/common/response"
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

func (u *UserController) Info(ctx *gin.Context) {
	err, user := account.UserService.GetUserInfo(ctx.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(ctx, err.Error())
		return
	}
	response.Success(ctx, user)
}
