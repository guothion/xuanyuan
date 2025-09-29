package mapper

import (
	"context"

	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	*BaseRepository[model.User]
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	base := NewBaseRepository[model.User](db)
	return &UserRepo{BaseRepository: base}
}

func (r *UserRepo) GetByID(ctx context.Context, id uint) (*model.User, error) {
	return r.BaseRepository.GetByID(ctx, id)
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	return r.BaseRepository.Create(ctx, user)
}

func (r *UserRepo) Update(ctx context.Context, user *model.User) error {
	return r.BaseRepository.Update(ctx, user)
}
