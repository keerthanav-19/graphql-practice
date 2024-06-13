package author

import (
	"log"
	"samplegolang/data/postgresconnector"
	"time"

	"github.com/graphql-go/graphql"
)

type Author struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateAuthor(params graphql.ResolveParams) (interface{}, error) {
	//var err error
	name, _ := params.Args["name"].(string)
	email, _ := params.Args["email"].(string)
	createdAt := time.Now()

	var lastInsertId int

	dbPool, err := postgresconnector.PostgresConnector()

	err = dbPool.QueryRow("INSERT INTO authors(name, email, created_at) VALUES($1, $2, $3) returning id;", name, email, createdAt).Scan(&lastInsertId)

	if err != nil {
		log.Println("error:", err)
	}

	defer dbPool.Close()

	newAuthor := &Author{
		ID:        lastInsertId,
		Name:      name,
		Email:     email,
		CreatedAt: createdAt,
	}

	return newAuthor, nil
}
