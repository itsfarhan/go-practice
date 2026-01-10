package main

import (
	"fmt"
)

func main() {
	var title string
	var copies int
	var author string
	var edition string
	var inStock bool
	var royaltyPercentage float64
	var x int
	title = "Farhan Go-Practice"
	copies = 5
	author = "Farhan Ahmed"
	edition = "first edition"
	inStock = true
	royaltyPercentage = 7.5
	fmt.Println(title)
	fmt.Println(copies)
	fmt.Println(author)
	fmt.Println(edition)
	fmt.Println(inStock)
	fmt.Println(royaltyPercentage)
	fmt.Println(x) // since value has not been assigned, therefore value is '0' by default
	y := 1 // Directly initialize the value without var
	fmt.Println(y)
}
