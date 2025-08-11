package bootstrap

import (
	"fmt"
	"github.com/guothion/xuanyuan/pkg/global"
	"github.com/guothion/xuanyuan/pkg/util"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"runtime"
	"strings"
)

var levelMap = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"panic": logrus.PanicLevel,
	"trace": logrus.TraceLevel,
	"fatal": logrus.FatalLevel,
}

func InitLogger() {
	cfg := global.App.Config.Log

	// 创建根目录
	createRootDir()

	rotateLogger := &lumberjack.Logger{
		Filename:   cfg.RootDir + "/" + cfg.Filename,
		MaxSize:    cfg.Size,
		MaxBackups: cfg.Backups,
		LocalTime:  true,
		MaxAge:     cfg.Age,
		Compress:   cfg.Compress,
	}
	// MultiWriter 允许你将数据同时写入多个 io.Writer
	ioWriter := io.MultiWriter(rotateLogger, os.Stdout)
	logrus.SetOutput(ioWriter)
	// 设置上报，会增加文件名等信息
	logrus.SetReportCaller(true)
	// 自定义 formmater
	logrus.SetFormatter(setFormatter())
	// 设置等级
	logrus.SetLevel(getLogLevel(cfg.Level))
	return
}

func createRootDir() {
	if ok, _ := util.PathExists(global.App.Config.Log.RootDir); !ok {
		_ = os.Mkdir(global.App.Config.Log.RootDir, os.ModePerm)
	}
}

// 这个最终返回 logrus 的相关对象
func setFormatter() logrus.Formatter {
	// logrus[ textFormatter | JSONFormatter ]
	formatter := &logrus.JSONFormatter{
		DisableTimestamp: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			name := frame.Function
			if names := strings.SplitN(frame.Function, "/", 2); len(names) > 2 {
				name = names[1]
			}
			return fmt.Sprintf("%s:%d", name, frame.Line), ""
		},
	}
	return formatter
}

func getLogLevel(conf string) logrus.Level {
	if level, ok := levelMap[conf]; ok {
		return level
	}
	// default is infolevel
	return logrus.InfoLevel
}
