package main

import (
	"AppLibrary/library"
	"fmt"
)

func main() {

	book := library.Book{Title: "El economista callejero", Author: "Axel Kaiser", Year: 2021}
	library.AddBock(book)
	book2 := library.Book{Title: "El Odio a los Ricos", Author: "Axel Kaiser y Rainer Zitelmann", Year: 2023}
	library.AddBock(book2)

	fmt.Println("Libros en la libreria:")
	library.ListBooks()
}
