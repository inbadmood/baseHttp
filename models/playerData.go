package models

type PlayerWallet struct {
	Balance     float64 `json:"balance"`     // 玩家錢包
	CreditRound string  `json:"creditRound"` // 交易唯一碼
}
