package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/common/response"
	"github.com/guothion/xuanyuan/pkg/global"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CustomRecovery() gin.HandlerFunc {
	logConfig := global.App.Config.Log
	return gin.RecoveryWithWriter(&lumberjack.Logger{
		Filename:   logConfig.RootDir + "/" + logConfig.Filename,
		MaxSize:    logConfig.Size,
		MaxAge:     logConfig.Age,
		MaxBackups: logConfig.Backups,
		Compress:   logConfig.Compress,
	}, response.ServerError)
}
