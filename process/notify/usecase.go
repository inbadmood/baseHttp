package notify

type UseCase interface {
	SendMsg(msg []string)
}
