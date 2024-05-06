package main

import "fmt"

// Estructura de un libro
type book struct {
	name      string
	author    string
	genre     string
	available bool
}

// Librería representa la colección de libros y proporciona operaciones relacionadas
type Library struct {
	storage []book
}

// Función principal
func main() {
	fmt.Println("Gestiona tu librería")

	// Crear una nueva librería
	lib := NewLibrary()

	// Agregar un nuevo libro
	lib.CreateBook("nuevo", "Elver Gatiesa", "comedia")

	// Imprimir la lista de libros
	fmt.Println("Libros almacenados:")
	lib.PrintBooks()

	// Encontrar un libro por título o autor
	bookFound := lib.FindBookTitleOrAuthor("nuevonuevo")
	if bookFound.name != "" {
		fmt.Println("Libro encontrado:", bookFound)
	} else {
		fmt.Println("Libro no encontrado.")
	}

	// Actualizar el estado de un libro
	lib.UpdateBookStatus("nuevo")
	fmt.Println("Estado del libro 'nuevo' actualizado.")

	// Imprimir la lista de libros actualizada
	fmt.Println("Libros almacenados actualizados:")
	lib.PrintBooks()
}

// Crea una nueva librería
func NewLibrary() *Library {
	return &Library{}
}

// Agrega un nuevo libro a la librería
func (lib *Library) CreateBook(name, author, genre string) {
	newBook := book{name: name, author: author, genre: genre, available: true}
	lib.storage = append(lib.storage, newBook)
}

// Encuentra un libro por título o autor
func (lib *Library) FindBookTitleOrAuthor(titleOrAuthor string) book {
	var bookFound book

	for _, bookObject := range lib.storage {
		if bookObject.name == titleOrAuthor || bookObject.author == titleOrAuthor {
			bookFound = bookObject
			break
		}
	}

	return bookFound
}

// Actualiza el estado de un libro
func (lib *Library) UpdateBookStatus(title string) {
	for i := range lib.storage {
		if lib.storage[i].name == title {
			lib.storage[i].available = !lib.storage[i].available
			break
		}
	}
}

// Imprime todos los libros en la librería
func (lib *Library) PrintBooks() {
	for _, book := range lib.storage {
		fmt.Println(book)
	}
}
