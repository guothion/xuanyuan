package bootstrap

import (
	"github.com/guothion/xuanyuan/pkg/common"
	"time"
)

var (
	ctx     = common.NewContext(time.Minute)
	Session = &sessionService{}
)
