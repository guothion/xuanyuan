package bootstrap

import (
	"fmt"
	"github.com/guothion/xuanyuan/pkg/global"
	"github.com/robfig/cron/v3"
	"time"
)

func InitializeCron() {
	global.App.Cron = cron.New(cron.WithSeconds())

	go func() {
		// 我们要执行的脚本
		global.App.Cron.AddFunc("0 0 2 * * *", func() {
			fmt.Println(time.Now())
		})
		global.App.Cron.Start()
		defer global.App.Cron.Stop()
		select {}
	}()
}
