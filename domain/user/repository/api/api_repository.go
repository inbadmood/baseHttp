package api

import (
	"baseHttp/entities"
	"fmt"
	"time"
)

type userAPIRepository struct {
	URL   string
	token string
}

func NewUserAPIRepository(uRL string, token string) entities.UserAPIRepository {
	return &userAPIRepository{
		URL:   uRL,
		token: token,
	}
}

func (_r *userAPIRepository) GetUserInfo() (newSerial string, err error) {
	// you can implement call api here

	// mock response
	newSerial = fmt.Sprint(time.Now().UnixNano())

	return
}
