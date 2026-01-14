package bookstore

type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) Book { //Here we simulate buying a book by reducing the number of copies by 1
	if b.Copies > 0 {
		b.Copies--
	}
	return b
}
