package model

type StepKey string
type ProcessCode string

type Step struct {
	ID
	// Type 步骤类型  例如：AskingApproval 与 Operation 两种
	ProcessCode   ProcessCode `json:"process_code" gorm:"size:50;comment:关联的流程 process_code;index:idx_process_code;uniqueIndex:uk_process_step"`
	StepKey       StepKey     `json:"step_key" gorm:"size:50;comment:步骤唯一标识，如 submit, approve_hr, approve_mgr;uniqueIndex:uk_process_step"`
	Name          string      `json:"name" gorm:"size:100;comment:步骤显示名称，如 提交申请、HR 审批"`
	Type          string      `json:"type" gorm:"size:30;comment:步骤类型：approval, notification, auto_task, decision 等"`
	Description   *string     `json:"description,omitempty" gorm:"comment:描述"`
	Status        StatusEnum  `json:"status" gorm:"default:pending;comment:节点状态"`
	SortOrder     int         `json:"sort_order" gorm:"default:0;comment:排序序号，用于确定执行顺序"`
	Required      int8        `json:"required" gorm:"size:2;default:1;comment:是否必须完成"`
	Assignees     string      `json:"assignees" gorm:"comment:可以处理当前节点的人"`
	DurationLimit int         `json:"duration_limit" gorm:"comment:处理时限（分钟），超时可提醒"`
	Tags          *string     `json:"tags" gorm:"type:json;comment:标签数组"`
	Config        *string     `json:"config" gorm:"comment:额外配置，如表单字段、条件表达式等"`
	Timestamp
}

func (Step) TableName() string {
	return "process_steps"
}

func (Step) InsertableFields() FieldSet {
	return FieldSet{
		"process_code":   true,
		"step_key":       true,
		"name":           true,
		"type":           true,
		"description":    true,
		"status":         true,
		"sort_order":     true,
		"required":       true,
		"assignees":      true,
		"duration_limit": true,
		"tags":           true,
		"config":         true,
	}
}

func (Step) UpdatableFields() FieldSet {
	return FieldSet{
		"name":           true,
		"type":           true,
		"description":    true,
		"status":         true,
		"sort_order":     true,
		"required":       true,
		"assignees":      true,
		"duration_limit": true,
		"tags":           true,
		"config":         true,
	}
}

func (m *Step) SetStatus(newStatus StatusEnum) {
	m.Status = newStatus
}
