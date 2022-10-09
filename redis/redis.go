package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

type RedisService struct{}

var RedisServiceImpl RedisService

func (r *RedisService) Init() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:55000",
		Password: "redispw", // no password set
		DB:       0,         // use default DB
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("redis can not start")
	}
	fmt.Println(pong)
	if pong != "PONG" {
		panic("redis can not link")
	}
	return rdb
}

func (r *RedisService) HGet() {
}
