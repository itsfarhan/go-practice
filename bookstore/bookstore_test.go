package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
	// if !cmp.Equal(want, got) {
	// 	t.Error(cmp.Diff(want, got)) // Here, we log the difference between want and got using cmp.Diff. It will show what is missing in got compared to want using '-' and what is extra in got compared to want using '+'
	// }

	//This tells cmp.Equal to ignore any unexported fields on the bookstore.Book struct,preventing the unpleasant panic message.
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
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

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "Harry Potter"},
		2: {ID: 2, Title: "SuperWoman"},
	}

	want := []bookstore.Book{
		{ID: 1, Title: "Harry Potter"},
		{ID: 2, Title: "SuperWoman"},
	}

	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	// if !cmp.Equal(want, got) {
	// 	t.Error(cmp.Diff(want, got))
	// }
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:           "Harry Potter",
		PriceCents:      1000,
		DiscountPercent: 10,
	}
	want := 900
	got := b.NetPriceCents()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "Harry Potter",
		PriceCents: 4000,
	}
	want := 3000
	err := b.SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.PriceCents
	if want != got {
		t.Errorf("want updated price %d, got %d", want, got)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "Harry Potter",
		PriceCents: 4000,
	}
	err := b.SetPriceCents(-1)
	if err == nil {
		t.Fatal("want error while setting invalid price, but got nil")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "Harry Potter",
	}
	cats := []bookstore.Category{
		bookstore.CategoryComic,
		bookstore.CategoryFantasy,
		bookstore.CategoryAction,
	}
	for _, cat := range cats {
		err := b.SetCategory(cat)
		if err != nil {
			t.Fatal(err)
		}
		got := b.Category()
		if cat != got {
			t.Errorf("want categort %q, got %q", cat, got)
		}
	}
}

// 	want := "Comic"
// 	err := b.SetCategory(want)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	got := b.Category()
// 	if want != got {
// 		t.Errorf("want category %q, got %q", want, got)
// 	}
// }

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "Harry Potter",
	}
	err := b.SetCategory(999)
	if err == nil {
		t.Errorf("No error, got nit")
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
