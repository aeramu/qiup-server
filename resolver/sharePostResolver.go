package resolver

import(
	"context"
	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/repository"
	"github.com/aeramu/qiup-server/service"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SharePostResolver struct{
	post *entity.SharePost
}
func (r *SharePostResolver) ID()(graphql.ID){
	return graphql.ID(r.post.ID.Hex())
}
func (r *SharePostResolver) Timestamp()(int32){
	return int32(r.post.ID.Timestamp().Unix())
}
func (r *SharePostResolver) Account()(*ShareAccountResolver){
	shareAccountRepository := repository.NewShareAccountRepository()
	account := shareAccountRepository.GetDataByIndex("_id",r.post.AccountID)
	return &ShareAccountResolver{account}
}
func (r *SharePostResolver) Body()(string){
	return r.post.Body
}

func (r *Resolver) SharePost(args struct{
	ID graphql.ID
})(*SharePostResolver){
	sharePostRepository := repository.NewSharePostRepository()
	id,_ := primitive.ObjectIDFromHex(string(args.ID))
	post := sharePostRepository.GetDataByIndex("_id",id)
	return &SharePostResolver{post}
}

func (r *Resolver) SharePostList()([]*SharePostResolver){
	sharePostRepository := repository.NewSharePostRepository()
	postList := sharePostRepository.GetDataList()
	var sharePostList []*SharePostResolver
	for _,post := range(postList) {
		sharePostList = append(sharePostList,&SharePostResolver{post})
	}
	return sharePostList
}

func (r *Resolver) PostSharePost(ctx context.Context, args struct{
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