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
	PriceCents int
	DiscountPercent int
}

type Catalog map[int]Book

func GetBook(catalog map[int]Book, ID int) (Book, error) {
	b, ok := catalog[ID] // Here, ok is boolean (true or false)
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

func (b Book) NetPriceCents() int { // To declare a method, we use the syntax func (receiver Type) MethodName() ReturnType
	savings := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - savings
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
