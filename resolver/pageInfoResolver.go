package resolver

import (
	"github.com/aeramu/qiup-server/entity"
	"github.com/graph-gophers/graphql-go"
)

//PageInfoResolver graphql
type PageInfoResolver struct {
	menfessPostList []entity.Node
}

//StartCursor get startcursor
func (r *PageInfoResolver) StartCursor() *graphql.ID {
	if len(r.menfessPostList) == 0 {
		return nil
	}
	startCursor := graphql.ID(r.menfessPostList[0].GetID().Hex())
	return &startCursor
}

// EndCursor get endcursor
func (r *PageInfoResolver) EndCursor() *graphql.ID {
	if len(r.menfessPostList) == 0 {
		return nil
	}
	endCursor := graphql.ID(r.menfessPostList[len(r.menfessPostList)-1].GetID().Hex())
	return &endCursor
}
