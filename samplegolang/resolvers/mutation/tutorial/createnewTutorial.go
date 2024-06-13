package tutorail

import (
	"github.com/graphql-go/graphql"
)

type Tutorial struct {
	ID       int
	Title    string
	Author   string
	Comments []Comment
}

type Comment struct {
	Body string
}

func CreateTutorial(params graphql.ResolveParams) (interface{}, error) {
	//tutorials := data.Populate()
	tutorial := Tutorial{
		Title:  params.Args["title"].(string),
		Author: params.Args["author"].(string),
	}
	//tutorials = append(tutorials, tutorial)
	return tutorial, nil
}
