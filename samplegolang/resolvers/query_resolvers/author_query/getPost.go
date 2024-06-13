package author_query

import (
	"log"
	"samplegolang/data/postgresconnector"
	"time"

	"github.com/graphql-go/graphql"
)

type Post struct {
	ID        int
	Title     string
	Content   string
	AuthorID  int
	CreatedAt time.Time
}

func GetPost(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(int)

	post := &Post{}
	dbPool, err := postgresconnector.PostgresConnector()

	err = dbPool.QueryRow("select id, title, content, author_id from posts where id = $1;", id).Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID)

	if err != nil {
		log.Println("error:", err)
	}

	defer dbPool.Close()

	return post, nil
}
