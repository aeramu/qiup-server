package resolver

import (
	"context"

	"github.com/aeramu/qiup-server/domain"
	"github.com/aeramu/qiup-server/repository"
	"github.com/aeramu/qiup-server/service"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ShareAccountResolver graphql
type ShareAccountResolver struct {
	account *domain.ShareAccount
}

//ID query
func (r *ShareAccountResolver) ID() graphql.ID {
	return graphql.ID(r.account.ID.Hex())
}

//Username query
func (r *ShareAccountResolver) Username() string {
	return r.account.Username
}

//Name query
func (r *ShareAccountResolver) Name() string {
	return r.account.ShareProfile.Name
}

//Bio query
func (r *ShareAccountResolver) Bio() string {
	return r.account.ShareProfile.Bio
}

//ProfilePhoto query
func (r *ShareAccountResolver) ProfilePhoto() string {
	return r.account.ShareProfile.ProfilePhoto
}

//CoverPhoto query
func (r *ShareAccountResolver) CoverPhoto() string {
	return r.account.ShareProfile.CoverPhoto
}

//ShareAccount query
func (r *Resolver) ShareAccount(args struct {
	ID graphql.ID
}) *ShareAccountResolver {
	shareAccountRepository := repository.NewShareAccountRepository()
	id, _ := primitive.ObjectIDFromHex(string(args.ID))
	account := shareAccountRepository.GetDataByIndex("_id", id)
	return &ShareAccountResolver{account}
}

// MyShareAccount query
func (r *Resolver) MyShareAccount(ctx context.Context) *ShareAccountResolver {
	token := ctx.Value("token").(string)
	shareAccountRepository := repository.NewShareAccountRepository()
	id, _ := primitive.ObjectIDFromHex(service.DecodeJWT(token))
	account := shareAccountRepository.GetDataByIndex("_id", id)
	return &ShareAccountResolver{account}
}

// IsUsernameAvailable query
func (r *Resolver) IsUsernameAvailable(args struct {
	Username string
}) bool {
	shareAccountRepository := repository.NewShareAccountRepository()
	account := shareAccountRepository.GetDataByIndex("username", args.Username)
	if account == nil {
		return true
	}
	return false
}

// SetShareUsername mutation
func (r *Resolver) SetShareUsername(ctx context.Context, args struct {
	Username string
}) string {
	token := ctx.Value("token").(string)
	shareAccountRepository := repository.NewShareAccountRepository()
	if shareAccountRepository.GetDataByIndex("username", args.Username) != nil {
		return "Username already taken"
	}
	id, _ := primitive.ObjectIDFromHex(service.DecodeJWT(token))
	account := &domain.ShareAccount{
		ID:       id,
		Username: args.Username,
	}
	shareAccountRepository.PutData(account)
	return ""
}

// SetShareProfile mutation
func (r *Resolver) SetShareProfile(ctx context.Context, args struct {
	Name         string
	Bio          string
	ProfilePhoto string
	CoverPhoto   string
}) *ShareAccountResolver {
	token := ctx.Value("token").(string)
	profile := &domain.ShareProfile{
		Name:         args.Name,
		Bio:          args.Bio,
		ProfilePhoto: args.ProfilePhoto,
		CoverPhoto:   args.CoverPhoto,
	}
	shareAccountRepository := repository.NewShareAccountRepository()
	id, _ := primitive.ObjectIDFromHex(service.DecodeJWT(token))
	account := shareAccountRepository.UpdateData(id, "shareProfile", profile)
	return &ShareAccountResolver{account}
}

// UploadImage mutation
func (r *Resolver) UploadImage(args struct {
	Directory string
}) string {
	directory := args.Directory + "/" + service.GenerateUUID() + ".jpg"
	s3Repository := repository.NewS3Repository()
	url := s3Repository.PutImage(directory)
	return url
}
