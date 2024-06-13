package mutation

import (
	"samplegolang/resolvers/mutation/author"
	tutorail "samplegolang/resolvers/mutation/tutorial"
	"samplegolang/schemas/authorSchema"
	"samplegolang/schemas/tutorialschema"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createTutorial": &graphql.Field{
			Type:        tutorialschema.TutorialType,
			Description: "Create a new Tutorial",
			Args:        tutorialschema.CreateTutorail,
			Resolve:     tutorail.CreateTutorial,
		},
		"createAuthor": &graphql.Field{
			Type:        authorSchema.AuthorType,
			Description: "Create a new author",
			Args:        authorSchema.CreateAuthor,
			Resolve:     author.CreateAuthor,
		},
		"updateAuthor": &graphql.Field{
			Type:        authorSchema.AuthorType,
			Description: "Update an author",
			Args:        authorSchema.UpdateAuthor,
			Resolve:     author.UpdateAuthor,
		},
		"deleteAuthor": &graphql.Field{
			Type:        authorSchema.AuthorType,
			Description: "Delete an author",
			Args:        authorSchema.DeleteAuthor,
			Resolve:     author.DeleteAuthor,
		},
		"createPost": &graphql.Field{
			Type:        authorSchema.PostType,
			Description: "Create a post",
			Args:        authorSchema.CreatePost,
			Resolve:     author.CreatePost,
		},
		"updatePost": &graphql.Field{
			Type:        authorSchema.PostType,
			Description: "Update a post",
			Args:        authorSchema.UpdatePost,
			Resolve:     author.UpdatePost,
		},
		"deletePost": &graphql.Field{
			Type:        authorSchema.PostType,
			Description: "Delete a post",
			Args:        authorSchema.DeletePost,
			Resolve:     author.DeletePost,
		},
	},
})
