package library

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	database "AppLibrary/db"
)

// Interfaz para representar elementos de la biblioteca
type LibraryItem interface {
	GetDetails() string
}

var items []LibraryItem

func AddItem(item LibraryItem) {
	var doc interface{}
	switch v := item.(type) {
	case PrintedBook:
		doc = v
	case Ebook:
		doc = v
	default:
		fmt.Println("No se acepta este tipo de elemento: %T", v)
	}

	_, err := database.Collection.InsertOne(context.TODO(), doc)
	if err != nil {
		items = append(items, item)
	}
}

func ListItems() {

	cursor, err := database.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}

func CreatePrintedBook() PrintedBook {

	var title, author string
	var year, pages int

	fmt.Print("Ingrese Title: ")
	fmt.Scanln(&title)
	fmt.Print("Ingrese Autor: ")
	fmt.Scanln(&author)
	fmt.Print("Ingrese Año de publicación: ")
	fmt.Scanln(&year)
	fmt.Print("Ingrese numero de páginas: ")
	fmt.Scanln(&pages)

	return PrintedBook{
		Book: Book{
			Title:  title,
			Author: author,
			Year:   year,
		},
		Pages: pages,
	}
}

func CreateEbook() Ebook {
	var title, author string
	var year, fileSizeMB int

	fmt.Print("Ingrese Title: ")
	fmt.Scanln(&title)
	fmt.Print("Ingrese Autor: ")
	fmt.Scanln(&author)
	fmt.Print("Ingrese Año de publicación: ")
	fmt.Scanln(&year)
	fmt.Print("Ingrese Tamaño en MB: ")
	fmt.Scanln(&fileSizeMB)

	return Ebook{
		Book: Book{
			Title:  title,
			Author: author,
			Year:   year,
		},
		FileSizeMB: fileSizeMB,
	}
}
