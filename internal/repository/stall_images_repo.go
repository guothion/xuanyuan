package repository

import (
	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type StallImagesRepository struct {
	*BaseRepository[model.StallImages]
}

func NewStallImagesRepository(db *gorm.DB) *StallImagesRepository {
	return &StallImagesRepository{
		BaseRepository: NewBaseRepository[model.StallImages](db),
	}
}
