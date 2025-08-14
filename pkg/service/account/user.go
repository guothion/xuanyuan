package account

import (
	"errors"
	"github.com/guothion/xuanyuan/pkg/api/common/request"
	"github.com/guothion/xuanyuan/pkg/mapper"
	"github.com/guothion/xuanyuan/pkg/model"
	"github.com/guothion/xuanyuan/pkg/util"
	"strconv"
)

type userService struct{}

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

func (us *userService) GetUserInfo(uid string) (err error, user *model.User) {
	intId, err := strconv.Atoi(uid)
	err, user = mapper.User.GetUserInfoById(intId)
	if err != nil {
		err = errors.New("当前用户不存在")
	}
	return
}
