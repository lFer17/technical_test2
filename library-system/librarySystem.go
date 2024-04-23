package main

import "fmt"

type book struct {
	name      string
	author    string
	genre     string
	available bool
}

func main() {
	fmt.Println("Gestiona tu libreria")
	var storageLocal []book
	CreateBook(&storageLocal, "nuevo", "Elver Gatiesa", "comedia")
	fmt.Printf("Almacenado actualmente: %s \n", storageLocal)
	fmt.Println(FindBookTittleOrAuthor(&storageLocal, "nuevonuevo"))
	UpdateBookStatus(&storageLocal, "nuevo")
	fmt.Println(FindBookTittleOrAuthor(&storageLocal, "nuevo"))
}

func CreateBook(storage *[]book, name string, author string, genre string) {
	var newBook = book{name: name, author: author, genre: genre, available: true}
	*storage = append(*storage, newBook)
}

func FindBookTittleOrAuthor(storage *[]book, bookTitleOrAuthor string) book {

	var BookFound book

	for _, bookObject := range *storage {

		if bookObject.name == bookTitleOrAuthor || bookObject.author == bookTitleOrAuthor {
			BookFound = bookObject
			return BookFound
		}
	}
	return BookFound
}

func UpdateBookStatus(storage *[]book, bookTitle string) {
	for objectPosition, bookObject := range *storage {
		if bookObject.name == bookTitle && bookObject.available == true {
			(*storage)[objectPosition].available = false
		} else {
			(*storage)[objectPosition].available = true
		}
	}
}

func DeleteBook(storage *[]book, bookTitle string) {

	for objectPosition, bookobject := range *storage {
		if bookobject.name == bookTitle {
			(*storage)[objectPosition] = (*storage)[len(*storage)-1]
			*storage = (*storage)[:len(*storage)-1]
		}
	}

}
