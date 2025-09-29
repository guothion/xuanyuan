package mapper

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	// 根据 ID 查询
	GetByID(ctx context.Context, id string) (any, error)
	// 创建
	Create(ctx context.Context, entity any) error
	// 更新
	Update(ctx context.Context, entity any) error
	// FindBy 查询（可扩展）
	FindBy(ctx context.Context, dest *[]*any, query string, args ...any) ([]any, error)
}

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{
		db: db,
	}
}

func (r *BaseRepository[T]) GetByID(ctx context.Context, id uint) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).First(entity, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("query failed %w", err)
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Create(ctx context.Context, entity *T) error {
	err := r.db.WithContext(ctx).Create(entity).Error
	return err
}

func (r *BaseRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r *BaseRepository[T]) FindBy(ctx context.Context, dest *[]*T, query string, args ...any) error {
	return r.db.WithContext(ctx).Where(query, args...).Find(dest).Error
}
