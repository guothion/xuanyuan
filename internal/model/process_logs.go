package model

type ProcessLogs struct {
	ID
	ProcessId    int     `json:"process_id" gorm:"comment:流程实例 ID;index:idx_process_id"`
	StepKey      StepKey `json:"step_key" gorm:"size:50;comment:操作的步骤 key;index:idx_step_key"`
	StepName     string  `json:"step_name" gorm:"size:100;comment:步骤显示名称，冗余便于查询"`
	OperatorId   int     `json:"operator_id" gorm:"comment:操作人用户 ID;index:idx_operator"`
	OperatorName string  `json:"operator_name" gorm:"size:50;comment:操作人姓名"`
	OperatorRole *string `json:"operator_role" gorm:"size:50;comment:操作人角色"`
	Action       string  `json:"action" gorm:"comment:操作类型;index:idx_action"`
	Comment      string  `json:"comment" gorm:"type:text;comment:操作备注或审批意见"`
	Description  string  `json:"description" gorm:"type:text;comment:系统自动生成的描述，如“系统因超时自动拒绝"`
	FromStatus   *string `json:"from_status" gorm:"size:20;comment:操作前状态"`
	ToStatus     *string `json:"to_status" gorm:"size:20;comment:操作后状态"`
	ClientIp     *string `json:"client_ip" gorm:"size:45;comment:操作 IP 地址"`
	UserAgent    *string `json:"user_agent" gorm:"type:text;comment:客户端信息"`
	ExtraData    *string `json:"extra_data" gorm:"type:json;comment:额外数据，如委派人、跳转规则等"`
	Timestamp
}

func (ProcessLogs) TableName() string {
	return "process_logs"
}

func (ProcessLogs) InsertableFields() FieldSet {
	return FieldSet{
		"process_id":    true,
		"step_key":      true,
		"step_name":     true,
		"operator_id":   true,
		"operator_name": true,
		"operator_role": true,
		"action":        true,
		"comment":       true,
		"description":   true,
		"from_status":   true,
		"to_status":     true,
		"client_ip":     true,
		"user_agent":    true,
		"extra_data":    true,
	}
}

func (ProcessLogs) UpdatableFields() FieldSet {
	return FieldSet{
		"step_name":     true,
		"operator_id":   true,
		"operator_name": true,
		"operator_role": true,
		"action":        true,
		"comment":       true,
		"description":   true,
		"from_status":   true,
		"to_status":     true,
		"client_ip":     true,
		"user_agent":    true,
		"extra_data":    true,
	}
}
