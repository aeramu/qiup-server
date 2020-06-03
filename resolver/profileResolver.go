package resolver

import(
	"context"
	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/service"
	"github.com/aeramu/qiup-server/repository"
	"github.com/graph-gophers/graphql-go"
)

type ProfileResolver struct{
	profile *entity.Profile
}
func (r *ProfileResolver) ID()(graphql.ID){
	return graphql.ID(r.profile.ID)
}
func (r *ProfileResolver) Username()(string){
	return r.profile.Username
}
func (r *ProfileResolver) Name()(string){
	return r.profile.Name
}
func (r *ProfileResolver) Bio()(string){
	return r.profile.Bio
}
func (r *ProfileResolver) ProfilePhoto()(string){
	return r.profile.ProfilePhoto
}
func (r *ProfileResolver) CoverPhoto()(string){
	return r.profile.CoverPhoto
}

func (r *Resolver) Profile(args struct{
	ID graphql.ID
})(*AccountResolver){
	accountRepository := repository.NewAccountRepository()
	account := accountRepository.GetDataByIndex("_id",string(args.ID))
	return &AccountResolver{account}
}

func (r *Resolver) MyProfile(ctx context.Context)(*AccountResolver){
	token := ctx.Value("token").(string)
	accountRepository := repository.NewAccountRepository()
	account := accountRepository.GetDataByIndex("_id",service.DecodeJWT(token))
	return &AccountResolver{account}
}

func (r *Resolver) IsUsernameAvailable(args struct{
	Username string
})(bool){
	accountRepository := repository.NewAccountRepository()
	account := accountRepository.GetDataByIndex("username",args.Username)
	if account == nil {
		return true
	} else{
		return false
	}
}

func (r *Resolver) EditProfile(ctx context.Context, args struct{
	Name string
	Bio string
	ProfilePhoto string
	CoverPhoto string
})(*AccountResolver){
	token := ctx.Value("token").(string)
	profile := &entity.Profile{
		Name: args.Name,
		Bio: args.Bio,
		ProfilePhoto: args.ProfilePhoto,
		CoverPhoto: args.CoverPhoto,
	}
	accountRepository := repository.NewAccountRepository()
	account := accountRepository.UpdateData(service.DecodeJWT(token),"profile",profile)
	return &AccountResolver{account}
}

func (r *Resolver) UploadImage(args struct{
	Directory string
})(string){
	directory := args.Directory + "/" + service.GenerateUUID() + ".jpg"
	s3Repository := repository.NewS3Repository()
	url := s3Repository.PutImage(directory)
	return url
}