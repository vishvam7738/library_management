package main

import (
	"log"
	"net/http"
)

func main() {
	// Register routes
	RegisterRoutes()

	// Start server
	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
