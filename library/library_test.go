package library

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	books = nil
	b := Book{"El economista callejero", "Axel Kaiser", 2021}
	AddBook(b)

	assert.Equal(t, 1, len(books), "Espera 1 libro en la lista")
	assert.Equal(t, b, books[0], "Espera el mismo 'El economista callejero'")

}

func TestListBooks(t *testing.T) {
	books = nil
	b1 := Book{"Libro 1", "Autor 1", 2001}
	b2 := Book{"Libro 2", "Autor 2", 2002}
	AddBook(b1)
	AddBook(b2)

	result := GetBooks()

	assert.Equal(t, 2, len(result), "Espera 2 libros")
	assert.Equal(t, []Book{b1, b2}, result, "Espera 'Libro 1' y 'Libro 2'")
}
