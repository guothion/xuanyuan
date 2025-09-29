package dto

import "github.com/guothion/xuanyuan/pkg/model"

type StepDTO struct {
	ID
	ProcessCode   string        `json:"process_code"`
	StepKey       model.StepKey `json:"step_key"`
	Name          string        `json:"name"`
	Type          string        `json:"type"`
	Description   string        `json:"description"`
	Status        string        `json:"status"`
	SortOrder     int           `json:"sort_order"`
	Required      bool          `json:"required"`
	Assignees     []string      `json:"assignees"`
	DurationLimit int           `json:"duration_limit"`
	Tags          []string      `json:"tags"`
	Config        string        `json:"config"`
	Timestamp
}

type Process struct {
	ID
	Code        string    `json:"code"`
	BizId       string    `json:"biz_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Steps       []StepDTO `json:"steps"`
	Timestamp
}
