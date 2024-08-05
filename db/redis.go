package db

import (
	"kodinggo/internal/config"

	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	rd := redis.NewClient(&redis.Options{
		Addr:     config.GetRedisHost(),
		Password: "",
		DB:       config.GetRedisDB(),
	})

	return rd
}
