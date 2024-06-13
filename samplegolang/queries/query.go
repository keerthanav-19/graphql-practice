package queries

import (
	"samplegolang/resolvers/query_resolvers/author_query"
	"samplegolang/resolvers/query_resolvers/cac_query"
	"samplegolang/resolvers/query_resolvers/tutorial_query"
	"samplegolang/schemas/authorSchema"
	"samplegolang/schemas/cacSchema"
	"samplegolang/schemas/tutorialschema"

	"github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getTutorialList": &graphql.Field{
				Type:        graphql.NewList(tutorialschema.TutorialType),
				Description: "Get Tutorial List",
				Resolve:     tutorial_query.TutorialList,
			},
			"tutorial": &graphql.Field{
				Type:        tutorialschema.TutorialType,
				Description: "Get Tutorial By ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: tutorial_query.TutorialID,
			},
			"author": &graphql.Field{
				Type:        authorSchema.AuthorType,
				Description: "Get an author",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: author_query.GetAuthor,
			},
			"authorList": &graphql.Field{
				Type:        graphql.NewList(authorSchema.AuthorType),
				Description: "List of authors",
				Resolve:     author_query.AuthorList,
			},
			"post": &graphql.Field{
				Type:        authorSchema.PostType,
				Description: "Get an Post",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: author_query.GetPost,
			},
			"postList": &graphql.Field{
				Type:        graphql.NewList(authorSchema.PostType),
				Description: "List of posts",
				Resolve:     author_query.PostList,
			},
			"pipelineList": &graphql.Field{
				Type:        graphql.NewList(cacSchema.PipelineType),
				Description: "List of pipelines",
				Resolve:     cac_query.PipelineList,
			},
		},
	},
)
