package controller

import (
	"context"
	"github.com/guothion/xuanyuan/internal/api/common/request"
	"github.com/guothion/xuanyuan/internal/service/stall"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/internal/api/common/response"
)

type StallsController struct{}

func (c *StallsController) GetStalls(ctx *gin.Context) {
	// 解析分页参数
	pageStr := ctx.DefaultQuery("page", "1")
	sizeStr := ctx.DefaultQuery("pageSize", "10")
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(sizeStr)

	// 查询分页数据
	stalls, total, err := stall.StallsService.GetPage(ctx, page, pageSize)
	if err != nil {
		response.BusinessFail(ctx, err.Error())
		return
	}

	response.Success(ctx, gin.H{
		"list":     stalls,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func (c *StallsController) Create(ctx *gin.Context) {
	var (
		req  request.StallCreate
		_ctx context.Context
		err  error
	)
	_ctx = ctx.Request.Context()

	if err = ctx.ShouldBind(&req); err != nil {
		response.ValidateFail(ctx, request.GetErrorMsg(req, err))
		return
	}

	if err = stall.StallsService.Create(_ctx, req); err != nil {
		response.BusinessFail(ctx, err.Error())
	}
	response.Success(ctx, nil)
}
