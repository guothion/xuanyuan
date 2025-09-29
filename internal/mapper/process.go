package mapper

import (
	"context"

	"github.com/guothion/xuanyuan/internal/global"
	"github.com/guothion/xuanyuan/internal/model"
)

type ProcessMapper struct{}

func (m *ProcessMapper) CreateRow(ctx context.Context, process *model.Process) (err error) {
	err = global.App.DB.WithContext(ctx).Create(&process).Error
	return
}

//func (m *ProcessMapper) GetByID(ctx context.Context, id uint) (process *model.Process, err error) {
//	err := global.App.DB.WithContext(ctx).
//		First(process, id).Error
//
//	return
//}

type StepMapper struct{}

func (m *StepMapper) CreateRow(ctx context.Context, step *model.Step) (err error) {
	err = global.App.DB.WithContext(ctx).Create(step).Error
	return
}

func (m *StepMapper) Get(ctx context.Context, step *model.Step) (err error) {
	result := global.App.DB.WithContext(ctx).First(step)
	return result.Error
}

//func (m *StepMapper) GetRows(ctx context.Context, processCode model.ProcessCode) (err error) {
//	result := global.App.DB.WithContext(ctx).Where("process_code=?", processCode).Find(&model.Process{})
//}

type StepLogMapper struct{}

func (m *StepLogMapper) CreateRow(ctx context.Context, processLog *model.ProcessLogs) (err error) {
	err = global.App.DB.WithContext(ctx).Create(processLog).Error
	return
}

func (m *StepLogMapper) Get(ctx context.Context, stepLog *model.ProcessLogs) (err error) {
	result := global.App.DB.WithContext(ctx).First(stepLog)
	return result.Error
}
