package main

import (
	"net/http"

	"github.com/vishvam7738/library-management/handlers" // use the correct import path
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/books", handlers.BooksHandler)
	http.HandleFunc("/books/add", handlers.AddBookHandler)
	http.HandleFunc("/books/delete", handlers.DeleteBookHandler)
}
