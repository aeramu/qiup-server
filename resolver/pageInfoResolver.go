package resolver

import (
	"github.com/graph-gophers/graphql-go"
)

type node interface {
	GetID() string
}

//PageInfoResolver graphql
type PageInfoResolver struct {
	nodeList []node
}

//StartCursor get startcursor
func (r *PageInfoResolver) StartCursor() *graphql.ID {
	if len(r.nodeList) == 0 {
		return nil
	}
	startCursor := graphql.ID(r.nodeList[0].GetID())
	return &startCursor
}

// EndCursor get endcursor
func (r *PageInfoResolver) EndCursor() *graphql.ID {
	if len(r.nodeList) == 0 {
		return nil
	}
	endCursor := graphql.ID(r.nodeList[len(r.nodeList)-1].GetID())
	return &endCursor
}
