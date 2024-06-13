package tutorialschema

import (
	"github.com/graphql-go/graphql"
)

var TutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TutorialType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(CommentType),
			},
		},
	},
)

var CommentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CommentType",
		Fields: graphql.Fields{
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
