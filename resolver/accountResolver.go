package resolver

import(
	"context"
	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/service"
	"github.com/aeramu/qiup-server/repository"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountResolver struct{
	account *entity.Account
}
func (r *AccountResolver) ID()(graphql.ID){
	return graphql.ID(r.account.ID.Hex())
}
func (r *AccountResolver) Email()(string){
	return r.account.Email
}

func (r *Resolver) MyAccount(ctx context.Context)(*AccountResolver){
	token := ctx.Value("token").(string)
	accountRepository := repository.NewAccountRepository()
	id,_ := primitive.ObjectIDFromHex(service.DecodeJWT(token))
	account := accountRepository.GetDataByIndex("_id",id)
	return &AccountResolver{account}
}

func (r *Resolver) IsEmailAvailable(args struct{
	Email string
})(bool){
	accountRepository := repository.NewAccountRepository()
	account := accountRepository.GetDataByIndex("email",args.Email)
	if account == nil {
		return true
	} else{
		return false
	}
}

func (r *Resolver) Register(args struct{
	Email string
	Password string
})(string){
	accountRepository := repository.NewAccountRepository()
	if accountRepository.GetDataByIndex("email",args.Email) != nil {
		return "Email already registered"
	}
	account := &entity.Account{
		ID: primitive.NewObjectID(),
		Email: args.Email,
		Password: service.Hash(args.Password),
	}
	accountRepository.PutData(account)
	return service.GenerateJWT(account.ID.Hex())
}

func (r *Resolver) Login(args struct{
	Email string
	Password string
})(string){
	accountRepository := repository.NewAccountRepository()
	account := accountRepository.GetDataByIndex("email",args.Email)
	if account == nil {
		return "Email not registered"
	}
	if service.Hash(args.Password) != account.Password {
		return "Wrong password"
	}
	return service.GenerateJWT(account.ID.Hex())
}