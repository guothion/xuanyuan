package schedule

import "github.com/robfig/cron/v3"

type Job = cron.Job

// FuncJob 包装 func() error 为 Job 接口
type FuncJob struct {
	f    func()
	name string
}

func (f FuncJob) Run()         { f.f() }
func (f FuncJob) Name() string { return f.name }

// NewFuncJob 创建一个基于函数的 Job
func NewFuncJob(name string, f func()) Job {
	return FuncJob{name: name, f: f}
}
