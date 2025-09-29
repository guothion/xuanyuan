package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/internal/api/common/request"
	"github.com/guothion/xuanyuan/internal/api/common/response"
	"github.com/guothion/xuanyuan/internal/model"
	"github.com/guothion/xuanyuan/internal/service/account"
)

type UserController struct{}

func (u *UserController) Register(ctx *gin.Context) {
	var (
		req  request.Register
		_ctx context.Context
		err  error
		user model.User
	)
	_ctx = ctx.Request.Context()

	if err = ctx.ShouldBind(&req); err != nil {
		response.ValidateFail(ctx, request.GetErrorMsg(req, err))
		return
	}

	if err, user = account.UserService.Register(_ctx, req); err != nil {
		response.ValidateFail(ctx, err.Error())
	} else {
		var tokenData account.TokenOutPut
		tokenData, err, _ = account.JwtService.CreateToken(account.AppGuardName, user)
		if err != nil {
			response.BusinessFail(ctx, err.Error())
			return
		}
		response.Success(ctx, tokenData)
	}
}

func (u *UserController) Login(ctx *gin.Context) {
	var (
		req       request.Login
		_ctx      context.Context
		err       error
		user      *model.User
		tokenData account.TokenOutPut
	)
	_ctx = ctx.Request.Context()
	if err = ctx.ShouldBind(&req); err != nil {
		response.ValidateFail(ctx, request.GetErrorMsg(req, err))
		return
	}

	if err, user = account.UserService.Login(_ctx, req); err != nil {
		response.BusinessFail(ctx, err.Error())
	} else {
		tokenData, err, _ = account.JwtService.CreateToken(account.AppGuardName, user)
		if err != nil {
			response.BusinessFail(ctx, err.Error())
			return
		}
		// set authorization cookie
		ctx.SetCookie(
			"Authorization",
			tokenData.AccessToken,
			tokenData.ExpiresIn,
			"/",
			"localhost",
			false,
			true,
		)
		type LoginResp struct {
			AccessToken string `json:"access_token"`
			ExpiresIn   int    `json:"expires_in"`
			Type        string `json:"type"`
		}
		resp := LoginResp{
			AccessToken: tokenData.AccessToken,
			ExpiresIn:   tokenData.ExpiresIn,
			Type:        req.Type,
		}
		response.Success(ctx, resp)
	}
}

func (u *UserController) Info(ctx *gin.Context) {
	var (
		_ctx context.Context
		err  error
		user *model.User
	)
	_ctx = ctx.Request.Context()
	err, user = account.UserService.GetUserInfo(_ctx, ctx.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(ctx, err.Error())
		return
	}
	response.Success(ctx, user)
}
