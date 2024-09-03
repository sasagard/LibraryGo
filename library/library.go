package library

import "fmt"

var books []Book

func AddBook(b Book) {
	books = append(books, b)
}

func ListBooks() {

	for _, book := range books {
		fmt.Printf("- %s, Autor %s (%d)\n", book.Title, book.Author, book.Year)
	}
}

func GetBooks() []Book {
	return books
}
