package repository

import (
	serverLog "baseApiServer/log"
	"baseApiServer/process/authenticate"
	"github.com/jinzhu/gorm"
	"os"
)

type mysqlAuthRepository struct {
	Conn *gorm.DB
}

var logObj serverLog.LoggerInterface

func init() {
	logObj = serverLog.NewLogger(os.Getenv("HOSTNAME"), "authRepo", "authRepo")
}

// NewMysqlRepository will create an object that represent the article.Repository interface
func NewMysqlRepository(conn *gorm.DB) authenticate.MysqlAuthRepository {
	return &mysqlAuthRepository{
		Conn: conn,
	}
}
