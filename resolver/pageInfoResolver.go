package resolver

import(
	"github.com/aeramu/qiup-server/entity"
	"github.com/graph-gophers/graphql-go"
)

type PageInfoResolver struct{
	justPostList []*entity.JustPost
}
func (r *PageInfoResolver) StartCursor()(graphql.ID){
	startCursor := r.justPostList[0].ID.Hex()
	return graphql.ID(startCursor)
}
func (r *PageInfoResolver) EndCursor()(graphql.ID){
	endCursor := r.justPostList[len(r.justPostList)-1].ID.Hex()
	return graphql.ID(endCursor)
}