package bookstore

//"errors"
// "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
	ID   int
}

func GetBook(catalog map[int]Book, ID int) Book {
	return catalog[ID]
}

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
