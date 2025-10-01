package repository

import (
	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type StallApplicationRepository struct {
	*BaseRepository[model.StallApplication]
}

func NewStallApplicationRepository(db *gorm.DB) *StallApplicationRepository {
	return &StallApplicationRepository{
		BaseRepository: NewBaseRepository[model.StallApplication](db),
	}
}
