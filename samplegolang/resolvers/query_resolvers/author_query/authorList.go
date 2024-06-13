package author_query

import (
	"log"
	"samplegolang/data/postgresconnector"

	"github.com/graphql-go/graphql"
)

func AuthorList(p graphql.ResolveParams) (interface{}, error) {
	dbPool, err := postgresconnector.PostgresConnector()
	rows, err := dbPool.Query("SELECT id, name, email FROM authors;")

	if err != nil {
		log.Println("error:", err)
	}
	var authors []*Author

	for rows.Next() {
		author := &Author{}
		err = rows.Scan(&author.ID, &author.Name, &author.Email)
		if err != nil {
			log.Println("error:", err)
		}
		authors = append(authors, author)
	}
	defer dbPool.Close()

	return authors, nil
}
