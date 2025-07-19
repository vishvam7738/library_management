package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{}
var idCounter = 1

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "📚 Welcome to the Library Management System!")
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var b Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	b.ID = idCounter
	idCounter++
	books = append(books, b)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE method allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	idStr := query.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Book with ID %d deleted", id)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}
