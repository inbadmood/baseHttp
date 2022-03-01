package api

import (
	"baseHttp/config"
	"github.com/go-redis/redis"
	"log"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"baseHttp/utils"

	_userDelivery "baseHttp/domain/user/delivery/http"
	_userApiRepository "baseHttp/domain/user/repository/api"
	_userMysqlRepository "baseHttp/domain/user/repository/mysql"
	_userRedisRepository "baseHttp/domain/user/repository/redis"
	_userUsecase "baseHttp/domain/user/usecase"
	_middleware "baseHttp/middleware"
)

var mysqlClient *gorm.DB
var redisClient *redis.Client
var timeoutContext = time.Duration(15) * time.Second

func initSetup() {
	// 初始化logging
	utils.InitLogging()
	// 綁定環境變數
	if err := config.InitialEnvConfiguration(); err != nil {
		utils.Logging.SendErrorLog(utils.GetFunctionName(), err.Error(), string(debug.Stack()))
		log.Fatal(err)
	}

	// 初始化database
	mysqlClient = config.NewMysqlConnect()
	redisClient = config.NewRedisConnection()
}

func StartAPIServer() {
	initSetup()

	// gin
	g := gin.New()
	g.GET("", func(c *gin.Context) {}) // default router

	// default middleware
	defaultMiddleware := _middleware.InitDefaultMiddleware()
	g.Use(defaultMiddleware.RecoverPanic(), defaultMiddleware.Logger(), defaultMiddleware.CORSMiddleware())

	apiRouter := g.Group("api")

	// user
	userApiRepository := _userApiRepository.NewUserAPIRepository(config.EnvConfig.API.Url, config.EnvConfig.API.Token)
	userDataRepository := _userMysqlRepository.NewUserDataRepository(mysqlClient)
	userRedisRepository := _userRedisRepository.NewUserRedisRepository(redisClient)
	userUsecase := _userUsecase.NewUserUsecase(userDataRepository, userApiRepository, userRedisRepository, timeoutContext)
	_userDelivery.NewUserHandler(apiRouter, userUsecase)

	if config.EnvConfig.Env == "local" {
		_ = g.Run("localhost:" + config.EnvConfig.Port)
	} else {
		_ = g.Run(":" + config.EnvConfig.Port)
	}
}
