package main

import (
	"log"
	"net/http"
)

func main() {
	//l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	mux := http.NewServeMux()

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}