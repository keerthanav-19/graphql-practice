package author

import (
	"log"
	"samplegolang/data/postgresconnector"
	"time"

	"github.com/graphql-go/graphql"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}

func CreatePost(params graphql.ResolveParams) (interface{}, error) {
	title, _ := params.Args["title"].(string)
	content, _ := params.Args["content"].(string)
	authorId, _ := params.Args["author_id"].(int)
	createdAt := time.Now()

	var lastInsertId int
	dbPool, err := postgresconnector.PostgresConnector()

	err = dbPool.QueryRow("INSERT INTO posts(title, content, author_id, created_at) VALUES($1, $2, $3, $4) returning id;", title, content, authorId, createdAt).Scan(&lastInsertId)

	if err != nil {
		log.Println("error:", err)
	}

	newPost := &Post{
		ID:        lastInsertId,
		Title:     title,
		Content:   content,
		AuthorID:  authorId,
		CreatedAt: createdAt,
	}

	return newPost, nil
}
