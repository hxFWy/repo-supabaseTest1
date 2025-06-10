package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"sort"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

//go:embed seeds/*.sql
var seedFS embed.FS

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

		runSeeds(dsn)
	}

	if cmd == "down" {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	}

}

func runSeeds(dsn string) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("sql.Open:", err)
	}
	defer db.Close()

	entries, err := fs.ReadDir(seedFS, "seeds")
	if err != nil {
		log.Fatal("ReadDir seeds:", err)
	}

	// By mieć deterministyczną kolejność – sortujemy wg nazw plików
	var names []string
	for _, e := range entries {
		if !e.IsDir() && path.Ext(e.Name()) == ".sql" {
			names = append(names, e.Name())
		}
	}
	sort.Strings(names)

	ctx := context.Background()
	for _, fname := range names {
		sqlBytes, err := seedFS.ReadFile("seeds/" + fname)
		if err != nil {
			log.Fatalf("ReadFile %s: %v", fname, err)
		}
		query := string(sqlBytes)
		fmt.Printf("[INFO] Seeding %s …\n", fname)
		if _, err := db.ExecContext(ctx, query); err != nil {
			log.Fatalf("Exec seed %s: %v", fname, err)
		}
	}
	fmt.Println("[INFO] All seeds applied")
}
