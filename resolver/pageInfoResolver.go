package resolver

import(
	"github.com/aeramu/qiup-server/entity"
	"github.com/graph-gophers/graphql-go"
)

type PageInfoResolver struct{
	justPostList []*entity.JustPost
}
func (r *PageInfoResolver) StartCursor()(*graphql.ID){
	if (len(r.justPostList)==0){
		return nil
	}
	startCursor := graphql.ID(r.justPostList[0].ID.Hex())
	return &startCursor
}
func (r *PageInfoResolver) EndCursor()(*graphql.ID){
	if (len(r.justPostList)==0){
		return nil
	}
	endCursor := graphql.ID(r.justPostList[len(r.justPostList)-1].ID.Hex())
	return &endCursor
}