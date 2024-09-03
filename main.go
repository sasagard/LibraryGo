package main

import (
	"AppLibrary/library"
	"fmt"
)

func main() {

	book := library.Book{
		Title:  "EL ECONOMISTA CALLEJERO",
		Author: "Axel Kaiser",
		Year:   2021,
	}

	fmt.Printf("Mi primer libro es %s del autor %s (%d)\n", book.Title, book.Author, book.Year)
}
