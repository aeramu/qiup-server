package usecase

import "github.com/aeramu/qiup-server/entity"

//Interactor interface
type Interactor interface {
	menfessPost
	menfessRoom
}

type menfessPost interface {
	MenfessPost(id string) entity.MenfessPost
	MenfessPostFeed(first int, after string) []entity.MenfessPost
	MenfessPostChild(parentID string, first int, after string) []entity.MenfessPost
	MenfessPostRooms(roomIDs []string, first int, after string) []entity.MenfessPost
	PostMenfessPost(name string, avatar string, body string, parentID string, repostID string, roomID string) entity.MenfessPost
	UpvoteMenfessPost(accountID string, postID string) entity.MenfessPost
	DownvoteMenfessPost(accountID string, postID string) entity.MenfessPost
}

type menfessRoom interface {
	//MenfessRoomList() []entity.MenfessRoom
}

//InteractorConstructor constructor
type InteractorConstructor struct {
	MenfessPostRepo MenfessPostRepo
}

//New Construct Interactor
func (i InteractorConstructor) New() Interactor {
	return &interactor{
		menfessPostRepo: i.MenfessPostRepo,
	}
}
