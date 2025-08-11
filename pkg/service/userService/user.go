package userService

import (
	"github.com/guothion/xuanyuan/pkg/common"
	"github.com/guothion/xuanyuan/pkg/mapper"
	"github.com/guothion/xuanyuan/pkg/model"
)

type userService struct{}

func (us *userService) Create(ctx *common.Context, req *model.User) (err error) {
	return mapper.User.Create(ctx, req)
}

func (us *userService) Update(ctx *common.Context, req *model.User) (err error) {
	return mapper.User.Update(ctx, req)
}
