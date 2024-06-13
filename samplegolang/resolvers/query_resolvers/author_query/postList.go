package author_query

import (
	"log"
	"samplegolang/data/postgresconnector"

	"github.com/graphql-go/graphql"
)

func PostList(p graphql.ResolveParams) (interface{}, error) {
	dbPool, err := postgresconnector.PostgresConnector()
	if err != nil {
		log.Println("error:", err)
	}
	rows, err := dbPool.Query("SELECT id, title, content, author_id FROM posts;")
	defer dbPool.Close()

	if err != nil {
		log.Println("error:", err)
	}

	var posts []*Post

	for rows.Next() {
		post := &Post{}
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID)
		if err != nil {
			log.Println("error:", err)
		}

		posts = append(posts, post)
	}

	return posts, nil
}
