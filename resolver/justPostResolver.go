package resolver

import(
	//"context"
	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/repository"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	post := justPostRepository.GetDataByIndex("_id",r.post.ParentID.Hex())
	return &JustPostResolver{post}
}
func (r *JustPostResolver) Name()(string){
	return r.post.Name
}
func (r *JustPostResolver) Body()(string){
	return r.post.Body
}
func (r *JustPostResolver) Child()([]*JustPostResolver){
	justPostRepository := repository.NewJustPostRepository()
	postList := justPostRepository.GetDataListByIndex("parentID",r.post.ParentID.Hex())
	var justPostList []*JustPostResolver
	for _,post := range(postList) {
		justPostList = append(justPostList,&JustPostResolver{post})
	}
	return justPostList
}

func (r *Resolver) PostJustPost(args struct{
	Name string
	Body string
	ParentID graphql.ID
})(*JustPostResolver){
	var post *entity.JustPost
	if (string(args.ParentID) == ""){
		post = &entity.JustPost{
			ID: primitive.NewObjectID(),
			Name: args.Name,
			Body: args.Body,
		}
	} else{
		parentID,_ := primitive.ObjectIDFromHex(string(args.ParentID))
		post = &entity.JustPost{
			ID: primitive.NewObjectID(),
			ParentID: parentID,
			Name: args.Name,
			Body: args.Body,
		}
	}
	justPostRepository := repository.NewJustPostRepository()
	justPostRepository.PutData(post)
	return &JustPostResolver{post}
}