package redis

import (
	"baseHttp/entities"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type userRedisRepository struct {
	conn *redis.Client
}

func NewUserRedisRepository(conn *redis.Client) entities.UserRedisRepository {
	return &userRedisRepository{conn}
}

var userInfoKey = "userInfo:"

func (_r *userRedisRepository) GetUserInfo(ctx context.Context, userID int) (info string, err error) {
	key := userInfoKey + fmt.Sprint(userID)
	info, err = _r.conn.WithContext(ctx).Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return info, nil
		}
		return
	}

	return
}

func (_r *userRedisRepository) SetUserInfo(ctx context.Context, userInfo entities.User) (err error) {
	key := userInfoKey + fmt.Sprint(userInfo.ID)
	expireTime := time.Duration(600) * time.Second

	userInfoByte, err := json.Marshal(userInfo)
	if err != nil {
		return
	}
	_, err = _r.conn.WithContext(ctx).Set(key, string(userInfoByte), expireTime).Result()
	if err != nil {
		return
	}

	return
}

func (_r *userRedisRepository) DeleteUserInfo(ctx context.Context, userID int) (err error) {
	key := userInfoKey + fmt.Sprint(userID)
	_, err = _r.conn.WithContext(ctx).Del(key).Result()
	if err != nil {
		return
	}

	return
}
