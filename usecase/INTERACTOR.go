package usecase

import (
	"github.com/aeramu/qiup-server/entity"
)

type interactor struct {
	menfessPostRepo MenfessPostRepo
	menfessRoomRepo MenfessRoomRepo
}

//MenfessPostRepo interface
type MenfessPostRepo interface {
	NewID() string
	GetDataByID(id string) entity.MenfessPost
	GetDataListByParentID(parentID string, first int, after string, ascSort bool) []entity.MenfessPost
	GetDataListByRoomIDs(roomIDs []string, first int, after string, ascSort bool) []entity.MenfessPost
	PutData(name string, avatar string, body string, parentID string, repostID string, roomID string) entity.MenfessPost
	UpdateUpvoterIDs(postID string, accountID string, exist bool)
	UpdateDownvoterIDs(postID string, accountID string, exist bool)
}

//MenfessRoomRepo interface
type MenfessRoomRepo interface {
	GetDataList() []entity.MenfessRoom
}
