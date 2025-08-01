package main

import (
	"log"
	"net/http"

	"github.com/vishvam7738/library-management/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/books", handlers.BooksHandler)
	mux.HandleFunc("/books/add", handlers.AddBookHandler)
	mux.HandleFunc("/books/delete", handlers.DeleteBookHandler)
}

func main() {
	mux := http.NewServeMux()
	RegisterRoutes(mux)

	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
