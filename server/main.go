package main

import (
	"log"
	"net/http"
	"os"

	"github.com/davidl21/algotracker/server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	ph := handlers.NewProblems(l)

	serveMux := mux.NewRouter()
	
	// register handlers
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.ServeHTTP)

	// putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	// postRouter := serveMux.Methods(http.MethodPost).Subrouter()

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", serveMux)
	if err != nil {
		log.Fatal(err)
	}
}