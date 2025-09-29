package repository

import (
	"context"

	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	*BaseRepository[model.User]
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository[model.User](db),
	}
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string, selectFields ...string) (*model.User, error) {
	return r.GetOne(ctx, map[string]interface{}{"email": email}, selectFields)
}

func (r *UserRepository) GetByUserOrName(ctx context.Context, email string, name string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		Or("username=?", name).
		First(&user).Error
	return &user, err
}
