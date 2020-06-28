package usecase

import (
	"time"

	"github.com/aeramu/qiup-server/entity"
)

//Interactor interface
type Interactor interface {
	MenfessPost(id string) entity.MenfessPost
	MenfessPostFeed(first int, after string) []entity.MenfessPost
	MenfessPostChild(parentID string, first int, after string) []entity.MenfessPost
	PostMenfessPost(name string, avatar string, body string, parentID string) entity.MenfessPost
	UpvoteMenfessPost(accountID string, postID string) entity.MenfessPost
	DownvoteMenfessPost(accountID string, postID string) entity.MenfessPost
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

type interactor struct {
	menfessPostRepo MenfessPostRepo
}

//MenfessPostRepo interface
type MenfessPostRepo interface {
	NewID() string
	GetDataByID(id string) entity.MenfessPost
	GetDataListByParentID(parentID string, first int, after string, ascSort bool) []entity.MenfessPost
	PutData(post entity.MenfessPost)
	UpdateUpvoterIDs(postID string, accountID string, exist bool)
	UpdateDownvoterIDs(postID string, accountID string, exist bool)
}

func (i *interactor) MenfessPost(id string) entity.MenfessPost {
	post := i.menfessPostRepo.GetDataByID(id)
	return post
}

func (i *interactor) MenfessPostFeed(first int, after string) []entity.MenfessPost {
	postList := i.menfessPostRepo.GetDataListByParentID("", first, after, false)
	return postList
}

func (i *interactor) MenfessPostChild(parentID string, first int, after string) []entity.MenfessPost {
	postList := i.menfessPostRepo.GetDataListByParentID(parentID, first, after, true)
	return postList
}

func (i *interactor) PostMenfessPost(name string, avatar string, body string, parentID string) entity.MenfessPost {
	id := i.menfessPostRepo.NewID()
	post := entity.MenfessPostConstructor{
		ID:        id,
		Timestamp: int(time.Now().Unix()),
		Name:      name,
		Avatar:    avatar,
		Body:      body,
		ParentID:  parentID,
	}.New()
	i.menfessPostRepo.PutData(post)
	return post
}

func (i *interactor) UpvoteMenfessPost(accountID string, postID string) entity.MenfessPost {
	post := i.menfessPostRepo.GetDataByID(postID)
	if post.IsDownvoted(accountID) {
		exist := post.Downvote(accountID)
		i.menfessPostRepo.UpdateDownvoterIDs(postID, accountID, exist)
	}
	exist := post.Upvote(accountID)
	i.menfessPostRepo.UpdateUpvoterIDs(postID, accountID, exist)
	return post
}

func (i *interactor) DownvoteMenfessPost(accountID string, postID string) entity.MenfessPost {
	post := i.menfessPostRepo.GetDataByID(postID)
	if post.IsUpvoted(accountID) {
		exist := post.Upvote(accountID)
		i.menfessPostRepo.UpdateUpvoterIDs(postID, accountID, exist)
	}
	exist := post.Downvote(accountID)
	i.menfessPostRepo.UpdateDownvoterIDs(postID, accountID, exist)
	return post
}
