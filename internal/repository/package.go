package repository

import (
	"gorm.io/gorm"
)

// UserRepo 需在 DB 初始化完成后再注入
var UserRepo *UserRepository
var ProcessRepo *ProcessRepository
var ProcessStepRepo *StepRepository
var ProcessLogRepo *ProcessLogRepository
var AdminLogRepo *AdminLogsRepository
var CommentsRepo *CommentsRepository
var FavoritesRepo *FavoritesRepository
var NotificationRepo *NotificationsRepository
var StallApplicationRepo *StallApplicationRepository
var StallCategoriesRepo *StallCategoriesRepository
var StallImagesRepo *StallImagesRepository
var StallsRepo *StallsRepository

// Init 初始化仓库实例，避免在 DB 尚未就绪时创建导致空指针
func Init(db *gorm.DB) {
	UserRepo = NewUserRepo(db)
	ProcessRepo = NewProcessRepository(db)
	ProcessStepRepo = NewStepRepository(db)
	ProcessLogRepo = NewProcessLogRepository(db)
	AdminLogRepo = NewAdminLogsRepository(db)
	CommentsRepo = NewCommentsRepository(db)
	FavoritesRepo = NewFavoritesRepository(db)
	NotificationRepo = NewNotificationsRepository(db)
	StallApplicationRepo = NewStallApplicationRepository(db)
	StallCategoriesRepo = NewStallCategoriesRepository(db)
	StallImagesRepo = NewStallImagesRepository(db)
	StallsRepo = NewStallsRepository(db)
}
