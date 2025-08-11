package mapper

import (
	"context"
	"fmt"
	"github.com/guothion/xuanyuan/pkg/model"
	"strings"
)

var (
	_ = fmt.Errorf
	_ = strings.TrimSpace
)

type userMapper struct{}

func (m *userMapper) Create(ctx context.Context, user *model.User) (err error) {

	return nil
}

func (m *userMapper) Update(ctx context.Context, user *model.User) (err error) {

	return nil
}
