package model

import "time"

type ID struct {
	ID uint `json:"id" gorm:"primary_key"`
}

type Timestamp struct {
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type SoftDeletes struct {
	DeletedAt time.Time `json:"deleted_at" gorm:"index"`
}
