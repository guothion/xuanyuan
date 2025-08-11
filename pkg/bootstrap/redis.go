package bootstrap

import (
	"context"
	"github.com/guothion/xuanyuan/pkg/global"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitializeRedis() *redis.Client {
	redisConf := global.App.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConf.Host + ":" + redisConf.Port,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logrus.Errorln("Redis connect ping failed,err:", err)
		return nil
	}
	return client
}
