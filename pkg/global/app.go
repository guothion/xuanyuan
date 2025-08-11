package global

import (
	"github.com/guothion/xuanyuan/pkg/config"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	DB          *gorm.DB
	Redis       *redis.Client
	Cron        *cron.Cron
}

// 这里我们使用全局的变量来绑定
var App = &Application{}
