package cac_query

import (
	"log"
	"samplegolang/data/postgresconnector"

	"github.com/graphql-go/graphql"
)

type Pipeline struct {
	ID       int
	Pathname string
}

func PipelineList(p graphql.ResolveParams) (interface{}, error) {
	dbPool, err := postgresconnector.PostgresConnector()
	rows, err := dbPool.Query("SELECT id, pathname FROM cacpipeline;")

	if err != nil {
		log.Println("error:", err)
	}
	var pipelines []*Pipeline

	for rows.Next() {
		pipeline := &Pipeline{}
		err = rows.Scan(&pipeline.ID, &pipeline.Pathname)
		if err != nil {
			log.Println("error:", err)
		}
		pipelines = append(pipelines, pipeline)
	}
	defer dbPool.Close()

	return pipelines, nil
}
