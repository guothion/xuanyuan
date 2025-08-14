package account

import (
	"errors"
	"github.com/guothion/xuanyuan/pkg/api/common/request"
	"github.com/guothion/xuanyuan/pkg/common"
	"github.com/guothion/xuanyuan/pkg/mapper"
	"github.com/guothion/xuanyuan/pkg/model"
	"github.com/guothion/xuanyuan/pkg/util"
)

type userService struct{}

func (us *userService) Update(ctx *common.Context, req *model.User) (err error) {
	return mapper.User.Update(ctx, req)
}

func (us *userService) Register(params request.Register) (err error, user model.User) {
	if err = mapper.User.GetUserIDByEmail(params.Email); err != nil {
		return
	}
	err, user = mapper.User.CreateUser(params)
	return
}

func (us *userService) Login(params request.Login) (err error, user *model.User) {
	err, user = mapper.User.GetUserInfoByEmail(params.Email)
	if err != nil {
		err = errors.New("用户名不存在")
		return
	}
	if isOk := util.BcryptMakeCheck([]byte(params.Password), user.PasswordHash); !isOk {
		err = errors.New("密码错误")
		return
	}
	return
}
