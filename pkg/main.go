package main

import (
	"github.com/guothion/xuanyuan/pkg/bootstrap"
	"github.com/guothion/xuanyuan/pkg/global"
)

func main() {
	// 初始化配置
	bootstrap.InitConfig()
	// 初始化日志配置
	bootstrap.InitLogger()
	// 初始化 DB
	global.App.DB = bootstrap.InitDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 初始化验证器
	bootstrap.InitializeValidator()
	// 初始化redis
	global.App.Redis = bootstrap.InitializeRedis()

	bootstrap.InitializeCron()

	// 运行服务器
	bootstrap.RunServer()
}
