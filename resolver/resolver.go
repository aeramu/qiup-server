package resolver

import(
	"gitlab.com/kentanggoreng/quip-server/entity"
	"gitlab.com/kentanggoreng/quip-server/service"
	"gitlab.com/kentanggoreng/quip-server/repository"
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
	}
	type Mutation{
		login(email: String!, password: String!): String!
		register(email: String!, username: String!, password: String!): String!
	}
	type Account{
		id: ID!
		email: String!
		username: String!
		profile: Profile
	}
	type Profile{
		name: String
		bio: String
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
		return "email not registered"
	}
	if args.Password != account.Password {
		return "wrong password"
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
		return "email already registered"
	}
	if accountRepository.GetDataByIndex("username",args.Username) != nil {
		return "username already registered"
	}

	account := &entity.Account{
		ID: service.GenerateUUID(),
		Email: args.Email,
		Username: args.Username,
		Password: args.Password,
	}
	accountRepository.PutData(account)

	return service.GenerateJWT(account.ID)
}

func (r *Resolver) Account(args struct{
	ID graphql.ID
})(*AccountResolver){
	accountRepository := repository.NewAccountRepository()
	account := accountRepository.GetDataByIndex("id",string(args.ID))

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