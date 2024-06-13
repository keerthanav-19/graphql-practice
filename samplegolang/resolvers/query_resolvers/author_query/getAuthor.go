package author_query

import (
	"log"
	"samplegolang/data/postgresconnector"
	"time"

	"github.com/graphql-go/graphql"
)

type Author struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

func GetAuthor(params graphql.ResolveParams) (interface{}, error) {
	//var err error
	id, _ := params.Args["id"].(int)

	author := &Author{}
	dbPool, err := postgresconnector.PostgresConnector()

	err = dbPool.QueryRow("select id, name, email from authors where id = $1;", id).Scan(&author.ID, &author.Name, &author.Email)
	defer dbPool.Close()

	if err != nil {
		log.Println("error:", err)
	}
	return author, nil

}

func GetAuthorPost(p graphql.ResolveParams) (interface{}, error) {
	author := &Author{}
	if post, ok := p.Source.(*Post); ok {
		dbPool, err := postgresconnector.PostgresConnector()

		err = dbPool.QueryRow("select id, name, email from authors where id = $1", post.AuthorID).Scan(&author.ID, &author.Name, &author.Email)

		if err != nil {
			log.Println("error:", err)
		}

	}
	return author, nil
}
