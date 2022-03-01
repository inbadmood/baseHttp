package models

// createPlayer 傳入參數
type CreatePlayerRequest struct {
	AgentID  string `json:"agentID"`
	PlayerID string `json:"playerID"`
	Currency string `json:"currency"`
}
