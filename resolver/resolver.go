package resolver

import(
	"context"
	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/service"
	"github.com/aeramu/qiup-server/repository"
	"github.com/graph-gophers/graphql-go"
)

type Resolver struct{}

var Schema = `
  	schema{
		query: Query
		mutation: Mutation
  	}
  	type Query{
		hello: String!
		account(id: ID!): Account
		me: Account!
	}
	type Mutation{
		login(email: String!, password: String!): String!
		register(email: String!, username: String!, password: String!): String!
		editProfile(name: String!, bio: String!, profilePhoto: String!, coverPhoto: String!): Account!
		uploadImage(directory: String!): String!
	}
	type Account{
		id: ID!
		email: String!
		username: String!
		profile: Profile!
	}
	type Profile{
		name: String!
		bio: String!
		profilePhoto: String!
		coverPhoto: String!
	}
`

func (r *Resolver) Hello()(string){
	return "Hello world!"
}

func (r *Resolver) Login(args struct{
	Email string
	Password string
})(string){
	accountRepository := repository.NewAccountRepository()
	account := accountRepository.GetDataByIndex("email",args.Email)
	if account == nil {
		account = accountRepository.GetDataByIndex("username",args.Email)
		if account == nil{
			return "Email or username not registered"
		}
	}
	if service.Hash(args.Password) != account.Password {
		return "Wrong password"
	}
	return service.GenerateJWT(account.ID)
}

func (r *Resolver) Register(args struct{
	Email string
	Username string
	Password string
})(string){
	accountRepository := repository.NewAccountRepository()
	if accountRepository.GetDataByIndex("email",args.Email) != nil {
		return "Email already registered"
	}
	if accountRepository.GetDataByIndex("username",args.Username) != nil {
		return "Username already registered"
	}
	account := &entity.Account{
		ID: service.GenerateUUID(),
		Email: args.Email,
		Username: args.Username,
		Password: service.Hash(args.Password),
	}
	accountRepository.PutData(account)
	return service.GenerateJWT(account.ID)
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

func (r *Resolver) Account(args struct{
	ID graphql.ID
})(*AccountResolver){
	accountRepository := repository.NewAccountRepository()
	account := accountRepository.GetDataByIndex("_id",string(args.ID))
	return &AccountResolver{account}
}

func (r *Resolver) Me(ctx context.Context)(*AccountResolver){
	token := ctx.Value("token").(string)
	accountRepository := repository.NewAccountRepository()
	account := accountRepository.GetDataByIndex("_id",service.DecodeJWT(token))
	return &AccountResolver{account}
}

type AccountResolver struct{
	account *entity.Account
}
func (r *AccountResolver) ID()(graphql.ID){
	return graphql.ID(r.account.ID)
}
func (r *AccountResolver) Email()(string){
	return r.account.Email
}
func (r *AccountResolver) Username()(string){
	return r.account.Username
}
func (r *AccountResolver) Profile()(*ProfileResolver){
	return &ProfileResolver{r.account.Profile}
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