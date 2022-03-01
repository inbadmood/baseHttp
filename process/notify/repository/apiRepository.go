package repository

import (
	serverLog "baseApiServer/log"
	"baseApiServer/models"
	"baseApiServer/process/notify"
	"baseApiServer/utils"
	"time"

	"os"

	"crypto/tls"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type ApiRepository struct {
	env        string
	client     *resty.Client
	notifyOpen bool
	apiURL     string
	chatID     int
	token      string
}

var logObj serverLog.LoggerInterface

func init() {
	logObj = serverLog.NewLogger(os.Getenv("HOSTNAME"), "notify", "notify")
}

func NewTeleRepository(env string, c *resty.Client, viper *viper.Viper) notify.ApiRepository {

	return &ApiRepository{
		env:        env,
		client:     c,
		notifyOpen: viper.GetBool("notifyOpen"),
		apiURL:     viper.GetString("apiURL"),
		chatID:     viper.GetInt("chatID"),
		token:      viper.GetString("robotToken"),
	}
}

func (_a ApiRepository) GetNotifyIsOpen() bool {
	return _a.notifyOpen
}

func (_a ApiRepository) SendMsg(msg string) error {
	urlrequest := fmt.Sprintf(_a.apiURL, _a.token, "sendMessage")
	title := "環境 : " + _a.env + "\n"
	msg = title + msg

	msgRune := []rune(msg)
	sliceMsgLen := 800
	if len(msgRune) >= 800 {
		sliceMsgLen = 800
	} else {
		sliceMsgLen = len(msgRune)
	}
	data := models.TelegramReq{
		ChatID: _a.chatID,
		Text:   string(msgRune[:sliceMsgLen]),
	}
	ByteData, _ := json.Marshal(data)

	result := models.TelegramResp{}

	apiResponse, err := _a.client.
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(ByteData)).
		SetResult(&result).
		Post(urlrequest)

	json, _ := json.Marshal(models.LogNotify{
		UrlRequest: urlrequest,
		Response:   apiResponse.String(),
	})

	logObj.LogInfo("SendMsg", "Info", "", string(json), "", utils.GetTimeToString(time.Now()))

	if err != nil {
		logObj.LogError("SendMsg", "Error", "", "API Response Error", err.Error(), utils.GetTimeToString(time.Now()))
		return err
	}

	return nil
}
