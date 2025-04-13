package main

import (
	"log"
	"net/http"

	"skipr/internal/api"
)

func main() {
	r := api.SetupRouter()
	log.Println("server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
