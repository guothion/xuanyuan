package model

type Favorites struct {
	ID
	UserId  int `json:"user_id" gorm:"column:user_id;uniqueIndex;composite:uk_user_stall"`
	StallId int `json:"stall_id" gorm:"column:stall_id;uniqueIndex;composite:uk_user_stall"`
	Timestamp
}

func (Favorites) TableName() string {
	return "favorites"
}

func (Favorites) InsertableFields() FieldSet {
	return FieldSet{
		"user_id":  true,
		"stall_id": true,
	}
}

func (Favorites) UpdatableFields() FieldSet {
	return FieldSet{
		"user_id":  true,
		"stall_id": true,
	}
}
