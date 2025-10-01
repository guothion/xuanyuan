package repository

import (
	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type StallCategoriesRepository struct {
	*BaseRepository[model.StallCategories]
}

func NewStallCategoriesRepository(db *gorm.DB) *StallCategoriesRepository {
	return &StallCategoriesRepository{
		BaseRepository: NewBaseRepository[model.StallCategories](db),
	}
}
