package authorSchema

import (
	"samplegolang/resolvers/query_resolvers/author_query"

	"github.com/graphql-go/graphql"
)

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Post",
		Description: "A Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"author_id": &graphql.Field{
				Type:    AuthorType,
				Resolve: author_query.GetAuthorPost,
			},
		},
	},
)
