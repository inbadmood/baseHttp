package middleware

import (
	"baseHttp/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type DefaultMiddleware struct {
}

func InitDefaultMiddleware() *DefaultMiddleware {
	return &DefaultMiddleware{}
}

func (m *DefaultMiddleware) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowOrigins := map[string]bool{
			"http://localhost:8080": true,
		}
		origin := c.Request.Header.Get("Origin")
		if allowOrigins[origin] {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	}
}

func (m *DefaultMiddleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 開始時間
		startTime := time.Now()
		// 處理請求
		c.Next()
		// 結束時間
		endTime := time.Now()
		// 執行時間
		latencyTime := endTime.Sub(startTime)
		// 請求方式
		reqMethod := c.Request.Method
		// 請求路由
		reqURI := c.Request.RequestURI
		// 狀態碼
		statusCode := c.Writer.Status()

		// 請求IP
		clientIP, exists := c.Get("clientIP")
		if !exists {
			clientIP = c.ClientIP() + "(gin)"
		}
		log := map[string]interface{}{
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIP":    clientIP,
			"reqMethod":   reqMethod,
			"reqURI":      reqURI,
			"startTime":   startTime,
		}
		logJSON, _ := json.Marshal(log)
		fmt.Printf(string(logJSON) + "\n")
	}
}

// RecoverPanic 回收panic
func (m *DefaultMiddleware) RecoverPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			errMessage := recover()
			if errMessage != nil {
				utils.Logging.SendErrorLogObj(utils.Logging.MakeLogging(utils.InterfaceToString(errMessage)))
				c.AbortWithStatusJSON(http.StatusOK, "")
				return
			}
		}()
		c.Next()
	}
}
