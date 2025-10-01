package repository

import (
	"github.com/guothion/xuanyuan/internal/model"
	"gorm.io/gorm"
)

type FavoritesRepository struct {
	*BaseRepository[model.Favorites]
}

func NewFavoritesRepository(db *gorm.DB) *FavoritesRepository {
	return &FavoritesRepository{
		BaseRepository: NewBaseRepository[model.Favorites](db),
	}
}
