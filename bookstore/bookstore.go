package bookstore

//"errors"
// "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
	ISBN   int
}

func GetAllBooks(catalog []Book, ISBN int) Book {
	for _, b := range catalog {
		if b.ISBN == ISBN {
			return b
		}
	}
	return Book{} // why Book{} -
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
