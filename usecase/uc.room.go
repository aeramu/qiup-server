package usecase

import "github.com/aeramu/qiup-server/entity"

func (i *interactor) RoomList() []entity.Room {
	roomList := i.repo.GetRoomList()
	return roomList
}

func (i *interactor) Room(id string) entity.Room {
	room := i.repo.GetRoom(id)
	return room
}
