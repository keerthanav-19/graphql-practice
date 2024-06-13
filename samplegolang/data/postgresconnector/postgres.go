package postgresconnector

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "user123"
	DB_NAME     = "cactesting"
	sslmode     = "disable"
)

type DB struct {
	*sql.DB
}

// func PostgresConnector(is_db bool) (*pgxpool.Pool, bool) {
// 	var databaseUrl string
// 	if is_db {
// 		databaseUrl = "postgres://" + DB_USER + ":" + DB_PASSWORD + "@localhost:5432/graqhqlDB"
// 	}
// 	fmt.Println("DB connection")
// 	conn, err := pgxpool.Connect(context.Background(), databaseUrl)
// 	//conn, err := sql.Open("postgres", databaseUrl)
// 	if err != nil {
// 		log.Printf("Unable to connect to database: %v\n", err)
// 	}
// 	return conn, true
// }

func PostgresConnector() (*DB, error) {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to db!")

	return &DB{db}, nil
}
