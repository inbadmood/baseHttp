package server

import (
	"baseApiServer/delivery/agentDelivery"
	"baseApiServer/delivery/gameDelivery"
	"baseApiServer/utils"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"

	_authRepo "baseApiServer/process/authenticate/repository"
	_authUseCase "baseApiServer/process/authenticate/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var serverConfig *viper.Viper

var authMysqlConn *gorm.DB
var authRedisConn *redis.Client

var serverPort int

// 初始化
func initRun() {
	serverConfig = utils.SetConfigPath()
	serverPort = serverConfig.GetInt("server.port")
	authMysqlConn = utils.NewMysql(serverConfig, "platform", "master")
	authRedisConn = utils.NewRedis(serverConfig, "player")
}

func Run() {
	initRun()
	// 載入router root
	r := SetupRouter()
	r.Run(":" + strconv.Itoa(serverPort))
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// cors
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// authenticate api
	authMysqlRepo := _authRepo.NewMysqlRepository(authMysqlConn)
	authRedisRepo := _authRepo.NewRedisRepository(authRedisConn)
	authUseCase := _authUseCase.NewAuthUseCase(authMysqlRepo, authRedisRepo)

	gameDelivery.NewGameProviderHandler(r, authUseCase, serverConfig)
	agentDelivery.NewAgentOperatorHandler(r, authUseCase, serverConfig)

	return r
}
