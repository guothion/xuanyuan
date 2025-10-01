package repository

import (
	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type CommentsRepository struct {
	*BaseRepository[model.Comments]
}

func NewCommentsRepository(db *gorm.DB) *CommentsRepository {
	return &CommentsRepository{
		BaseRepository: NewBaseRepository[model.Comments](db),
	}
}
