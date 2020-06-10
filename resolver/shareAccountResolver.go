package resolver

import(
	"context"
	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/service"
	"github.com/aeramu/qiup-server/repository"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShareAccountResolver struct{
	account *entity.ShareAccount
}
func (r *ShareAccountResolver) ID()(graphql.ID){
	return graphql.ID(r.account.ID.Hex())
}
func (r *ShareAccountResolver) Username()(string){
	return r.account.Username
}
func (r *ShareAccountResolver) Name()(string){
	return r.account.ShareProfile.Name
}
func (r *ShareAccountResolver) Bio()(string){
	return r.account.ShareProfile.Bio
}
func (r *ShareAccountResolver) ProfilePhoto()(string){
	return r.account.ShareProfile.ProfilePhoto
}
func (r *ShareAccountResolver) CoverPhoto()(string){
	return r.account.ShareProfile.CoverPhoto
}

func (r *Resolver) ShareAccount(args struct{
	ID graphql.ID
})(*ShareAccountResolver){
	shareAccountRepository := repository.NewShareAccountRepository()
	id,_ := primitive.ObjectIDFromHex(string(args.ID))
	account := shareAccountRepository.GetDataByIndex("_id",id)
	return &ShareAccountResolver{account}
}

func (r *Resolver) MyShareAccount(ctx context.Context)(*ShareAccountResolver){
	token := ctx.Value("token").(string)
	shareAccountRepository := repository.NewShareAccountRepository()
	id,_ := primitive.ObjectIDFromHex(service.DecodeJWT(token))
	account := shareAccountRepository.GetDataByIndex("_id",id)
	return &ShareAccountResolver{account}
}

func (r *Resolver) IsUsernameAvailable(args struct{
	Username string
})(bool){
	shareAccountRepository := repository.NewShareAccountRepository()
	account := shareAccountRepository.GetDataByIndex("username",args.Username)
	if account == nil {
		return true
	} else{
		return false
	}
}

func (r *Resolver) SetShareUsername(ctx context.Context, args struct{
	Username string
})(string){
	token := ctx.Value("token").(string)
	shareAccountRepository := repository.NewShareAccountRepository()
	if shareAccountRepository.GetDataByIndex("username",args.Username) != nil {
		return "Username already taken"
	}
	id,_ := primitive.ObjectIDFromHex(service.DecodeJWT(token))
	account := &entity.ShareAccount{
		ID: id,
		Username: args.Username,
	}
	shareAccountRepository.PutData(account)
	return ""
}

func (r *Resolver) SetShareProfile(ctx context.Context, args struct{
	Name string
	Bio string
	ProfilePhoto string
	CoverPhoto string
})(*ShareAccountResolver){
	token := ctx.Value("token").(string)
	profile := &entity.ShareProfile{
		Name: args.Name,
		Bio: args.Bio,
		ProfilePhoto: args.ProfilePhoto,
		CoverPhoto: args.CoverPhoto,
	}
	shareAccountRepository := repository.NewShareAccountRepository()
	id,_ := primitive.ObjectIDFromHex(service.DecodeJWT(token))
	account := shareAccountRepository.UpdateData(id,"shareProfile", profile)
	return &ShareAccountResolver{account}
}

func (r *Resolver) UploadImage(args struct{
	Directory string
})(string){
	directory := args.Directory + "/" + service.GenerateUUID() + ".jpg"
	s3Repository := repository.NewS3Repository()
	url := s3Repository.PutImage(directory)
	return url
}