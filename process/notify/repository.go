package notify

type ApiRepository interface {
	GetNotifyIsOpen() bool
	SendMsg(string) error
}
