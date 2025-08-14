package model

import (
	"gorm.io/gorm"
	"time"
)

type ID struct {
	ID uint `json:"id" gorm:"primary_key;auto_increment;comment:ID"`
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
