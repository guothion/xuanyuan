package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/common/response"
	"github.com/guothion/xuanyuan/pkg/api/schedule"
	"github.com/sirupsen/logrus"
)

var sched *schedule.CronScheduler

func init() {
	c := &CronRoute{}
	routes[c.BasePath()] = c
}

// 注入调度器实例
func SetScheduler(s *schedule.CronScheduler) {
	logrus.Println("SetScheduler: ", s)
	sched = s
}

type CronRoute struct{}

func (cron *CronRoute) BasePath() string { return "/cron" }

func (cron *CronRoute) RegisterRouter(router *gin.RouterGroup) {
	cronGroup := router.Group(cron.BasePath())
	{
		cronGroup.GET("", func(c *gin.Context) {
			response.Success(c, gin.H{"jobs": "Gin + Cron 服务运行中"})
		})
		// 查看所有定时任务
		cronGroup.GET("/jobs", func(c *gin.Context) {
			if sched == nil {
				response.BusinessFail(c, "sched 未初始化")
				return
			}
			response.Success(c, gin.H{"jobs": sched.ListJobs()})
		})
	}
}
