package mapper

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/guothion/xuanyuan/internal/api/common/request"
	"github.com/guothion/xuanyuan/internal/global"
	"github.com/guothion/xuanyuan/internal/model"
	"github.com/guothion/xuanyuan/internal/util"
	"gorm.io/gorm"
)

var (
	_ = fmt.Errorf
	_ = strings.TrimSpace
)

type userMapper struct {
}

func (m *userMapper) GetUserIDByEmail(ctx context.Context, email string) (err error) {
	result := global.App.DB.WithContext(ctx).Where("email = ?", email).Select("id").First(&model.User{})
	if result.RowsAffected != 0 {
		err = errors.New("邮箱已经存在")
		return
	}
	return nil
}

func (m *userMapper) GetUserInfo(ctx context.Context, email string, username string) (user *model.User, err error) {
	err = global.App.DB.WithContext(ctx).Where("email = ?", email).
		Or("username = ?", username).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("用户不存在")
		} else {
			err = errors.New("查询出错")
		}
		return
	}
	return
}

func (m *userMapper) CreateUser(ctx context.Context, ur request.Register) (user model.User, err error) {
	user = model.User{
		Username: ur.Name,
		Email:    &ur.Email,
		Password: util.BcryptMake([]byte(ur.Password)),
		Role:     ur.Role,
	}
	err = global.App.DB.WithContext(ctx).Create(&user).Error
	return
}

func (m *userMapper) GetUserInfoById(ctx context.Context, uid int) (err error, user *model.User) {
	err = global.App.DB.WithContext(ctx).Select("id", "username", "email", "role", "created_at", "updated_at").First(&user, uid).Error
	return
}
