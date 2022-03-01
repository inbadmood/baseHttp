package usecase

import (
	"baseApiServer/process/notify"
	"strings"
)

type UseCase struct {
	notifyRepo notify.ApiRepository
}

func NewNotifyUseCase(t notify.ApiRepository) notify.UseCase {
	return &UseCase{
		notifyRepo: t,
	}
}

func (_u UseCase) SendMsg(msg []string) {
	if _u.notifyRepo.GetNotifyIsOpen() == false {
		return
	}

	MsgString := strings.Join(msg, "\n")
	_u.notifyRepo.SendMsg(MsgString)
}
