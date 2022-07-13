package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/yeongbok77/TaskManager/settings"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		zap.L().Error("redis connect failed", zap.Error(err))
		return err
	}
	return nil
}
func Close() {
	err := rdb.Close()
	if err != nil {
		zap.L().Error("redis close failed", zap.Error(err))
		return
	}
}
