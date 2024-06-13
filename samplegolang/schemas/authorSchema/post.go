package authorSchema

import (
	"github.com/graphql-go/graphql"
)

var CreatePost = graphql.FieldConfigArgument{
	"title": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"content": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"author_id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}

var UpdatePost = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"title": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"content": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"author_id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}

var DeletePost = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}
