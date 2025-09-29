package repository

import (
	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type ProcessRepository struct {
	*BaseRepository[model.Process]
}

func NewProcessRepository(db *gorm.DB) *ProcessRepository {
	return &ProcessRepository{
		BaseRepository: NewBaseRepository[model.Process](db),
	}
}
