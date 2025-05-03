package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/davidl21/algotracker/server/data"
	"github.com/davidl21/algotracker/server/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}	

	// intialize db
	ctx := context.Background()
	dbURL := os.Getenv("DATABASE_URL")

	store, err := data.NewStore(ctx, dbURL)
	if err != nil {
		log.Fatalf("Failed to create store: %v", err)
	} 

	if err := store.Ping(ctx); err != nil {
		log.Fatalf("Database connection check failed: %v", err)
	} else {
		log.Println("Successfully pinged database!")
	}
	defer store.Close()
	
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	// register handlers
	ph := handlers.NewProblems(l, store)
	frh := handlers.NewFolders(l, store)
	serveMux := mux.NewRouter()
	
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.ServeHTTP)
	getRouter.HandleFunc("/folders", frh.ServeHTTP)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/folders", frh.CreateFolder)

	log.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", serveMux)
	if err != nil {
		log.Fatal(err)
	}
}