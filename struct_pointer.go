package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Book
	var Book2 Book

	Book1.title = "Go语言"
	Book1.author = "chenxi"
	Book2.title = "PHP语言"
	Book2.author = "gengxin"

	printBook(&Book1)
	printBook(&Book2)
}

func printBook(book *Book) {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book book_id: %d\n", book.book_id)
}
