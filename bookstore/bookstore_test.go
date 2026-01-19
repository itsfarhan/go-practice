package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
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
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatalf("%v", err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuyErrorIfNoCopies(t *testing.T) {
	t.Parallel()
	var b = bookstore.Book{
		Title:  "Farhan Go-Practice",
		Author: "Farhan",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Fatalf("want error buying from 0 copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	got := bookstore.GetAllBooks(catalog)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
