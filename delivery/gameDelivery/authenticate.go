package gameDelivery

import (
	"baseApiServer/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (_h *GameProviderHandler) Authenticate(c *gin.Context) {
	authenticateRequest := models.AuthenticateRequest{}
	byteValue, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	err := json.Unmarshal(byteValue, &authenticateRequest)
	if err != nil {
		_h.ResponseErrorMessage(c, http.StatusOK, "", models.ErrUnmarshalInputData, "ErrUnmarshalInputData")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
