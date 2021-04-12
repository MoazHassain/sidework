package main

import "fmt"

// type book struct {
// 	Title  string
// 	Author string
// 	ISBN   string
// 	Price  float32
// 	Pages  int
// }

func main() {
	// var b1 book
	// b1.Title = "An Introduction to programming in go"
	// b1.Author = "Caleb Doxsy"
	// b1.ISBN = "978-1478355823"
	// b1.Price = 10.50
	// b1.Pages = 165

	// fmt.Println(b1)
	// fmt.Println(b1.Title)
	// fmt.Println(b1.ISBN)
	// fmt.Println(b1.Author)
	// fmt.Println(b1.Price)

	// struct literals
	b1 := struct {
		title  string
		author string
		isbn   string
		price  float32
		pages  int
	}{
		title:  "an introduction to programming in GO",
		author: "caleb doxy",
		isbn:   "978-1478355823",
		price:  10.00,
		pages:  165,
	}

	fmt.Println(b1)

	var w1 weight

	fmt.Println(w1)
}
