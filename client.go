package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var RedisInstance = newClient()

type Redis struct {
	Client *redis.Client
	Ctx    context.Context
}

var host = *flag.String("redis.hostname", "127.0.0.1", "Hostname, eg 127.0.0.1")
var port = *flag.Int("redis.port", 6379, "Port redis is running on")

func newClient() Redis {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		DialTimeout:  1 & time.Second,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
		PoolFIFO:     false,
		PoolTimeout:  5 * time.Second,
		MinIdleConns: 1,
		MaxIdleConns: 5,
	})

	return Redis{Client: redisClient, Ctx: context.Background()}
}
