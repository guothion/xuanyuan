package model

type Notifications struct {
	ID
	UserId     int    `json:"user_id" gorm:"index:idx_user_read"`
	Title      string `json:"title" gorm:"column:title;size:100"`
	Content    string `json:"content" gorm:"column:content;type:text"`
	ReadStatus bool   `json:"read_status" gorm:"column:read_status;type:bool;default:false;index:idx_user_read"`
	TargetType string `json:"target_type" gorm:"column:target_type;size:50;comment:关联类型：stall, comment"`
	TargetId   int    `json:"target_id" gorm:"comment:关联 ID"`
	Timestamp
}

func (Notifications) TableName() string {
	return "notifications"
}

func (Notifications) InsertableFields() FieldSet {
	return FieldSet{
		"user_id":     true,
		"title":       true,
		"content":     true,
		"read_status": true,
		"target_type": true,
		"target_id":   true,
	}
}

func (Notifications) UpdatableFields() FieldSet {
	return FieldSet{
		"title":       true,
		"content":     true,
		"read_status": true,
		"target_type": true,
	}
}
