package library

import "fmt"

// Interfaz para representar elementos de la biblioteca
type LibraryItem interface {
	GetDetails() string
}

// Estructura de PrintedBook
type PrintedBook struct {
	Book
	Pages int
}

// Estructura de Ebook
type Ebook struct {
	Book
	FileSizeMB int
}

// Funcion para obtener los detalles de PrintedBook
func (pb PrintedBook) GetDetails() string {
	return fmt.Sprintf("%s (%d) - Autor: %s - Pag: %d", pb.Title, pb.Year, pb.Author, pb.Pages)
}

// Funcion para obtener los detalles de Ebook
func (eb Ebook) GetDetails() string {
	return fmt.Sprintf("%s (%d) - Autor: %s - Tamaño: %d MB", eb.Title, eb.Year, eb.Author, eb.FileSizeMB)
}

var items []LibraryItem

func AddItem(item LibraryItem) {
	items = append(items, item)
}

func ListItems() {
	for _, item := range items {
		fmt.Println(item.GetDetails())
	}
}
