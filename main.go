package main

import (
	"AppLibrary/library"
	"fmt"
)

func main() {

	fmt.Println("Agregando elementos a biblioteca")
	printedBook := library.PrintedBook{
		Book: library.Book{
			Title:  "The Great Gatsby",
			Author: "F. Scott Fitzgerald",
			Year:   1925,
		},
		Pages: 180,
	}
	eBook := library.Ebook{
		Book: library.Book{
			Title:  "The Catcher in the Rye",
			Author: "J. D. Salinger",
			Year:   1951,
		},
		FileSizeMB: 500,
	}

	library.AddItem(printedBook)
	library.AddItem(eBook)

	library.ListItems()
}
