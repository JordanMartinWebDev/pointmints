package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jordanmartinwebdev/pointmints/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("warning: assuming default configuration. .env unreadable: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	apiCfg := apiConfig{}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("DATABASE_URL environment variable is not set")
		log.Println("Running without CRUD endpoints")
	} else {
		db, err := sql.Open("libsql", dbURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbURL, err)
			os.Exit(1)
		}
		defer db.Close()
		dbQueries := database.New(db)

		apiCfg.DB = dbQueries
	}

	mux := http.NewServeMux()
	corsMux := middlewareCors(mux)

	//Endpoints

	//Users
	mux.HandleFunc("GET /v1/readiness", apiCfg.handlerReady)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           corsMux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
