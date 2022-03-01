package delivery

// HTTPResponse 通用response格式
type HTTPResponse struct {
	Result    int         `json:"result"`
	Retrieve  interface{} `json:"ret,omitempty"`
	ErrorCode string      `json:"code,omitempty"`
	Message   string      `json:"msg,omitempty"`
}
