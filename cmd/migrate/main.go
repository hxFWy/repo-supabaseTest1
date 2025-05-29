package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

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

	fmt.Println("[INFO] Successfully connected to PostgreSQL")

	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		dsn)

	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	}

}
