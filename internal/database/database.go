package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {

	user := os.Getenv("user")
	password := os.Getenv("password")
	host := os.Getenv("host")
	port := os.Getenv("port")
	dbname := os.Getenv("dbname")

	// Connection string (DSN)
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=require",
		user, password, host, port, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	return db, nil
}
