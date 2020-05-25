package resolver

import(
	"gitlab.com/kentanggoreng/quip-server/entity"
	"gitlab.com/kentanggoreng/quip-server/service"
	"gitlab.com/kentanggoreng/quip-server/repository"
)

type Resolver struct{}

var Schema = `
  	schema{
		query: Query
		mutation: Mutation
  	}
  	type Query{
		hello: String!
	}
	type Mutation{
		login(email: String!, password: String!): String!
		register(email: String!, username: String!, password: String!): String!
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
	account,_ := accountRepository.GetDataByIndex("email",args.Email)
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
	account := &entity.Account{
		ID: service.GenerateUUID(),
		Email: args.Email,
		Username: args.Username,
		Password: args.Password,
	}

	accountRepository := repository.NewAccountRepository()
	accountRepository.PutData(account)

	return service.GenerateJWT(account.ID)
}