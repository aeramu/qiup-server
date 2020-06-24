package resolver

import(
	"github.com/aeramu/qiup-server/entity"
	"github.com/graph-gophers/graphql-go"
)

type PageInfoResolver struct{
	menfessPostList []*entity.MenfessPost
}
func (r *PageInfoResolver) StartCursor()(*graphql.ID){
	if (len(r.menfessPostList) == 0){
		return nil
	}
	startCursor := graphql.ID(r.menfessPostList[0].ID.Hex())
	return &startCursor
}
func (r *PageInfoResolver) EndCursor()(*graphql.ID){
	if (len(r.menfessPostList) == 0){
		return nil
	}
	endCursor := graphql.ID(r.menfessPostList[len(r.menfessPostList) - 1].ID.Hex())
	return &endCursor
}