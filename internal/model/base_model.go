package model

import (
	"gorm.io/gorm"
	"time"
)

type ID struct {
	ID uint `json:"id" gorm:"primary_key;auto_increment;comment:ID"`
}

type CreatedBy struct {
	CreatedBy string `json:"created_by" gorm:"comment:Created By"`
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
}

type SoftDeletes struct {
	IsDeleted int8           `json:"is_deleted" gorm:"size:8,comment:是否删除"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"comment:删除时间"`
}

type ListCommon struct {
	Total    int64 `json:"total"`
	PageSize int   `json:"page_size"`
	Page     int   `json:"page"`
}

// FieldSet 可更新字段集合
type FieldSet map[string]bool

type ModelWithFields interface {
	TableName() string
	InsertableFields() FieldSet // 可插入字段
	UpdatableFields() FieldSet  // 可更新字段
}
