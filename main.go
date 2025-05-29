package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"supabase-testProject1/cmd/api"
	"supabase-testProject1/internal/database"

	"github.com/joho/godotenv"
)

func main() {

	loadDotEnv()

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("[ERROR] Could not connect to PostgreSQL: %v", err)
	}

	fmt.Println("[INFO] Successfully connected to PostgreSQL")

	server := api.NewAPIServer("0.0.0.0:20384", db)
	if err := server.Run(); err != nil {
		log.Fatalf("[ERROR] Could not start the API Server: %v", err)
	}

	//startServer(db)
}

func loadDotEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("[ERROR] Error loading .env file")
	}

	fmt.Println("[INFO] Loaded .env file")
}

func startServer(db *sql.DB) {

	fmt.Println("[INFO] Starting the server")

	var userPosition string
	username := "testUser" // or pull this from somewhere

	// Run the query with a placeholder
	err := db.QueryRowContext(
		context.Background(),
		"SELECT position FROM \"Users\" WHERE username = $1",
		username,
	).Scan(&userPosition)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			fmt.Printf("No user found with username %q\n", username)
		default:
			log.Printf("Error executing query: %v\n", err)
		}
		return
	}

	// Success!
	fmt.Println("User position is:", userPosition)
}
