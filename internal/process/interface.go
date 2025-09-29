package process

import "github.com/guothion/xuanyuan/pkg/common"

type CustomLogicFunc func(ctx *common.Context, processContext *ProcessContext) error

type ProcessContext struct {
}
