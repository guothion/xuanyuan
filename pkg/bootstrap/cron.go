package bootstrap

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/guothion/xuanyuan/pkg/api/schedule"
	"github.com/guothion/xuanyuan/pkg/routes"
)

var sched *schedule.CronScheduler

func InitializeCron() {
	go func() {
		// 创建调度器
		sched = schedule.NewCronScheduler()

		// 注册任务
		registerJobs()

		// 注入调度器到路由（用于 API 查询）
		routes.SetScheduler(sched)

		// 启动调度器
		sched.Start()

		// 优雅关闭
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Println("Shutdown signal received...")

		// 停止 Cron
		sched.Stop()
	}()
}

func registerJobs() {
	// === 方式一：添加 Job（结构体）===
	//sched.AddJob("example", "@every 10s", &jobs.ExampleJob{})
	//sched.AddJob("cleanup", "30 */1 * * * *", &jobs.ExampleJob{})

	// === 方式二：添加 Func（函数）===
	//sched.AddFunc("heartbeat", "@every 30s", func() {
	//	fmt.Printf("💓 心跳任务执行: %s\n", time.Now().Format(time.DateTime))
	//})
}
