package models

type AuthenticateResponse struct {
	ErrorCode     int                `json:"errorCode"`
	ErrorMessage  string             `json:"errorMessage"`
	ExecutionTime string             `json:"executionTime"`
	Data          PlayerDataResponse `json:"data"`
}
type PlayerDataResponse struct {
	PlayerID int    `json:"playerID"`
	GameID   int    `json:"gameID"`
	GameName string `json:"gameName"`
	GameType int    `json:"gameType"`
}
