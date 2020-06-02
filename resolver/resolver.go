package resolver

import(
	"context"
	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/service"
	"github.com/aeramu/qiup-server/repository"
)

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

type ProfileResolver struct{
	profile *entity.Profile
}
func (r *ProfileResolver) Name()(string){
	if r.profile == nil{
		return ""
	}
	return r.profile.Name
}
func (r *ProfileResolver) Bio()(string){
	if r.profile == nil{
		return ""
	}
	return r.profile.Bio
}
func (r *ProfileResolver) ProfilePhoto()(string){
	if r.profile == nil{
		return ""
	}
	return r.profile.ProfilePhoto
}
func (r *ProfileResolver) CoverPhoto()(string){
	if r.profile == nil{
		return ""
	}
	return r.profile.CoverPhoto
}