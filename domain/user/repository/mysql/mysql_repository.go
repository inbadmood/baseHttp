package mysql

import (
	_struct "baseHttp/domain/user/repository/mysql/struct"
	"baseHttp/entities"
	"context"
	"gorm.io/gorm"
)

type userMysqlRepository struct {
	conn *gorm.DB
}

func NewUserDataRepository(conn *gorm.DB) entities.UserDataRepository {
	return &userMysqlRepository{conn}
}

// you make access to outside source/tools here

func (_r *userMysqlRepository) UpdateUserSerial(ctx context.Context, userID int, newSerial string) (err error) {
	updateColumns := map[string]interface{}{
		"serial": newSerial,
	}

	err = _r.conn.WithContext(ctx).
		Table("User").
		Where("id = ?", userID).
		Model(&_struct.User{}).
		Updates(updateColumns).Error
	if err != nil {
		return
	}

	return
}

func (_r *userMysqlRepository) GetUserList(ctx context.Context) (userList []entities.User, err error) {
	var selectResult []_struct.User
	err = _r.conn.
		WithContext(ctx).
		Table("User").
		Find(&selectResult).
		Error
	if err != nil {
		return
	}

	userList = make([]entities.User, len(selectResult))
	for index, data := range selectResult {
		user := entities.User{
			ID:     data.ID,
			Name:   data.Name,
			Serial: data.Serial,
		}

		userList[index] = user
	}

	return
}

func (_r *userMysqlRepository) GetSingleUserInfo(ctx context.Context, userID int) (user entities.User, err error) {
	var selectResult _struct.User
	err = _r.conn.
		WithContext(ctx).
		Table("User").
		Where("id = ?", userID).
		Find(&selectResult).
		Error
	if err != nil {
		return
	}

	user = entities.User{
		ID:     selectResult.ID,
		Name:   selectResult.Name,
		Serial: selectResult.Serial,
	}

	return
}
