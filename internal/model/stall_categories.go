package model

type StallCategoryStatus string

type StallCategories struct {
	ID
	Name      string              `json:"name" gorm:"type:varchar(50);unique;comment:分类名：小吃、手作、鲜花等;"`
	Icon      string              `json:"icon" gorm:"size:255;comment:图表"`
	SortOrder int                 `json:"sort_order" gorm:"default:0"`
	Status    StallCategoryStatus `json:"status" gorm:"default:active"`
	Timestamp
	SoftDeletes
}

func (StallCategories) TableName() string {
	return "stall_categories"
}

func (StallCategories) InsertableFields() FieldSet {
	return FieldSet{
		"name":       true,
		"icon":       true,
		"sort_order": true,
		"status":     true,
	}
}

func (StallCategories) UpdatableFields() FieldSet {
	return FieldSet{
		"name":       true,
		"icon":       true,
		"sort_order": true,
		"status":     true,
	}
}
