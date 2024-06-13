package cacSchema

import (
	"github.com/graphql-go/graphql"
)

var PipelineType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Pipeline",
	Description: "Pipeline Pathname",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"pathname": &graphql.Field{
			Type: graphql.String,
		},
	},
},
)
