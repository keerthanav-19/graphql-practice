package authorSchema

import (
	"github.com/graphql-go/graphql"
)

var CreateAuthor = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"email": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var UpdateAuthor = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"email": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var DeleteAuthor = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}
