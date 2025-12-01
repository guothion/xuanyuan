package model

import "time"

type Stalls struct {
	ID
	UserId      UserId    `json:"user_id" gorm:"comment:摊主用户ID;index:idx_user_id"`
	Name        string    `json:"name" gorm:"size:100;comment:摊位名称"`
	Description string    `json:"description" gorm:"type:TEXT;comment:摊位描述"`
	Category    string    `json:"category" gorm:"size:50;comment:分类：小吃、手作、服饰等;index:idx_category"`
	CoverImage  string    `json:"cover_image" gorm:"size:255;comment:封面图"`
	Location    string    `json:"location" gorm:"type:JSON;comment:位置信息"`
	Status      string    `json:"status" gorm:"default:pending;comment:状态;index:idx_status"`
	Reason      string    `json:"reason" gorm:"size:255;comment:"原因"`
	StartTime   time.Time `json:"start_time" gorm:"type:DATETIME;comment:计划开始时间;index:idx_start_time"`
	EndTime     time.Time `json:"end_time" gorm:"type:DATETIME;comment:计划结束时间"`
	Timestamp
	SoftDeletes
}

func (Stalls) TableName() string {
	return "stalls"
}

func (Stalls) InsertableFields() FieldSet {
	return FieldSet{
		"user_id":     true,
		"name":        true,
		"description": true,
		"category":    true,
		"cover_image": true,
		"location":    true,
		"status":      true,
		"reason":      true,
		"start_time":  true,
		"end_time":    true,
	}
}

func (Stalls) UpdatableFields() FieldSet {
	return FieldSet{
		"description": true,
		"category":    true,
		"cover_image": true,
		"location":    true,
		"status":      true,
		"reason":      true,
		"start_time":  true,
		"end_time":    true,
	}
}
