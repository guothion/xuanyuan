package model

type CommentStatus string

type Comments struct {
	ID
	UserId   int           `json:"user_id"`
	StallId  int           `json:"stall_id" gorm:"column:stall_id;index:idx_stall_id"`
	Content  string        `json:"content" gorm:"type:text"`
	ParentId int           `json:"parent_id" gorm:"default:null;comment:回复评论 Id;index:idx_parent_id"`
	Status   CommentStatus `json:"status" gorm:"default:normal"`
	Timestamp
	SoftDeletes
}

func (Comments) TableName() string {
	return "comments"
}

func (Comments) InsertableFields() FieldSet {
	return FieldSet{
		"user_id":   true,
		"stall_id":  true,
		"content":   true,
		"parent_id": true,
		"status":    true,
	}
}

func (Comments) UpdatableFields() FieldSet {
	return FieldSet{
		"user_id":   true,
		"stall_id":  true,
		"content":   true,
		"parent_id": true,
		"status":    true,
	}
}
