package repository

import (
	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type StepRepository struct {
	*BaseRepository[model.Step]
}

func NewStepRepository(db *gorm.DB) *StepRepository {
	return &StepRepository{
		BaseRepository: &BaseRepository[model.Step]{
			db: db,
		},
	}
}
