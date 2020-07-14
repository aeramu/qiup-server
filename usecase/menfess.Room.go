package usecase

import "github.com/aeramu/qiup-server/entity"

func (i *interactor) MenfessRoomList() []entity.MenfessRoom {
	post := i.menfessRepo.GetRoomList()
	return post
}
