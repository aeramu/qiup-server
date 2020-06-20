package resolver

import(
	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/repository"
	"github.com/graph-gophers/graphql-go"
)

type JustPostResolver struct{
	post *entity.JustPost
}
func (r *JustPostResolver) ID()(graphql.ID){
	return graphql.ID(r.post.ID.Hex())
}
func (r *JustPostResolver) Timestamp()(int32){
	return int32(r.post.ID.Timestamp().Unix())
}
func (r *JustPostResolver) Parent()(*JustPostResolver){
	justPostRepository := repository.NewJustPostRepository()
	post := justPostRepository.GetDataByIndex("_id", r.post.ParentID)
	if (post == nil) {
		return nil
	} else {
		return &JustPostResolver{post}
	}
}
func (r *JustPostResolver) Name()(string){
	return r.post.Name
}
func (r *JustPostResolver) Avatar()(string){
	return r.post.Avatar
}
func (r *JustPostResolver) Body()(string){
	return r.post.Body
}
func (r *JustPostResolver) Child(args struct{
	First int32
	After graphql.ID
})(*JustPostConnectionResolver){
	justPostRepository := repository.NewJustPostRepository()
	justPostList := justPostRepository.GetDataListByIndex("parentID", r.post.ID, args.First, entity.ID(string(args.After)))
	return &JustPostConnectionResolver{justPostList}
}

type JustPostConnectionResolver struct{
	justPostList []*entity.JustPost
}
func (r *JustPostConnectionResolver) Edges()([]*JustPostResolver){
	var justPostResolverList []*JustPostResolver
	for _,post := range(r.justPostList) {
		justPostResolverList = append(justPostResolverList, &JustPostResolver{post})
	}
	return justPostResolverList
}
func (r *JustPostConnectionResolver) PageInfo()(*PageInfoResolver){
	return &PageInfoResolver{r.justPostList}
}

func (r *Resolver) JustPost(args struct{
	ID graphql.ID
})(*JustPostResolver){
	justPostRepository := repository.NewJustPostRepository()
	post := justPostRepository.GetDataByIndex("_id", entity.ID(string(args.ID)))
	return &JustPostResolver{post}
}

func (r *Resolver) JustPostList(args struct{
	First int32
	After graphql.ID
})(*JustPostConnectionResolver){
	justPostRepository := repository.NewJustPostRepository()
	justPostList := justPostRepository.GetDataList(args.First, entity.ID(string(args.After)))
	return &JustPostConnectionResolver{justPostList}
}

func (r *Resolver) PostJustPost(args struct{
	Name string
	Avatar string
	Body string
	ParentID graphql.ID
})(*JustPostResolver){
	post := &entity.JustPost{
		ID: entity.NewID(),
		ParentID: entity.ID(string(args.ParentID)),
		Name: args.Name,
		Avatar: args.Avatar,
		Body: args.Body,
	}
	justPostRepository := repository.NewJustPostRepository()
	justPostRepository.PutData(post)
	return &JustPostResolver{post}
}