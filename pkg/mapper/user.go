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

func (m *userMapper) Get(ctx context.Context, userId int) (result *model.User, error) {
	var user *model.User
}
