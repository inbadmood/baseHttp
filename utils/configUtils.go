package utils

import (
	serverLog "baseApiServer/log"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
	"time"
)

var logObj serverLog.LoggerInterface

func init() {
	logObj = serverLog.NewLogger(os.Getenv("HOSTNAME"), "utils", "utils")
}

func SetConfigPath() (ServerConfig *viper.Viper) {
	env := os.Getenv("ENV")
	ServerConfig = viper.New()
	ServerConfig.SetConfigName("config")
	ServerConfig.SetConfigType("yaml")
	ServerConfig.AddConfigPath("./config/app/" + env)
	err := ServerConfig.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err.Error())
	}
	return
}

// 建立redis
func NewRedis(ServerConfig *viper.Viper, name string) *redis.Client {
	maxActive := ServerConfig.GetInt("redis.maxActive")
	minIdle := ServerConfig.GetInt("redis.minIdle")
	host := ServerConfig.GetString("redis." + name + ".master.host")
	port := ServerConfig.GetString("redis." + name + ".master.port")
	IdleTimeout := ServerConfig.GetDuration("redis." + name + ".master.IdleTimeout")
	index, _ := strconv.Atoi(ServerConfig.GetString("redis." + name + ".master.index"))
	password := ServerConfig.GetString("redis." + name + ".master.password")
	// connectTimeout := ServerConfig.GetDuration("redis.ConnectTimeout")
	readTimeout := ServerConfig.GetDuration("redis.ReadTimeout")
	// writeTimeout := ServerConfig.GetDuration("redis.WriteTimeout")

	Conn := redis.NewClient(&redis.Options{
		Addr:         host + ":" + port,
		Password:     password,
		DB:           index,
		PoolSize:     maxActive, // Redis连接池大小
		MinIdleConns: minIdle,
		MaxRetries:   20, // 最大重试次数
		ReadTimeout:  readTimeout * time.Millisecond,
		IdleTimeout:  IdleTimeout * time.Second, // 空闲链接超时时间
	})

	_, err := Conn.Ping(context.Background()).Result()
	if err == redis.Nil {
		panic("Cannot connect Redis.")
	} else if err != nil {
		logObj.LogFatal("Redis Connect", "Fail", "", "Not able to establish connection to redis", "", GetTimeToString(time.Now()))
		panic(err)
	}

	return Conn
}

// 建立mysql
func NewMysql(ServerConfig *viper.Viper, name string, dbType string) *gorm.DB {
	dbName := ServerConfig.GetString("mysql." + name + "." + dbType + ".dbName")
	host := ServerConfig.GetString("mysql." + name + "." + dbType + ".host")
	port := ServerConfig.GetString("mysql." + name + "." + dbType + ".port")
	user := ServerConfig.GetString("mysql." + name + "." + dbType + ".user")
	password := ServerConfig.GetString("mysql." + name + "." + dbType + ".password")
	idleConns := ServerConfig.GetInt("mysql.idleConns")
	openConns := ServerConfig.GetInt("mysql.openConns")
	logOpen := ServerConfig.GetBool("mysql.logOpen")

	connectName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, dbName)

	retryCount := 10
	Conn, err := gorm.Open("mysql", connectName)
	if err != nil {
		for {
			if err != nil {
				if retryCount == 0 {
					logObj.LogFatal("MySQL Connect", "Fail", "", "Not able to connect to database", "", GetTimeToString(time.Now()))
				}
				log.Println(err)
				retryCount--
				time.Sleep(3 * time.Second)
				Conn, err = gorm.Open("mysql", connectName)
			} else {
				break
			}
		}
	}

	Conn.LogMode(logOpen)
	Conn.DB().SetMaxIdleConns(idleConns)
	Conn.DB().SetMaxOpenConns(openConns)
	Conn.DB().SetConnMaxLifetime(10)

	return Conn
}
