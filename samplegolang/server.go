package main

import (
	"fmt"
	"log"
	"net/http"
	"samplegolang/mutation"
	"samplegolang/queries"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queries.Query,
		Mutation: mutation.MutationType,
	},
)

func main() {

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// serve HTTP
	http.Handle("/graphql", h)
	fmt.Println("Server is running on port 9091")
	listenerr := http.ListenAndServe(":9091", nil)
	if listenerr != nil {
		log.Println("ListenAndServe: ", listenerr)
	}
}
