package repository

import (
	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type NotificationsRepository struct {
	*BaseRepository[model.Notifications]
}

func NewNotificationsRepository(db *gorm.DB) *NotificationsRepository {
	return &NotificationsRepository{
		BaseRepository: NewBaseRepository[model.Notifications](db),
	}
}
