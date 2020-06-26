package usecase

import "github.com/aeramu/qiup-server/entity"

//Interactor interface
type Interactor interface {
	MenfessPost(id string) *entity.MenfessPost
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
	GetDataByID(id string) *entity.MenfessPost
}

func (i *interactor) MenfessPost(id string) *entity.MenfessPost {
	post := i.menfessPostRepo.GetDataByID(id)
	return post
}
