package usecase

import (
	serverLog "baseApiServer/log"
	"baseApiServer/process/authenticate"
	"os"
)

type UseCase struct {
	authMysqlRepo authenticate.MysqlAuthRepository
	authRedisRepo authenticate.RedisAuthRepository
}

var logObj serverLog.LoggerInterface

func init() {
	logObj = serverLog.NewLogger(os.Getenv("HOSTNAME"), "auth", "auth")
}

// NewAuthUseCase 初始化
func NewAuthUseCase(mysql authenticate.MysqlAuthRepository, redis authenticate.RedisAuthRepository) authenticate.UseCase {
	return &UseCase{
		authMysqlRepo: mysql,
		authRedisRepo: redis,
	}
}
