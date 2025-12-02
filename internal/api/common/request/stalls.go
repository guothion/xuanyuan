package request

import (
	"github.com/guothion/xuanyuan/internal/model"
	"time"
)

type StallCreate struct {
	UserId      model.UserId `form:"user_id" json:"user_id" binding:"required"`
	Name        string       `form:"name" json:"name" binding:"required"`
	Description string       `form:"description" json:"description" binding:"required"`
	Category    string       `form:"category" json:"category" binding:"required"`
	CoverImage  string       `form:"cover_image" json:"cover_image" binding:"required,uri"`
	Location    string       `form:"location" json:"location" binding:"required"`
	Status      string       `form:"status" json:"status" binding:"required"`
	Reason      string       `form:"reason" json:"reason" binding:"required"`
	StartTime   time.Time    `form:"start_time" json:"start_time" binding:"required"`
	EndTime     time.Time    `form:"end_time" json:"end_time" binding:"required"`
}

func (StallCreate StallCreate) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"user_id.required":     "user id is required",
		"name.required":        "name is required",
		"description.required": "description is required",
		"category.required":    "category is required",
		"cover_image.required": "cover_image is required",
		"cover_image.uri":      "cover_image is not a uri",
		"location.required":    "location is required",
		"status.required":      "status is required",
		"reason.required":      "reason is required",
		"start_time.required":  "start_time is required",
		"end_time.required":    "end_time is required",
	}
}
