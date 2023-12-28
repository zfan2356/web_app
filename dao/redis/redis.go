package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	rdb *redis.Client
)

func InitRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", viper.GetString("redis.host"),
			viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		PoolSize: viper.GetInt("redis.pool_size"),
		DB:       viper.GetInt("redis.db"),
	})

	_, err = rdb.Ping().Result()
	return err
}

func Close() {
	err := rdb.Close()
	if err != nil {
		zap.L().Fatal("close redis failed: ", zap.Error(err))
	}
}
