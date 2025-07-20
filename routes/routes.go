package routes

import (
	"net/http"

	"github.com/vishvam7738/library-management/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/books", handlers.BooksHandler)
	mux.HandleFunc("/books/add", handlers.AddBookHandler)
	mux.HandleFunc("/books/delete", handlers.DeleteBookHandler)
}
