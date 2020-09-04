package usecase

import (
	"github.com/aeramu/qiup-server/entity"
)

type interactor struct {
	menfessRepo MenfessRepo
}

//MenfessRepo interface
type MenfessRepo interface {
	NewID() string
	GetPostByID(id string) entity.MenfessPost
	GetPostListByParentID(parentID string, first int, after string, ascSort bool) []entity.MenfessPost
	GetPostListByRoomIDs(roomIDs []string, first int, after string, ascSort bool) []entity.MenfessPost
	PutPost(name string, avatar string, body string, parentID string, repostID string, roomID string) entity.MenfessPost
	UpdateUpvoterIDs(postID string, accountID string, exist bool)
	UpdateDownvoterIDs(postID string, accountID string, exist bool)
	GetRoomList() []entity.MenfessRoom
	GetRoom(id string) entity.MenfessRoom
}
