package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBookstore(t *testing.T) {
	t.Parallel()
	// Placeholder for future tests related to the bookstore package
	_ = bookstore.Book{
		Title:  "Farhan Go-Practice",
		Author: "Farhan",
		Copies: 10,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	// Placeholder
	b := bookstore.Book{
		Title:  "Farhan Go-Practice",
		Author: "Farhan",
		Copies: 2,
	}

	want := 1
	result := bookstore.Buy(b)
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
