package resolver

import (
	"context"

	"github.com/aeramu/qiup-server/domain"
	"github.com/aeramu/qiup-server/repository"
	"github.com/aeramu/qiup-server/service"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SharePostResolver graphql
type SharePostResolver struct {
	post *domain.SharePost
}

//ID query
func (r *SharePostResolver) ID() graphql.ID {
	return graphql.ID(r.post.ID.Hex())
}

//Timestamp query
func (r *SharePostResolver) Timestamp() int32 {
	return int32(r.post.ID.Timestamp().Unix())
}

//Account query
func (r *SharePostResolver) Account() *ShareAccountResolver {
	shareAccountRepository := repository.NewShareAccountRepository()
	account := shareAccountRepository.GetDataByIndex("_id", r.post.AccountID)
	return &ShareAccountResolver{account}
}

//Body query
func (r *SharePostResolver) Body() string {
	return r.post.Body
}

// SharePost query
func (r *Resolver) SharePost(args struct {
	ID graphql.ID
}) *SharePostResolver {
	sharePostRepository := repository.NewSharePostRepository()
	id, _ := primitive.ObjectIDFromHex(string(args.ID))
	post := sharePostRepository.GetDataByIndex("_id", id)
	return &SharePostResolver{post}
}

// SharePostList query
func (r *Resolver) SharePostList() []*SharePostResolver {
	sharePostRepository := repository.NewSharePostRepository()
	postList := sharePostRepository.GetDataList()
	var sharePostList []*SharePostResolver
	for _, post := range postList {
		sharePostList = append(sharePostList, &SharePostResolver{post})
	}
	return sharePostList
}

// PostSharePost mutation
func (r *Resolver) PostSharePost(ctx context.Context, args struct {
	Body string
}) *SharePostResolver {
	token := ctx.Value("token").(string)
	accountID, _ := primitive.ObjectIDFromHex(service.DecodeJWT(token))
	post := &domain.SharePost{
		ID:        primitive.NewObjectID(),
		AccountID: accountID,
		Body:      args.Body,
	}
	sharePostRepository := repository.NewSharePostRepository()
	sharePostRepository.PutData(post)
	return &SharePostResolver{post}
}
