package model

type UserId uint

type Process struct {
	ID
	Code        ProcessCode `json:"code" gorm:"size:50;unique;comment:流程唯一编码;index:idx_code"`
	BizId       *string     `json:"biz_id" gorm:"size:64;comment:关联业务ID"`
	Name        string      `json:"name" gorm:"size:100;not null;comment:流程名称"`
	Description *string     `json:"description,omitempty" gorm:"type:text;comment:流程描述"`
	Status      string      `json:"status" gorm:"default:draft;comment:状态;index:idx_status"`
	Timestamp
}

func (Process) TableName() string {
	return "process"
}

func (Process) InsertableFields() FieldSet {
	return FieldSet{
		"code":        true,
		"biz_id":      true,
		"name":        true,
		"description": true,
		"status":      true,
	}
}

func (Process) UpdatableFields() FieldSet {
	return FieldSet{
		"name":        true,
		"description": true,
		"status":      true,
	}
}
