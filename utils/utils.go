package utils

import (
	"baseApiServer/models"
	"strconv"
	"time"
)

const sqlTimeFormat = "2006-01-02 15:04:05"

func GetDayAgo(t time.Time, ago int) time.Time {
	return t.AddDate(0, 0, -1*ago)
}

// 取指定小數位
func Decimal(value float64, number int) float64 {
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', number, 64), 64)
	return value
}

// GetTimeToString get t trans 2000-01-01 01:01:01.000
func GetTimeToString(t time.Time) string {
	return t.Format(sqlTimeFormat)
}
func GetSQLTimeFormatToTime(dateTime string) time.Time {
	t, _ := time.Parse(sqlTimeFormat, dateTime)
	return t
}

func IntToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(int64(i), 10)
}

func FloatToString(i float64) string {
	return strconv.FormatFloat(i, 'f', 2, 64)
}

func ErrorMsg(code int) (resp string) {
	switch code {
	case models.ErrInputAgentID:
		resp = "ErrInputAgentID"
	case models.ErrInputWagerID:
		resp = "ErrInputWagerID"
	case models.ErrInputGameID:
		resp = "ErrInputGameID"
	case models.ErrInputPlayerID:
		resp = "ErrInputPlayerID"
	case models.ErrGetMysqlMainWager:
		resp = "ErrGetMysqlMainWager"
	case models.ErrGetMysqlMainWagerID:
		resp = "ErrGetMysqlMainWagerID"
	case models.ErrGetMysqlSubWager:
		resp = "ErrGetMysqlSubWager"
	case models.ErrGetMysqlSubWagerCount:
		resp = "ErrGetMysqlSubWagerCount"
	case models.ErrTimeLessFiveMin:
		resp = "ErrTimeLessFiveMin"
	default:
		resp = "Exception."
	}

	return
}
