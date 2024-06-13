package author

import (
	"log"
	"samplegolang/data/postgresconnector"

	"github.com/graphql-go/graphql"
)

func UpdatePost(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(int)
	title, _ := params.Args["title"].(string)
	content, _ := params.Args["content"].(string)
	authorId, _ := params.Args["author_id"].(int)

	dbPool, err := postgresconnector.PostgresConnector()

	stmt, err := dbPool.Prepare("UPDATE posts SET title = $1, content = $2, author_id = $3 WHERE id = $4")
	if err != nil {
		log.Println("error:", err)
	}

	_, err2 := stmt.Exec(title, content, authorId, id)
	if err2 != nil {
		log.Println("error:", err2)
	}

	newPost := &Post{
		ID:       id,
		Title:    title,
		Content:  content,
		AuthorID: authorId,
	}

	return newPost, nil
}
