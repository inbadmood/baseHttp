package agentDelivery

import (
	"baseApiServer/log"
	"baseApiServer/models"
	"baseApiServer/process/authenticate"
	"baseApiServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

type AgentOperatorHandler struct {
	WagerUseCase authenticate.UseCase
	serverConfig *viper.Viper
}

var logObj log.LoggerInterface
var messageDecode bool

// NewAgentOperatorHandler router
func NewAgentOperatorHandler(r *gin.Engine, a authenticate.UseCase, config *viper.Viper) {
	logObj = log.NewLogger(os.Getenv("HOSTNAME"), "agentDelivery", "agentDelivery")

	handler := &AgentOperatorHandler{
		WagerUseCase: a,
		serverConfig: config,
	}
	messageDecode = config.GetBool("server.messageDecode")
	// 創建玩家
	r.POST("/agent/createPlayer", handler.CreatePlayer)
	r.GET("/agent/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}

// ResponseErrorMessage 取得回傳前端的錯誤訊息
func (_h *AgentOperatorHandler) ResponseErrorMessage(c *gin.Context, httpStatus int, encryptKey string, Code int, Data string) {
	if messageDecode {
		errOutPut := models.ErrorOutputData{
			Code:         Code,
			ErrorMessage: Data,
		}
		response, err := utils.MakeResponseEncryption("agentDelivery", encryptKey, errOutPut, false)
		if err != nil {
			c.JSON(httpStatus, gin.H{"errorCode": Code,
				"errorMessage": Data})
		}
		c.String(httpStatus, response)
	} else {
		c.JSON(httpStatus, gin.H{"errorCode": Code,
			"errorMessage": Data})
	}
}
