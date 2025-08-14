package mapper

import (
	"errors"
	"fmt"
	"github.com/guothion/xuanyuan/pkg/api/common/request"
	"github.com/guothion/xuanyuan/pkg/global"
	"github.com/guothion/xuanyuan/pkg/model"
	"github.com/guothion/xuanyuan/pkg/util"
	"strings"
)

var (
	_ = fmt.Errorf
	_ = strings.TrimSpace
)

type userMapper struct{}

func (m *userMapper) GetUserIDByEmail(email string) (err error) {
	result := global.App.DB.Where("email = ?", email).Select("id").First(&model.User{})
	if result.RowsAffected != 0 {
		err = errors.New("邮箱已经存在")
		return
	}
	return nil
}

func (m *userMapper) GetUserInfoByEmail(email string) (err error, user *model.User) {
	err = global.App.DB.Where("email = ?", email).First(&user).Error
	return
}

func (m *userMapper) CreateUser(ur request.Register) (err error, user model.User) {
	user = model.User{
		Username:     ur.Name,
		Email:        &ur.Email,
		PasswordHash: util.BcryptMake([]byte(ur.Password)),
		Role:         ur.Role,
	}
	err = global.App.DB.Create(&user).Error
	return
}

func (m *userMapper) GetUserInfoById(uid int) (err error, user *model.User) {
	err = global.App.DB.Select("id", "username", "email", "role").First(&user, uid).Error
	return
}
