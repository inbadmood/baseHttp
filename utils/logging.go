package utils

import (
	"baseHttp/config"
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"cloud.google.com/go/logging"
)

var Logging ISendLog

const (
	SeverityError    = "Error"
	SeverityWarning  = "Warning"
	SeverityCritical = "Critical"
	SeverityInfo     = "Info"
)

type logInterface struct {
	ISendLog   ISendLog
	severity   logging.Severity
	appName    string
	loggerType string
	projectID  string
}

type payload struct {
	App        string      `json:"app"`
	Function   string      `json:"function"`
	LogObj     interface{} `json:"logObj"`
	LogMessage interface{} `json:"message"`
}
type ErrorDebug struct {
	FunctionName string
	ErrorMsg     interface{}
	ErrorObj     interface{}
}
type ConsoleLogFormat struct {
	Severity     string
	FunctionName string
	ErrorMsg     interface{}
	AppName      string
	Log          interface{}
}
type LogFormat struct {
	Function         string
	LogObj           interface{}
	LogMessage       interface{}
	LogMessageAppend string
	Time             string
}

type ISendLog interface {
	SendLog(severity string, funcName string, logMsg, logObj interface{})
	SendSeverityLogObj(severity string, log LogFormat)
	SendLogWithMultiObjs(severity string, funcName string, logMsg string, logObj ...interface{})
	SendErrorLog(funcName string, logMsg, logObj interface{})
	SendErrorLogObj(log LogFormat)
	MakeLogging(logMsg string, logObj ...interface{}) (log LogFormat)
}

func InitLogging() {
	serviceName := os.Getenv("APP_NAME")
	if serviceName == "" {
		serviceName = "baseHttp"
	}
	Logging = &logInterface{
		appName:    serviceName,
		loggerType: "global",
		projectID:  "example",
	}
}
func (l *logInterface) SendErrorLog(funcName string, logMsg, logObj interface{}) {
	l.severity = logging.Error
	l.ConsoleLogging(funcName, logMsg, logObj)
}

func (l *logInterface) SendErrorLogObj(log LogFormat) {
	l.severity = logging.Error
	l.ConsoleLogging(log.Function, log.LogMessage, log.LogObj)
}

func (l *logInterface) SendLog(severity string, funcName string, logMsg, logObj interface{}) {
	l.severity = getLogSeverity(severity)
	l.ConsoleLogging(funcName, logMsg, logObj)
}

func (l *logInterface) SendSeverityLogObj(severity string, log LogFormat) {
	l.severity = getLogSeverity(severity)
	l.ConsoleLogging(log.Function, log.LogMessage, log.LogObj)
}

func (l *logInterface) SendLogWithMultiObjs(severity string, funcName string, logMsg string, logObj ...interface{}) {
	l.severity = getLogSeverity(severity)
	l.ConsoleLogging(funcName, logMsg, logObj)
}

func (l *logInterface) ConsoleLogging(funcName string, logMsg, logObj interface{}) {
	log := ConsoleLogFormat{
		Severity:     l.severity.String(),
		FunctionName: funcName,
		ErrorMsg:     logMsg,
		AppName:      l.appName,
		Log:          logObj,
	}
	if config.EnvConfig.Env == "local" {
		fmt.Println(log)
	} else {
		logJSON, _ := json.Marshal(log)
		fmt.Println(string(logJSON))
	}

}

func (l *logInterface) MakeLogging(logMsg string, logObj ...interface{}) (log LogFormat) {
	logObj = append(logObj, string(debug.Stack()))
	log = LogFormat{
		Function:   GetFunctionName(),
		LogObj:     logObj,
		LogMessage: logMsg,
		Time:       GetLogTimeFormat(time.Now()),
	}
	return
}

func getLogSeverity(level string) (severity logging.Severity) {
	switch level {
	case "Info":
		severity = logging.Info
	case "Error":
		severity = logging.Error
	case "Warning":
		severity = logging.Warning
	case "Critical":
		severity = logging.Critical
	default:
		severity = logging.Error
	}

	return
}
