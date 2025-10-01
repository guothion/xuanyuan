package model

import "time"

type StallApplicationStatus string

const (
	StallApplicationStatusPending StallApplicationStatus = "pending"
	StallApplicationStatusRunning StallApplicationStatus = "running"
	StallApplicationStatusSuccess StallApplicationStatus = "success"
	StallApplicationStatusFailed  StallApplicationStatus = "failed"
)

type StallApplication struct {
	ID
	UserId      UserId                 `json:"user_id" gorm:"column:user_id;type:INT;comment:申请人用户ID;index:idx_user_id"`
	Mobile      string                 `json:"mobile" gorm:"column:mobile;type:VARCHAR(15);comment:联系电话"`
	Name        string                 `json:"name" gorm:"size:100;comment:摊位名称"`
	Description string                 `json:"description" gorm:"type:TEXT;comment:摊位描述"`
	Category    string                 `json:"category" gorm:"size:50;comment:经营分类"`
	CoverImage  string                 `json:"cover_image" gorm:"size:255;comment:封面图"`
	ImageUrls   string                 `json:"image_urls" gorm:"type:JSON;comment:多图（JSON数组）"`
	Latitude    float32                `json:"latitude" gorm:"type:DECIMAL(10, 8);comment:期望纬度"`
	Longitude   float32                `json:"longitude" gorm:"type:DECIMAL(11, 8);comment:期望经度"`
	StartTime   time.Time              `json:"start_time" gorm:"type:DATETIME;comment:计划开始时间;index:idx_start_time"`
	EndTime     time.Time              `json:"end_time" gorm:"type:DATETIME;comment:计划结束时间"`
	Status      StallApplicationStatus `json:"status" gorm:"type:varchar(255);comment:申请状态;index:idx_status"`
	Reason      string                 `json:"reason" gorm:"size:255;comment:拒绝原因"`
	Timestamp
	SoftDeletes
}

func (StallApplication) TableName() string {
	return "stall_application"
}

func (StallApplication) InsertableFields() FieldSet {
	return FieldSet{
		"user_id":     true,
		"mobile":      true,
		"name":        true,
		"description": true,
		"category":    true,
		"cover_image": true,
		"image_urls":  true,
		"latitude":    true,
		"longitude":   true,
		"start_time":  true,
		"end_time":    true,
		"status":      true,
		"reason":      true,
	}
}

func (StallApplication) UpdatableFields() FieldSet {
	return FieldSet{
		"mobile":      true,
		"name":        true,
		"description": true,
		"category":    true,
		"cover_image": true,
		"image_urls":  true,
		"latitude":    true,
		"longitude":   true,
		"start_time":  true,
		"end_time":    true,
		"status":      true,
		"reason":      true,
	}
}
