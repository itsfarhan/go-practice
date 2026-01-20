package bookstore_test

import (
	"bookstore"
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "Harry potter", ISBN: 1},
		// {Title: "Superman", ISBN: 2},	
	}
	want := bookstore.Book{Title: "Harry potter", ISBN: 1}
		// {Title: "Superman", ISBN: 2},

	got := bookstore.GetAllBooks(catalog, 1) // here bookstore.GetAllBooks called catalog for getting all books
	if !cmp.Equal(want, got) { // here cmp is used to compare want and got 
		t.Error(cmp.Diff(want, got)) // Equal and diff are methods of cmp package and used to compare two values
	}
}

// func TestBookstore(t *testing.T) {
// 	t.Parallel()
// 	// Placeholder for future tests related to the bookstore package
// 	_ = bookstore.Book{
// 		Title:  "Farhan Go-Practice",
// 		Author: "Farhan",
// 		Copies: 10,
// 	}
// }

// func TestBuy(t *testing.T) {
// 	t.Parallel()
// 	// Placeholder
// 	b := bookstore.Book{
// 		Title:  "Farhan Go-Practice",
// 		Author: "Farhan",
// 		Copies: 2,
// 	}

// 	want := 1
// 	result, err := bookstore.Buy(b)
// 	if err != nil {
// 		t.Fatalf("%v", err)
// 	}
// 	got := result.Copies
// 	if want != got {
// 		t.Errorf("want %d, got %d", want, got)
// 	}
// }

// func TestBuyErrorIfNoCopies(t *testing.T) {
// 	t.Parallel()
// 	var b = bookstore.Book{
// 		Title:  "Farhan Go-Practice",
// 		Author: "Farhan",
// 		Copies: 0,
// 	}
// 	_, err := bookstore.Buy(b)
// 	if err == nil {
// 		t.Fatalf("want error buying from 0 copies, got nil")
// 	}
// }
