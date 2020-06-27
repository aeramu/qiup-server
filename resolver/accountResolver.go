package resolver

import (
	"context"

	"github.com/aeramu/qiup-server/old/domain"
	"github.com/aeramu/qiup-server/old/repository"
	"github.com/aeramu/qiup-server/old/service"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AccountResolver graphql
type AccountResolver struct {
	account *domain.Account
}

//ID query
func (r *AccountResolver) ID() graphql.ID {
	return graphql.ID(r.account.ID.Hex())
}

//Email query
func (r *AccountResolver) Email() string {
	return r.account.Email
}

//MyAccount query
func (r *Resolver) MyAccount(ctx context.Context) *AccountResolver {
	token := ctx.Value("token").(string)
	accountRepository := repository.NewAccountRepository()
	id, _ := primitive.ObjectIDFromHex(service.DecodeJWT(token))
	account := accountRepository.GetDataByIndex("_id", id)
	return &AccountResolver{account}
}

//IsEmailAvailable query
func (r *Resolver) IsEmailAvailable(args struct {
	Email string
}) bool {
	accountRepository := repository.NewAccountRepository()
	account := accountRepository.GetDataByIndex("email", args.Email)
	if account == nil {
		return true
	}
	return false
}

//Register mutation
func (r *Resolver) Register(args struct {
	Email    string
	Password string
}) string {
	accountRepository := repository.NewAccountRepository()
	if accountRepository.GetDataByIndex("email", args.Email) != nil {
		return "Email already registered"
	}
	account := &domain.Account{
		ID:       primitive.NewObjectID(),
		Email:    args.Email,
		Password: service.Hash(args.Password),
	}
	accountRepository.PutData(account)
	return service.GenerateJWT(account.ID.Hex())
}

//Login mutation
func (r *Resolver) Login(args struct {
	Email    string
	Password string
}) string {
	accountRepository := repository.NewAccountRepository()
	account := accountRepository.GetDataByIndex("email", args.Email)
	if account == nil {
		return "Email not registered"
	}
	if service.Hash(args.Password) != account.Password {
		return "Wrong password"
	}
	return service.GenerateJWT(account.ID.Hex())
}
