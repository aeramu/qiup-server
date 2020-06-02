package main

import (
    "log"
	"net/http"
	
    "github.com/friendsofgo/graphiql"
    "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/aeramu/qiup-server/resolver"
)

func main() {
    schema := graphql.MustParseSchema(resolver.Schema, &resolver.Resolver{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/query")
	if err != nil {
		panic(err)
	}
	http.Handle("/", graphiqlHandler)
    
    log.Println("Server ready at 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}