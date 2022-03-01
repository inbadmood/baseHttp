package usecase

import (
	"baseHttp/entities"
	"baseHttp/utils"
	"context"
	"encoding/json"
	"time"
)

type userUsecase struct {
	dataRepository  entities.UserDataRepository
	apiRepository   entities.UserAPIRepository
	redisRepository entities.UserRedisRepository
	contextTimeout  time.Duration
}

func NewUserUsecase(dataRepository entities.UserDataRepository, api entities.UserAPIRepository, redis entities.UserRedisRepository, contextTimeout time.Duration) entities.UserUsecase {
	return &userUsecase{
		dataRepository:  dataRepository,
		apiRepository:   api,
		redisRepository: redis,
		contextTimeout:  contextTimeout,
	}
}

// you make business logic here

func (_u *userUsecase) UpdateUserSerial(ctx context.Context, userID int, newSerial string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, _u.contextTimeout)
	defer cancel()

	newSerial = "1234" + newSerial

	err = _u.dataRepository.UpdateUserSerial(ctx, userID, newSerial)
	if err != nil {
		utils.Logging.SendLog(utils.SeverityError, utils.GetFunctionName(), "", err.Error())
		return
	}

	return
}

func (_u *userUsecase) GetUserList(ctx context.Context) (userList []entities.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, _u.contextTimeout)
	defer cancel()

	userList, err = _u.dataRepository.GetUserList(ctx)
	if err != nil {
		utils.Logging.SendLog(utils.SeverityError, utils.GetFunctionName(), "", err.Error())
		return
	}

	return
}

func (_u *userUsecase) MakeNewUserSerial(ctx context.Context, inputSerial string) (serial string, err error) {
	_, cancel := context.WithTimeout(ctx, _u.contextTimeout)
	defer cancel()

	apiSerial, err := _u.apiRepository.GetUserInfo()
	if err != nil {
		utils.Logging.SendLog(utils.SeverityError, utils.GetFunctionName(), "", err.Error())
		return
	}

	serial = apiSerial + "_" + inputSerial

	return
}

func (_u *userUsecase) GetSingleUserInfo(ctx context.Context, userID int) (userInfo entities.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, _u.contextTimeout)
	defer cancel()

	userInfoString, err := _u.redisRepository.GetUserInfo(ctx, userID)
	if err != nil {
		utils.Logging.SendLog(utils.SeverityError, utils.GetFunctionName(), "", err.Error())
		return
	}
	if userInfoString != "" {
		err = json.Unmarshal([]byte(userInfoString), &userInfo)
		if err != nil {
			utils.Logging.SendLog(utils.SeverityError, utils.GetFunctionName(), "", err.Error())
			return
		}
		return
	}

	// get mysql
	userInfo, err = _u.dataRepository.GetSingleUserInfo(ctx, userID)
	if err != nil {
		utils.Logging.SendLog(utils.SeverityError, utils.GetFunctionName(), "", err.Error())
		return
	}
	// set redis
	err = _u.redisRepository.SetUserInfo(ctx, userInfo)
	if err != nil {
		utils.Logging.SendLog(utils.SeverityInfo, utils.GetFunctionName(), "", err.Error())
		return
	}

	return
}
