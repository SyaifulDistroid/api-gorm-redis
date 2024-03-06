package redis

import (
	"context"
	"fmt"
	"service-journal/infrastructure/shared/constant"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Adress   string `json:"address"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

func LoadRedis(cfg RedisConfig) (client *redis.Client, err error) {

	ctx := context.Background()

	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Adress,
		Password: cfg.Password, // no password set
		DB:       cfg.Database, // use default DB
	})

	res := client.Ping(ctx)
	if res.Err() != nil {
		err = fmt.Errorf(constant.ErrConnectToRedis, res.Err().Error())
		return
	}

	if res.Val() != "PONG" {
		err = fmt.Errorf(constant.ErrConnectToRedis, constant.DISCONNECTED)
		return
	}

	return
}
