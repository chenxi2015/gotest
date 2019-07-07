package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	fmt.Println(Books{"Go语言", "chenxi", "Go语言教程", 1})

	fmt.Println(Books{title: "php语言", author: "chenxi"})

	var Book1 Books
	var Book2 Books

	Book1.title = "Java语言"
	Book1.author = "gengxin"

	Book2.title = "Python语言"
	Book2.subject = "Python语言教程"

	fmt.Printf("Book1 title: %s\n", Book1.title)
	fmt.Printf("Book1 author: %s\n", Book1.author)

	fmt.Printf("Book2 title: %s\n", Book2.title)
	fmt.Printf("Book2 author: %s\n", Book2.author)
	fmt.Printf("Book2 subject: %s\n", Book2.subject)
}
