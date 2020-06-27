package usecase

import (
	"github.com/aeramu/qiup-server/entity"
)

//Interactor interface
type Interactor interface {
	MenfessPost(id string) *entity.MenfessPost
	MenfessPostFeed(first int, after string) []*entity.MenfessPost
	MenfessPostChild(parentID string, first int, after string) []*entity.MenfessPost
	PostMenfessPost(name string, avatar string, body string, parentID string) *entity.MenfessPost
	UpvoteMenfessPost(accountID string, postID string) *entity.MenfessPost
	DownvoteMenfessPost(accountID string, postID string) *entity.MenfessPost
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
	Vote(postID string, accountID string, isUpvote bool)
	Unvote(postID string, accountID string, isUpvote bool)
}

func (i *interactor) MenfessPost(id string) *entity.MenfessPost {
	post := i.menfessPostRepo.GetDataByID(id)
	return post
}

func (i *interactor) MenfessPostFeed(first int, after string) []*entity.MenfessPost {
	postList := i.menfessPostRepo.GetDataListByParentID("", first, after, false)
	return postList
}

func (i *interactor) MenfessPostChild(parentID string, first int, after string) []*entity.MenfessPost {
	postList := i.menfessPostRepo.GetDataListByParentID(parentID, first, after, true)
	return postList
}

func (i *interactor) PostMenfessPost(name string, avatar string, body string, parentID string) *entity.MenfessPost {
	id, timestamp := i.menfessPostRepo.NewID()
	//TODO create in repo implementation
	post := &entity.MenfessPost{
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

func (i *interactor) UpvoteMenfessPost(accountID string, postID string) *entity.MenfessPost {
	post := i.menfessPostRepo.GetDataByID(postID)
	for index, id := range post.UpvoterIDs {
		if id == accountID {
			i.menfessPostRepo.Unvote(postID, accountID, true)
			post.UpvoteCount--
			post.UpvoterIDs[index] = post.UpvoterIDs[len(post.UpvoterIDs)-1]
			post.UpvoterIDs = post.UpvoterIDs[:len(post.UpvoterIDs)-1]
			return post
		}
	}
	i.menfessPostRepo.Vote(postID, accountID, true)
	post.UpvoteCount++
	post.UpvoterIDs = append(post.UpvoterIDs, accountID)
	return post
}

func (i *interactor) DownvoteMenfessPost(accountID string, postID string) *entity.MenfessPost {
	post := i.menfessPostRepo.GetDataByID(postID)
	for index, id := range post.DownvoterIDs {
		if id == accountID {
			i.menfessPostRepo.Unvote(postID, accountID, false)
			post.DownvoteCount--
			post.DownvoterIDs[index] = post.DownvoterIDs[len(post.DownvoterIDs)-1]
			post.DownvoterIDs = post.DownvoterIDs[:len(post.DownvoterIDs)-1]
			return post
		}
	}
	i.menfessPostRepo.Vote(postID, accountID, false)
	post.DownvoteCount++
	post.DownvoterIDs = append(post.DownvoterIDs, accountID)
	return post
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
