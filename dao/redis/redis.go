package redis

import (
	"fmt"
	"web_app/settings"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host,
			cfg.Port),
		Password: cfg.Password, // 密码
		DB:       cfg.DB,       // 数据库
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	_, err = client.Ping().Result()
	return err
}

func Close() {
	_ = client.Close()
}
