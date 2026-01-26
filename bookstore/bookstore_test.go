package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetBook(t *testing.T) {
	t.Parallel()
	// Here, catalog is a map with book ID as key and Book struct as value
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "Harry Potter"},
		2: {ID: 2, Title: "SuperWoman"},
	}
	want := bookstore.Book{ID: 2, Title: "SuperWoman"}

	// Here, we are testing the GetBook function to retrieve the book with ID 2
	got, err := bookstore.GetBook(catalog, 2) // got is a single Book struct and its a result of GetBook function. err is the error returned by GetBook function.
	if err != nil {                           // if err is not nil, then there was an error in getting the book
		// Here, we use t.Fatal to log the error and stop the test immediately
		t.Fatal(err) // if is use fatalf then I need to use format specifier like - t.Fatalf("error: %v", err)
	}
	// Here, we compare the want and got values using cmp.Equal
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got)) // Here, we log the difference between want and got using cmp.Diff. It will show what is missing in got compared to want using '-' and what is extra in got compared to want using '+'
	}
}

// Here, we are adding a test case to check for non-existing book ID
func TestGetBookBadIDreturnsError(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{}
	_, err := bookstore.GetBook(catalog, 999)
	if err == nil {
		t.Fatal("want error for non-exist ID, but I got nil")
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
// 		t.Fatalf("want error because cant but 0 copies, but got nil")
// 	}
// }
