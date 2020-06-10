package resolver

import(
	"context"
	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/repository"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SharePostResolver struct{
	post *entity.SharePost
}
func (r *SharePostResolver) ID()(graphql.ID){
	return graphql.ID(r.post.ID.Hex())
}
func (r *SharePostResolver) Timestamp()(int64){
	return r.post.ID.Timestamp().Unix()
}
func (r *SharePostResolver) AccountID()(*ShareAccountResolver){
	shareAccountRepository := repository.NewShareAccountRepository()
	account := shareAccountRepository.GetDataByIndex("_id",r.post.AccountID)
	return &ShareAccountResolver{account}
}
func (r *SharePostResolver) Body()(string){
	return r.post.Body
}

func (r *Resolver) postSharePost(ctx context.Context, args struct{
	Body string
})(*SharePostResolver){
	token := ctx.Value("token").(string)
	accountID,_ := primitive.ObjectIDFromHex(service.DecodeJWT(token))
	post := &entity.SharePost{
		ID: primitive.NewObjectID(),
		AccountID: accountID,
		Body: args.Body,
	}
	sharePostRepository := repository.NewSharePostRepository()
	sharePostRepository.PutData(post)
	return &SharePostResolver{post}
}