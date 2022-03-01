package models

// AuthenticateRequest 傳入參數
type AuthenticateRequest struct {
	AgentID  string `json:"agentID"`
	PlayerID string `json:"playerID"`
	GameID   string `json:"gameID"`
}
