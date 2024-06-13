package author

import (
	"log"
	"samplegolang/data/postgresconnector"

	"github.com/graphql-go/graphql"
)

func UpdateAuthor(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(int)
	name, _ := params.Args["name"].(string)
	email, _ := params.Args["email"].(string)

	dbPool, err := postgresconnector.PostgresConnector()

	stmt, err := dbPool.Prepare("UPDATE authors SET name = $1, email = $2 WHERE id = $3;")
	if err != nil {
		log.Println("error:", err)
	}

	_, err2 := stmt.Exec(name, email, id)
	if err2 != nil {
		log.Println("error:", err2)
	}

	defer dbPool.Close()

	newAuthor := &Author{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return newAuthor, nil
}
