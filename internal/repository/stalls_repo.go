package repository

import (
	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type StallsRepository struct {
	*BaseRepository[model.Stalls]
}

func NewStallsRepository(db *gorm.DB) *StallsRepository {
	return &StallsRepository{
		BaseRepository: NewBaseRepository[model.Stalls](db),
	}
}
