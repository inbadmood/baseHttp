package models

type TelegramResp struct {
	Ok          bool   `json:"ok"`
	Error_code  int    `json:"error_code"`
	Description string `json:"description"`
}

type TelegramReq struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

type LogNotify struct {
	UrlRequest string `json:"urlRequest"`
	Response   string `json:"response"`
}
