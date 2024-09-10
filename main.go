package main

import (
	"AppLibrary/library"
	"fmt"
)

func main() {
	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Listar libros existentes")
		fmt.Println("2. Ingresar nuevo libro")
		fmt.Println("3. Salir")

		var choice int
		fmt.Print("Ingrese una opción (1,2,3): ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\n")
			library.ListItems()
		case 2:
			fmt.Println("Elija el tipo de libro:")
			fmt.Println("1. Libro impreso")
			fmt.Println("2. Ebook")

			var typeChoice int
			fmt.Print("Ingrese una opción (1,2): ")
			fmt.Scanln(&typeChoice)

			switch typeChoice {
			case 1:
				library.AddItem(library.CreatePrintedBook())
			case 2:
				library.AddItem(library.CreateEbook())
			default:
				fmt.Println("Opción no disponible")
			}
		case 3:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no disponible")
		}
	}

}
