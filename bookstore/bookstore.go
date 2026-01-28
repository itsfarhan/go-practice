package bookstore

import (
	// "errors"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func GetBook(catalog map[int]Book, ID int) (Book, error) {
	b, ok := catalog[ID] // Here, ok is boolean (true or false)
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}

func GetAllBooks(catalog map[int]Book) []Book {
	books := []Book{}
	for ID := range catalog {
		fmt.Println(ID)
	}
	for _, b := range catalog {
		books = append(books, b)
	}
	return books
}

// modify maps in Go is done by first retrieving the struct, modifying its fields, and then reassigning it back to the map.

// b := catalog[1]
// b.Title = "New Title"
// catalog[1] = b

// var books = []Book {
// 	{Title: "Harry Potter"},
// 	{Title: "SuperGirl: Woman of Tommorrow"},
// }

// func Buy(b Book) (Book, error) {
// 	if b.Copies == 0 {
// 		return Book{}, errors.New("no copies available")
// 	}
// 	b.Copies--
// 	return b, nil
// }
