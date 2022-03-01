package entities

import "context"

type User struct {
	ID     int
	Name   string
	Serial string
}

type UserUsecase interface {
	GetUserList(ctx context.Context) (userList []User, err error)
	GetSingleUserInfo(ctx context.Context, userID int) (userInfo User, err error)
	MakeNewUserSerial(ctx context.Context, inputSerial string) (serial string, err error)
	UpdateUserSerial(ctx context.Context, userID int, newSerial string) (err error)
}

type UserDataRepository interface {
	GetUserList(ctx context.Context) (userList []User, err error)
	GetSingleUserInfo(ctx context.Context, userID int) (user User, err error)
	UpdateUserSerial(ctx context.Context, userID int, newSerial string) (err error)
}

type UserAPIRepository interface {
	GetUserInfo() (newSerial string, err error)
}

type UserRedisRepository interface {
	GetUserInfo(ctx context.Context, userID int) (info string, err error)
	SetUserInfo(ctx context.Context, userInfo User) (err error)
	DeleteUserInfo(ctx context.Context, userID int) (err error)
}
