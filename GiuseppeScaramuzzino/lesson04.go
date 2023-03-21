package main

import "fmt"

func main() {

	// var title string
	// var author string
	// var title, author string
	// var pages int

	// var (
	// 	title string
	// 	author string
	// 	pages int
	// )

	title, author := "TITLE", "AUTHOR"
	pages := 420

	fmt.Printf("The author %s, The Book %s, the page %d\n", author, title, pages)
	pages = 200
	fmt.Printf("The author %s, The Book %s, the page %d\n", author, title, pages)

	title, author = "title1", "author1"

	pages=1000

	fmt.Printf("%s %s %d\n",author, title, pages)
}
