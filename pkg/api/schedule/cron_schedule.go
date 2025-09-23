package schedule

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"sync"
)

type CronScheduler struct {
	cron    *cron.Cron
	jobs    map[string]Job          // 存储已注册的 Job
	entries map[string]cron.EntryID // 任务名 → EntryID
	specs   map[string]string       // 任务名 → cron 表达式
	mu      sync.RWMutex
}

func NewCronScheduler() *CronScheduler {
	return &CronScheduler{
		cron:    cron.New(cron.WithSeconds()),
		jobs:    make(map[string]Job),
		entries: make(map[string]cron.EntryID),
		specs:   make(map[string]string),
	}
}

// AddJob 添加实现了 Job 接口的任务
func (s *CronScheduler) AddJob(name, spec string, job Job) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if name == "" {
		logrus.Printf("cron job name is empty")
		return nil
	}

	if spec == "" {
		logrus.Printf("❌ 任务 %s 添加失败：cron 表达式为空", name)
		return nil
	}
	// 移除已有任务
	s.RemoveJob(name)
	entryID, err := s.cron.AddJob(spec, job)
	if err != nil {
		logrus.Printf("❌ 添加 Job 失败: %s (%s)", name, err)
		return err
	}

	s.jobs[name] = job
	s.entries[name] = entryID
	s.specs[name] = spec

	logrus.Printf("✅ Job 已添加: %s | Cron: %s", name, spec)
	return nil
}

// AddFunc 添加函数形式的任务（包装为 Job）
func (s *CronScheduler) AddFunc(name, spec string, f func()) error {
	job := NewFuncJob(name, f)
	return s.AddJob(name, spec, job)
}

// RemoveJob 移除任务
func (s *CronScheduler) RemoveJob(name string) {
	if id, exists := s.entries[name]; exists {
		s.cron.Remove(id)
		delete(s.entries, name)
		delete(s.jobs, name)
		delete(s.specs, name)
		logrus.Printf("🗑️ 任务已移除: %s", name)
	}
}

// ListJobs 列出所有任务（调试用）
func (s *CronScheduler) ListJobs() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make(map[string]string, len(s.specs))
	for name, spec := range s.specs {
		list[name] = spec
	}
	return list
}

// Start 启动调度器
func (s *CronScheduler) Start() {
	s.cron.Start()
	logrus.Printf("🚀 Cron 调度器已启动")
}

// Stop 停止调度器（优雅）
func (s *CronScheduler) Stop() {
	s.cron.Stop()
	logrus.Printf("🛑 Cron 调度器已停止")
}
