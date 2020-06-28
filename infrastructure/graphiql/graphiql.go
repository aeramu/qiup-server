package main

import (
	"log"
	"net/http"

	resolver "github.com/aeramu/qiup-server/implementation/graphql.resolver"
	cleanrepo "github.com/aeramu/qiup-server/implementation/mongodb.repository"
	"github.com/aeramu/qiup-server/usecase"
	"github.com/friendsofgo/graphiql"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	schema := graphql.MustParseSchema(resolver.Schema, &resolver.Resolver{
		Interactor: usecase.InteractorConstructor{
			MenfessPostRepo: cleanrepo.New(),
		}.New(),
	})
	http.Handle("/query", &relay.Handler{Schema: schema})

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/query")
	if err != nil {
		panic(err)
	}
	http.Handle("/", graphiqlHandler)

	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
