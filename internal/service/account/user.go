package account

import (
	"context"
	"errors"
	"github.com/guothion/xuanyuan/internal/repository"
	"strconv"

	"github.com/guothion/xuanyuan/internal/api/common/request"
	"github.com/guothion/xuanyuan/internal/model"
	"github.com/guothion/xuanyuan/internal/util"
)

type userService struct{}

func (us *userService) Register(ctx context.Context, params request.Register) (error, model.User) {
	isExit, err := repository.UserRepo.Exists(ctx, map[string]interface{}{"email": params.Email})
	if err != nil {
		return err, model.User{}
	}
	if isExit {
		return errors.New("当前邮箱已存在"), model.User{}
	}
	user := &model.User{
		Username: params.Name,
		Password: params.Password,
		Email:    &params.Email,
		Role:     params.Role,
	}
	err = repository.UserRepo.Create(ctx, user)
	return err, *user
}

func (us *userService) Login(ctx context.Context, params request.Login) (err error, user *model.User) {
	user, err = repository.UserRepo.GetByUserOrName(ctx, params.Email, params.Username)
	if err != nil {
		return
	}
	if isOk := util.BcryptMakeCheck([]byte(params.Password), user.Password); !isOk {
		err = errors.New("密码错误")
		return
	}
	return
}

func (us *userService) GetUserInfo(ctx context.Context, uid string) (err error, user *model.User) {
	intId, err := strconv.Atoi(uid)
	user, err = repository.UserRepo.GetByID(ctx, uint(intId), "id,username,email,role,created_at,updated_at")
	if err != nil {
		err = errors.New("当前用户不存在")
	}
	return
}
