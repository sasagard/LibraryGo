package library

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	items = nil

	pBook := PrintedBook{
		Book: Book{
			Title:  "The Great Gatsby",
			Author: "F. Scott Fitzgerald",
			Year:   1925,
		},
		Pages: 180,
	}

	eBook := Ebook{
		Book: Book{
			Title:  "The Catcher in the Rye",
			Author: "J. D. Salinger",
			Year:   1951,
		},
		FileSizeMB: 500,
	}

	AddItem(pBook)
	AddItem(eBook)

	assert.Equal(t, 2, len(items), "Espera 2 libro en la lista")
	assert.Equal(t, pBook, items[0], "Se espera el libro The Great Gatsby")
	assert.Equal(t, eBook, items[1], "Se espera el eBook The Catcher in the Rye")

}

// Test library.ListBooks
func TestListBooks(t *testing.T) {
	items = nil

	pb := PrintedBook{
		Book:  Book{Title: "Libro 1", Author: "Autor 1", Year: 2001},
		Pages: 150,
	}
	eb := Ebook{
		Book:       Book{Title: "Libro 2", Author: "Autor 2", Year: 2002},
		FileSizeMB: 3,
	}

	AddItem(pb)
	AddItem(eb)

	result := []string{}
	for _, item := range items {
		result = append(result, item.GetDetails())
	}

	assert.Equal(t, 2, len(result), "Se esperan 2 libros")
	assert.Equal(t, "Libro 1 (2001) - Autor: Autor 1 - Pag: 150", result[0], "Espera 'Libro 1'")
	assert.Equal(t, "Libro 2 (2002) - Autor: Autor 2 - Tamaño: 3 MB", result[1], "Espera 'Libro 2'")

}
