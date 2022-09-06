package redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)


type RedisClient struct {
	Client *redis.Client
	Context context.Context
}

var once sync.Once
var Redis *RedisClient

func Init() (err error) {
	once.Do(func() {
		Redis, err = NewClient(
			fmt.Sprintf("%s:%d", viper.GetString("redis.host"),viper.GetInt("redis.port")) ,
			"",
			viper.GetInt("redis.db"),
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