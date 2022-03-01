package repository

import (
	"baseApiServer/process/authenticate"

	"context"

	"github.com/go-redis/redis/v8"
)

type redisAuthRepository struct {
	Conn    *redis.Client
	context context.Context
}

// NewRedisRepository 建一個連線
func NewRedisRepository(conn *redis.Client) authenticate.RedisAuthRepository {
	wagerContext := context.Background()
	return &redisAuthRepository{
		Conn:    conn,
		context: wagerContext,
	}
}
