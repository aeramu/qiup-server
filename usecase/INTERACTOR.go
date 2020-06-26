package usecase

import (
	"fmt"

	"github.com/aeramu/qiup-server/entity"
)

//Interactor interface
type Interactor interface {
	MenfessPost(id string) *entity.MenfessPost
	MenfessPostFeed(first int, after string) []*entity.MenfessPost
	MenfessPostChild(parentID string, first int, after string) []*entity.MenfessPost
	PostMenfessPost(name string, avatar string, body string, parentID string) *entity.MenfessPost
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
	NewID() (string, int)
	GetDataByID(id string) *entity.MenfessPost
	GetDataListByParentID(parentID string, first int, after string, ascSort bool) []*entity.MenfessPost
	PutData(post *entity.MenfessPost)
}

func (i *interactor) MenfessPost(id string) *entity.MenfessPost {
	post := i.menfessPostRepo.GetDataByID(id)
	return post
}

func (i *interactor) MenfessPostFeed(first int, after string) []*entity.MenfessPost {
	postList := i.menfessPostRepo.GetDataListByParentID("", first, after, false)
	fmt.Println(postList)
	return postList
}

func (i *interactor) MenfessPostChild(parentID string, first int, after string) []*entity.MenfessPost {
	postList := i.menfessPostRepo.GetDataListByParentID(parentID, first, after, true)
	return postList
}

func (i *interactor) PostMenfessPost(name string, avatar string, body string, parentID string) *entity.MenfessPost {
	id, timestamp := i.menfessPostRepo.NewID()
	post := &entity.MenfessPost{
		// TODO make id generator
		ID:         id,
		Timestamp:  timestamp,
		Name:       name,
		Avatar:     avatar,
		Body:       body,
		ReplyCount: 0,
		ParentID:   parentID,
	}
	i.menfessPostRepo.PutData(post)
	return post
}
