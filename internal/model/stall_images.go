package model

type StallImages struct {
	ID
	StallId   int    `json:"stall_id" gorm:"column:stall_id;index:idx_stall_id"`
	ImageUrl  string `json:"image_url" gorm:"column:image_url;size:255"`
	SortOrder int    `json:"sort_order" gorm:"column:sort_order;default:0"`
	Timestamp
	SoftDeletes
}

func (StallImages) TableName() string {
	return "stall_images"
}

func (StallImages) InsertableFields() FieldSet {
	return FieldSet{
		"stall_id":   true,
		"image_url":  true,
		"sort_order": true,
	}
}

func (StallImages) UpdatableFields() FieldSet {
	return FieldSet{
		"image_url":  true,
		"sort_order": true,
	}
}
