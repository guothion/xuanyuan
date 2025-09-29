package jobs

import (
	"context"
	"log"
	"time"
)

type ExampleJob struct{}

func (j ExampleJob) Name() string { return "example-job" }
func (j ExampleJob) Run() {
	// 这里我们需要创建一个上下文对象，防止一直请求导致的内存崩溃
	var ctx = context.Background()
	context.WithTimeout(ctx, 5*time.Second)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	log.Printf("✅ [%s] 示例任务开始执行...", time.Now().Format("15:04:05"))
	// 模拟业务逻辑
	time.Sleep(100 * time.Millisecond)
	log.Printf("✅ [%s] 示例任务执行完成", time.Now().Format("15:04:05"))
	//return nil
}
