package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Role struct {
	ID          int
	Name        string
	Description string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Could not load .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not defined")
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("❌ Unable to connect to database: %v\n", err)
	}

	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("❌ DB is not responding: %v\n", err)
	}
	fmt.Println("✅ Connection successful! Database is alive.")
	fmt.Println("------------------------------------------------")

	rows, err := conn.Query(context.Background(), "SELECT id, name, description FROM auth.roles")
	if err != nil {
		log.Fatalf("❌ Query failed: %v\n", err)
	}
	defer rows.Close()

	fmt.Println("🔍 Searching for roles...")
	
	for rows.Next() {
		var r Role
		err := rows.Scan(&r.ID, &r.Name, &r.Description)
		if err != nil {
			log.Fatalf("Error reading row: %v", err)
		}
		fmt.Printf("👉 Role found: [ID: %d] %s - %s\n", r.ID, r.Name, r.Description)
	}
	
	if rows.Err() != nil {
		log.Fatalf("Error during iteration: %v", rows.Err())
	}
}