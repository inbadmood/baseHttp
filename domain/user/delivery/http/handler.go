package http

import (
	"baseHttp/entities"
	"baseHttp/entities/delivery"
	"baseHttp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userUsecase entities.UserUsecase
}

func NewUserHandler(rg *gin.RouterGroup, user entities.UserUsecase) {
	handler := &userHandler{
		userUsecase: user,
	}
	rg.GET("/user/list", handler.GetUserList)
	rg.GET("/user", handler.GetSingleUserInfo)
	rg.PUT("/update/user/serial", handler.UpdateUserSerial)
}

type updateUserSerialBody struct {
	ID     int    `json:"id"`
	Serial string `json:"serial"`
}

type getShowBannerQuery struct {
	UserID int `form:"userID"`
}

// you make business func flow here

func (_h *userHandler) GetUserList(c *gin.Context) {
	ctx := c.Request.Context()

	userList, err := _h.userUsecase.GetUserList(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, err.Error())
		return
	}

	var response delivery.HTTPResponse
	response.Result = 1
	response.Retrieve = userList

	c.JSON(http.StatusOK, response)
}

func (_h *userHandler) UpdateUserSerial(c *gin.Context) {
	ctx := c.Request.Context()

	// binding input
	var body updateUserSerialBody
	if bindErr := c.ShouldBindJSON(&body); bindErr != nil {
		utils.Logging.SendErrorLogObj(utils.Logging.MakeLogging(bindErr.Error(), c.Request))
		c.AbortWithStatusJSON(http.StatusOK, bindErr.Error())
		return
	}

	// get new serial
	newSerial, err := _h.userUsecase.MakeNewUserSerial(ctx, body.Serial)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, err.Error())
		return
	}

	// update to db
	err = _h.userUsecase.UpdateUserSerial(ctx, body.ID, newSerial)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, err.Error())
		return
	}

	var response delivery.HTTPResponse
	response.Result = 1

	c.JSON(http.StatusOK, response)
}

func (_h *userHandler) GetSingleUserInfo(c *gin.Context) {
	ctx := c.Request.Context()

	var query getShowBannerQuery
	if bindErr := c.ShouldBindQuery(&query); bindErr != nil {
		utils.Logging.SendErrorLogObj(utils.Logging.MakeLogging(bindErr.Error(), c.Request))
		c.AbortWithStatusJSON(http.StatusOK, bindErr.Error())
		return
	}

	user, err := _h.userUsecase.GetSingleUserInfo(ctx, query.UserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, err.Error())
		return
	}

	var response delivery.HTTPResponse
	response.Result = 1
	response.Retrieve = user

	c.JSON(http.StatusOK, response)
}
