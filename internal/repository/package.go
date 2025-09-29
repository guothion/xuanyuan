package repository

import (
	"gorm.io/gorm"
)

// UserRepo 需在 DB 初始化完成后再注入
var UserRepo *UserRepository
var ProcessRepo *ProcessRepository
var ProcessStepRepo *StepRepository
var ProcessLogRepo *ProcessLogRepository

// Init 初始化仓库实例，避免在 DB 尚未就绪时创建导致空指针
func Init(db *gorm.DB) {
	UserRepo = NewUserRepo(db)
	ProcessRepo = NewProcessRepository(db)
	ProcessStepRepo = NewStepRepository(db)
	ProcessLogRepo = NewProcessLogRepository(db)
}
