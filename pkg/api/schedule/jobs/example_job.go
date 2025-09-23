package jobs

import (
	"log"
	"time"
)

type ExampleJob struct{}

func (j ExampleJob) Name() string { return "example-job" }
func (j ExampleJob) Run() {
	log.Printf("✅ [%s] 示例任务开始执行...", time.Now().Format("15:04:05"))
	// 模拟业务逻辑
	time.Sleep(100 * time.Millisecond)
	log.Printf("✅ [%s] 示例任务执行完成", time.Now().Format("15:04:05"))
	//return nil
}
