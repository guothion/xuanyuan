package repository

import (
	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type AdminLogsRepository struct {
	*BaseRepository[model.AdminLogs]
}

func NewAdminLogsRepository(db *gorm.DB) *AdminLogsRepository {
	return &AdminLogsRepository{
		BaseRepository: NewBaseRepository[model.AdminLogs](db),
	}
}
