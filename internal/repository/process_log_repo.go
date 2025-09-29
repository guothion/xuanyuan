package repository

import (
	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type ProcessLogRepository struct {
	*BaseRepository[model.ProcessLogs]
}

func NewProcessLogRepository(db *gorm.DB) *ProcessLogRepository {
	return &ProcessLogRepository{
		BaseRepository: NewBaseRepository[model.ProcessLogs](db),
	}
}
