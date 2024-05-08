package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jordanmartinwebdev/pointmints/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("libsql", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbURL, err)
		os.Exit(1)
	}
	defer db.Close()
}
