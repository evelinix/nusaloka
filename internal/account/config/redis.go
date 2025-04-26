package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var RedisClient *redis.Client

func InitRedis() {
	addr := fmt.Sprintf("%s:%s", AccountConfig.RedisHost, AccountConfig.RedisPort)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: AccountConfig.RedisPass,
		DB:       0,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("[redis] Failed to connect to Redis")
	}

	log.Info().Msg("[redis] Redis connected successfully")
}
