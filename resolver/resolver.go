package resolver

import(
	//"gitlab.com/kentanggoreng/quip-server/entity"
	"gitlab.com/kentanggoreng/quip-server/service"
	repository "gitlab.com/kentanggoreng/quip-server/repository"
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
	account,_ := accountRepository.GetDataByIndex("Email",args.Email)
	if account == nil {
		return "email not registered"
	}
	if args.Password != account.Email {
		return "wrong password"
	}
	return service.GenerateJWT(account.ID)
}