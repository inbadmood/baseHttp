package agentDelivery

import (
	"baseApiServer/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (_h *AgentOperatorHandler) CreatePlayer(c *gin.Context) {
	createPlayerRequest := models.CreatePlayerRequest{}
	byteValue, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	err := json.Unmarshal(byteValue, &createPlayerRequest)
	if err != nil {
		_h.ResponseErrorMessage(c, http.StatusOK, "", models.ErrUnmarshalInputData, "ErrUnmarshalInputData")
		return
	}
}
