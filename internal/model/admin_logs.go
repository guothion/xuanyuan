package model

type AdminLogs struct {
	ID
	AdminId    int    `json:"admin_id" gorm:"column:admin_id;"`
	Action     string `json:"action" gorm:"column:action;size:100;comment:操作：approve_stall, ban_user"`
	TargetType string `json:"target_type" gorm:"size:50"`
	TargetId   int    `json:"target_id" gorm:"column:target_id;"`
	Detail     string `json:"detail" gorm:"column:detail;type:text"`
	Ip         string `json:"ip" gorm:"column:ip;size:45"`
	Timestamp
}

func (AdminLogs) TableName() string {
	return "admin_logs"
}

func (AdminLogs) InsertableFields() FieldSet {
	return FieldSet{
		"admin_id":    true,
		"action":      true,
		"target_type": true,
		"target_id":   true,
		"detail":      true,
		"ip":          true,
	}
}

func (AdminLogs) UpdatableFields() FieldSet {
	return FieldSet{
		"action":      true,
		"target_type": true,
		"target_id":   true,
		"detail":      true,
	}
}
