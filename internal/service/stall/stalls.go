package stall

import (
	"context"
	"github.com/guothion/xuanyuan/internal/api/common/request"
	"github.com/guothion/xuanyuan/internal/model"
	"github.com/guothion/xuanyuan/internal/repository"
)

type stallsService struct{}

func (s *stallsService) GetById(ctx context.Context, id uint) (stall *model.Stalls, err error) {
	stall, err = repository.StallsRepo.GetByID(ctx, id, "")
	return
}

func (s *stallsService) Create(ctx context.Context, stalls request.StallCreate) error {
	stall := &model.Stalls{
		UserId:      stalls.UserId,
		Name:        stalls.Name,
		Description: stalls.Description,
		Category:    stalls.Category,
		CoverImage:  stalls.CoverImage,
		Location:    stalls.Location,
		Status:      stalls.Status,
		Reason:      stalls.Reason,
		StartTime:   stalls.StartTime,
		EndTime:     stalls.EndTime,
	}
	err := repository.StallsRepo.Create(ctx, stall)
	return err
}

func (s *stallsService) GetPage(ctx context.Context, page int, pageSize int) (stalls []model.Stalls, total int64, err error) {
	// 可选过滤条件
	conditions := map[string]interface{}{}
	stalls, total, err = repository.StallsRepo.GetPage(ctx, page, pageSize, conditions, "start_time desc", nil)
	return
}
