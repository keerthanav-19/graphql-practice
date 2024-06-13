package authorSchema

import (
	"github.com/graphql-go/graphql"
)

var AuthorType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Author",
	Description: "A Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
	},
},
)
