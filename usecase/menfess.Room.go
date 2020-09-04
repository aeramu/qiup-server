package usecase

import "github.com/aeramu/qiup-server/entity"

func (i *interactor) MenfessRoomList() []entity.MenfessRoom {
	roomList := i.menfessRepo.GetRoomList()
	return roomList
}

func (i *interactor) MenfessRoom(id string) entity.MenfessRoom {
	room := i.menfessRepo.GetRoom(id)
	return room
}
