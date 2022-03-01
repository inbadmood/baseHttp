package models

import "errors"

var (
	ErrAuthInputData = errors.New("authInputDataError")
)

const (
	ErrInputAgentID = iota + 14000000 //
	ErrInputWagerID
	ErrInputGameID
	ErrInputPlayerID
	ErrUnmarshalInputData
	ErrTimeLessFiveMin
	ErrGetMysqlMainWager
	ErrGetMysqlMainWagerID
	ErrGetMysqlSubWager
	ErrGetMysqlSubWagerCount
	ErrGetNonCryptRes
	ErrGetEncryptRes
)

type ErrorOutputData struct {
	Code         int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
