package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/guothion/xuanyuan/internal/model"

	"gorm.io/gorm"
)

type Repository[T model.ModelWithFields] interface {
	// 根据 ID 查询
	GetByID(ctx context.Context, id uint, selectFields ...string) (*T, error)
	// 创建
	Create(ctx context.Context, model *T) error
	// 更新
	Update(ctx context.Context, model *T) error
	// 部分更新
	UpdatePartial(ctx context.Context, id uint, updates map[string]interface{}) error
	// 删除
	Delete(ctx context.Context, id uint) error
	// FindBy 查询（可扩展）
	FindBy(ctx context.Context, dest *[]*T, query string, args ...any) error
	// 获取某一条
	GetOne(ctx context.Context, conditions map[string]interface{}, selectFields []string, preloads ...string) (*T, error)
	// 获取列表
	GetList(ctx context.Context, conditions map[string]interface{}, selectFields []string, preloads ...string) ([]T, error)
	// 获取 page
	GetPage(ctx context.Context, page, pageSize int, conditions map[string]interface{}, order string, selectFields []string, preloads ...string) ([]T, int64, error)
	// 获取数量
	Count(ctx context.Context, conditions map[string]interface{}) (int64, error)
	// 是否存在
	Exists(ctx context.Context, conditions map[string]interface{}) (bool, error)
}

type BaseRepository[T model.ModelWithFields] struct {
	db *gorm.DB
}

func NewBaseRepository[T model.ModelWithFields](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{
		db: db,
	}
}

func (r *BaseRepository[T]) GetByID(ctx context.Context, id uint, selectFields ...string) (*T, error) {
	var modelInstance T
	query := r.db.WithContext(ctx)
	if len(selectFields) > 0 {
		query = query.Select(selectFields)
	}
	err := query.First(&modelInstance, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("query failed %w", err)
	}
	return &modelInstance, nil
}

func (r *BaseRepository[T]) Create(ctx context.Context, model *T) error {
	fields := (*model).InsertableFields()
	err := r.db.WithContext(ctx).Select(getSelectFields(fields)).Create(model).Error
	return err
}

func (r *BaseRepository[T]) Update(ctx context.Context, model *T) error {
	fields := (*model).UpdatableFields()
	return r.db.WithContext(ctx).
		Model(model).
		Select(getSelectFields(fields)).
		Updates(model).
		Error
}

func (r *BaseRepository[T]) UpdatePartial(ctx context.Context, id uint, updates map[string]interface{}) error {
	var modelInstance T
	query := r.db.WithContext(ctx)
	if err := query.First(&modelInstance, id).Error; err != nil {
		return err
	}

	updatable := modelInstance.UpdatableFields()
	filtered := make(map[string]interface{})
	for field, value := range updates {
		if updatable[field] {
			filtered[field] = value
		}
	}

	if len(filtered) == 0 {
		return errors.New("no updatable fields provided")
	}

	return query.Model(&modelInstance).Where("id = ?", id).Updates(filtered).Error
}

func (r *BaseRepository[T]) Delete(ctx context.Context, id uint) error {
	var modelInstance T
	return r.db.WithContext(ctx).Delete(&modelInstance, id).Error
}

func (r *BaseRepository[T]) FindBy(ctx context.Context, dest *[]*T, query string, args ...any) error {
	return r.db.WithContext(ctx).Where(query, args...).Find(dest).Error
}

// 辅助函数：从FieldSet获取字段列表
func getSelectFields(fields model.FieldSet) []string {
	var selected []string
	for field, enabled := range fields {
		if enabled {
			selected = append(selected, field)
		}
	}
	return selected
}

// GetOne 获取单条记录
func (r *BaseRepository[T]) GetOne(ctx context.Context, conditions map[string]interface{}, selectFields []string, preloads ...string) (*T, error) {
	var _model T
	query := r.db.WithContext(ctx).Where(conditions)

	if len(selectFields) > 0 {
		query = query.Select(selectFields)
	}

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.First(&_model).Error; err != nil {
		return nil, err
	}
	return &_model, nil
}

// GetList 获取多条记录
func (r *BaseRepository[T]) GetList(ctx context.Context, conditions map[string]interface{}, selectFields []string, preloads ...string) ([]T, error) {
	var models []T
	query := r.db.WithContext(ctx).Where(conditions)

	if len(selectFields) > 0 {
		query = query.Select(selectFields)
	}

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

// GetPage 分页查询
func (r *BaseRepository[T]) GetPage(ctx context.Context, page, pageSize int, conditions map[string]interface{}, order string, selectFields []string, preloads ...string) ([]T, int64, error) {
	var models []T
	var total int64

	query := r.db.Model(&models).WithContext(ctx).Where(conditions)

	if len(selectFields) > 0 {
		query = query.Select(selectFields)
	}

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if order != "" {
		query = query.Order(order)
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&models).Error; err != nil {
		return nil, 0, err
	}

	return models, total, nil
}

// Count 计数
func (r *BaseRepository[T]) Count(ctx context.Context, conditions map[string]interface{}) (int64, error) {
	var _model T
	var count int64
	query := r.db.Model(&_model).WithContext(ctx)
	if len(conditions) > 0 {
		query = query.Where(conditions)
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Exists 判断是否存在
func (r *BaseRepository[T]) Exists(ctx context.Context, conditions map[string]interface{}) (bool, error) {
	count, err := r.Count(ctx, conditions)
	return count > 0, err
}
