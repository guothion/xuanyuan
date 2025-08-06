package mapper

import (
	"context"
	"fmt"
	"github.com/guothion/xuanyuan/pkg/dataSource/mysql"
	"github.com/guothion/xuanyuan/pkg/model"
	"strings"
	"time"
)

var (
	_ = fmt.Errorf
	_ = strings.TrimSpace
)

type userMapper struct{}

func (m *userMapper) Create(ctx context.Context, user *model.User) (err error) {
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err = mysql.Creates(user)
	return
}
