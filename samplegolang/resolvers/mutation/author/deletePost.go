package author

import (
	"log"
	"samplegolang/data/postgresconnector"

	"github.com/graphql-go/graphql"
)

func DeletePost(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(int)

	dbPool, err := postgresconnector.PostgresConnector()

	stmt, err := dbPool.Prepare("DELETE FROM posts WHERE id = $1")
	if err != nil {
		log.Println("error:", err)
	}

	_, err2 := stmt.Exec(id)
	if err2 != nil {
		log.Println("error:", err2)
	}

	newPost := &Post{
		ID: id,
	}

	return newPost, nil
}
