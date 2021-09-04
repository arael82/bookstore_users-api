package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
)

func init() {
	//	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	//		username, password, host, port, database)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres", "postgres", "localhost", "5432", "coursepractice")

	Client, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("DB Successfully configured")
}
