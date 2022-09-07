package redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/sjxiang/bluebell/settings"
)


type RedisClient struct {
	Client *redis.Client
	Context context.Context
}

var once sync.Once
var Redis *RedisClient

func Init(cfg *settings.RedisConfig) (err error) {
	once.Do(func() {
		Redis, err = NewClient(
			fmt.Sprintf("%s:%d", cfg.Host, cfg.Port) ,
			cfg.Password,
			cfg.DB,
		)
	})
	return
}

func NewClient(address string, password string, db int) (*RedisClient, error) {
	rds := &RedisClient{}

	rds.Context = context.Background()
	rds.Client = redis.NewClient(&redis.Options{
		Addr: address,
		Password: password,
		DB: db,
	})

	// test
	if _, err := rds.Client.Ping(rds.Context).Result(); err != nil {
		return nil, err
	}

	return rds, nil
}

func Close() {
	_ = Redis.Client.Close()
}