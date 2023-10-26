package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"

	"github.com/youngdev-stack/minio-ffmpeg-lambda/global"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GlobalConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GlobalLog.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GlobalLog.Info("redis connect ping response:", zap.String("pong", pong))
		global.GlobalRedis = client
	}
}
