package log

import (
	"baseApiServer/models"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type LoggerInterface interface {
	LogApiInfo(event, recvMsg, token, response, costTime string)
	LogInfo(event, status, playerName, msg, errMsg, currentTime string)
	LogError(event, status, playerName, msg, errMsg, currentTime string)
	LogFatal(event, status, playerName, msg, errMsg, currentTime string)
}

type Logger struct {
	hostName string
	fileName string
	logrus   *logrus.Logger
}
type Data struct {
	Token       string `json:"Token,omitempty"`
	RecvMsg     string `json:"RecvMsg,omitempty"`
	CostTime    string `json:"CostTime,omitempty"`
	Host        string `json:"Host"`
	Event       string `json:"Event"`
	Status      string `json:"Status,omitempty"`
	Player      string `json:"PlayerName,omitempty"`
	ErrMsg      string `json:"ErrMsg,omitempty"`
	Response    string `json:"Response,omitempty"`
	Msg         string `json:"Msg,omitempty"`
	CurrentTime string `json:"currentTime,omitempty"`
}

var logEnv bool
var apiURL string
var token string
var chatID int
var env string

// NewLogger 帶入UseCase名稱寫相對應文件檔
func NewLogger(hostName, folderName, fileName string) LoggerInterface {
	// serverConfig := utils.SetConfigPath()
	env = os.Getenv("PROJECT_ENV")
	ServerConfig := viper.New()
	ServerConfig.SetConfigName("config")
	ServerConfig.SetConfigType("yaml")
	ServerConfig.AddConfigPath("./config/app/" + env)
	ServerConfig.ReadInConfig()

	log := logrus.New()
	// 設定日誌json格式
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误） 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 設置日誌等級以上
	log.SetLevel(logrus.InfoLevel)

	logEnv = ServerConfig.GetBool("log.debugLog")
	apiURL = ServerConfig.GetString("telegram-api.apiURL")
	token = ServerConfig.GetString("telegram-api.robotToken")
	chatID = ServerConfig.GetInt("telegram-api.chatID")

	return &Logger{
		hostName: hostName,
		fileName: fileName,
		logrus:   log,
	}
}

func (l *Logger) LogApiInfo(event, recvMsg, token, response, costTime string) {
	logInfo := l.logrus.WithFields(logrus.Fields{
		"Host":     l.hostName,
		"Event":    event,
		"Token":    token,
		"recvMsg":  recvMsg,
		"costTime": costTime,
	})
	logInfo.Info(response)

	txt, _ := json.Marshal(Data{Host: l.hostName, Event: event, Token: token, RecvMsg: recvMsg, CostTime: costTime, Response: response})
	if logEnv {
		log.Println(string(txt))
	}
}

func (l *Logger) LogInfo(event string, status string, playerName string, msg string, errMsg string, currentTime string) {
	logInfo := l.logrus.WithFields(logrus.Fields{
		"Host":        l.hostName,
		"Event":       event,
		"Status":      status,
		"Player":      playerName,
		"ErrMsg":      errMsg,
		"CurrentTime": currentTime,
	})
	logInfo.Info(msg)
	if logEnv {
		json, _ := json.Marshal(Data{Host: l.hostName, Event: event, Status: status, Player: playerName, ErrMsg: errMsg, CurrentTime: currentTime, Msg: msg})
		fmt.Println(string(json))
	}
}

func (l *Logger) LogError(event string, status string, playerName string, msg string, errMsg string, currentTime string) {
	logError := l.logrus.WithFields(logrus.Fields{
		"Host":        l.hostName,
		"Event":       event,
		"Status":      status,
		"Player":      playerName,
		"ErrMsg":      errMsg,
		"CurrentTime": currentTime,
	})

	logError.Error(msg)

	txt, _ := json.Marshal(Data{Host: l.hostName, Event: event, Status: status, Player: playerName, ErrMsg: errMsg, CurrentTime: currentTime, Msg: msg})
	if logEnv {
		log.Println(string(txt))
	}

	go func() {
		l.sendTelgram(l.buildTelReq(event, errMsg, string(txt)))
	}()
}

func (l *Logger) LogFatal(event string, status string, playerName string, msg string, errMsg string, currentTime string) {
	logError := l.logrus.WithFields(logrus.Fields{
		"Host":        l.hostName,
		"Event":       event,
		"Status":      status,
		"Player":      playerName,
		"ErrMsg":      errMsg,
		"CurrentTime": currentTime,
	})

	logError.Fatal(msg)

	txt, _ := json.Marshal(Data{Host: l.hostName, Event: event, Status: status, Player: playerName, ErrMsg: errMsg, CurrentTime: currentTime, Msg: msg})
	if logEnv {
		log.Println(string(txt))
	}
}

func (l *Logger) sendTelgram(msg models.TelegramReq) {
	client := resty.New()
	urlRequest := fmt.Sprintf(apiURL, token, "sendMessage")

	ByteData, _ := json.Marshal(msg)
	result := models.TelegramResp{}
	client.
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(ByteData)).
		SetResult(&result).
		Post(urlRequest)
}

func (l *Logger) buildTelReq(event string, errMsg string, msg string) models.TelegramReq {
	txt := l.hostName + "\n" + env + "\n" + event + "\n" + errMsg + "\n" + msg
	data := models.TelegramReq{
		ChatID: chatID,
		Text:   txt,
	}
	return data
}
