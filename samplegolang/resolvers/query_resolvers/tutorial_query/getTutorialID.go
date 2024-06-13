package tutorial_query

import (
	"samplegolang/data"

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

func TutorialID(p graphql.ResolveParams) (interface{}, error) {
	tutorials := data.Populate()
	id, ok := p.Args["id"].(int)
	if ok {
		// Parse our tutorial array for the matching id
		for _, tutorial := range tutorials {
			if int(tutorial.ID) == id {
				// return our tutorial
				return tutorial, nil
			}
		}
	}
	return nil, nil
}
