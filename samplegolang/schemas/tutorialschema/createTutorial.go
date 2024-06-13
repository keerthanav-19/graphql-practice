package tutorialschema

import (
	"github.com/graphql-go/graphql"
)

var CreateTutorail = graphql.FieldConfigArgument{
	"author": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"title": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}
