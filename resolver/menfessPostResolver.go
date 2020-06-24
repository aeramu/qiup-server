package resolver

import(
	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/repository"
	"github.com/graph-gophers/graphql-go"
)

type MenfessPostResolver struct{
	post *entity.MenfessPost
}
func (r *MenfessPostResolver) ID()(graphql.ID){
	return graphql.ID(r.post.ID.Hex())
}
func (r *MenfessPostResolver) Timestamp()(int32){
	return int32(r.post.ID.Timestamp().Unix())
}
func (r *MenfessPostResolver) Name()(string){
	return r.post.Name
}
func (r *MenfessPostResolver) Avatar()(string){
	return r.post.Avatar
}
func (r *MenfessPostResolver) Body()(string){
	return r.post.Body
}
func (r *MenfessPostResolver) ReplyCount()(int32){
	return r.post.ReplyCount
}
func (r *MenfessPostResolver) UpvoteCount()(int32){
	return r.post.UpvoteCount
}
func (r *MenfessPostResolver) DownvoteCount()(int32){
	return r.post.DownvoteCount
}
func (r *MenfessPostResolver) Parent()(*MenfessPostResolver){
	menfessPostRepository := repository.NewMenfessPostRepository()
	post := menfessPostRepository.GetDataByIndex("_id", r.post.ParentID)
	if (post == nil) {
		return nil
	} else {
		return &MenfessPostResolver{post}
	}
}
func (r *MenfessPostResolver) Child(args struct{
	First *int32
	After *graphql.ID
	Before *graphql.ID
	Sort *int32
})(*MenfessPostConnectionResolver){
	first := int32(20)
	if (args.First != nil){
		first = *args.First
	}
	after := entity.ID("ffffffffffffffffffffffff")
	if (args.After != nil){
		after = entity.ID(string(*args.After))
	}
	menfessPostRepository := repository.NewMenfessPostRepository()
	menfessPostList := menfessPostRepository.GetDataListByIndex("parentID", r.post.ID, first, after)
	return &MenfessPostConnectionResolver{menfessPostList}
}

type MenfessPostConnectionResolver struct{
	menfessPostList []*entity.MenfessPost
}
func (r *MenfessPostConnectionResolver) Edges()([]*MenfessPostResolver){
	var menfessPostResolverList []*MenfessPostResolver
	for _,post := range(r.menfessPostList) {
		menfessPostResolverList = append(menfessPostResolverList, &MenfessPostResolver{post})
	}
	return menfessPostResolverList
}
func (r *MenfessPostConnectionResolver) PageInfo()(*PageInfoResolver){
	return &PageInfoResolver{r.menfessPostList}
}

func (r *Resolver) MenfessPost(args struct{
	ID graphql.ID
})(*MenfessPostResolver){
	menfessPostRepository := repository.NewMenfessPostRepository()
	post := menfessPostRepository.GetDataByIndex("_id", entity.ID(string(args.ID)))
	return &MenfessPostResolver{post}
}

func (r *Resolver) MenfessPostList(args struct{
	First *int32
	After *graphql.ID
	Before *graphql.ID
	Sort *int32
})(*MenfessPostConnectionResolver){
	first := int32(20)
	if (args.First != nil){
		first = *args.First
	}
	after := entity.ID("ffffffffffffffffffffffff")
	if (args.After != nil){
		after = entity.ID(string(*args.After))
	}
	menfessPostRepository := repository.NewMenfessPostRepository()
	menfessPostList := menfessPostRepository.GetDataListByIndex("parentID", entity.ID(""), first, after)
	return &MenfessPostConnectionResolver{menfessPostList}
}

func (r *Resolver) PostMenfessPost(args struct{
	Name string
	Avatar string
	Body string
	ParentID graphql.ID
})(*MenfessPostResolver){
	post := &entity.MenfessPost{
		ID: entity.NewID(),
		Name: args.Name,
		Avatar: args.Avatar,
		Body: args.Body,
		ReplyCount: 0,
		UpvoteCount: 0,
		DownvoteCount: 0,
		ParentID: entity.ID(string(args.ParentID)),
	}
	menfessPostRepository := repository.NewMenfessPostRepository()
	menfessPostRepository.PutData(post)
	return &MenfessPostResolver{post}
}