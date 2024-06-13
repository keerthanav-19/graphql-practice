package tutorial_query

import (
	"samplegolang/data"

	"github.com/graphql-go/graphql"
)

func TutorialList(params graphql.ResolveParams) (interface{}, error) {
	tutorials := data.Populate()
	return tutorials, nil
}
