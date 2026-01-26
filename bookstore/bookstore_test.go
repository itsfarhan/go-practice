package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	got := bookstore.GetBook(catalog, 2)
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
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
