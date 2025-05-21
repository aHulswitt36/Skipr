package main

import (
	"log"
	"net/http"

	"skipr/internal/api"
	"skipr/internal/store"
)

func main() {
    // Initialize the database
    if err := store.InitDatabase(); err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }

    // Create the team and players
    if err := store.CreateTheMets(store.DB); err != nil {
        log.Fatalf("Failed to create team: %v", err)
    }

	r := api.SetupRouter()
	log.Println("server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
