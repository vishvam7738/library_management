package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// BookStore keeps track of books and next ID
var (
	Books  = []Book{}
	NextID = 1
)
