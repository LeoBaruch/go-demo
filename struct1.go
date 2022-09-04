package main

import (
	"fmt"
)

type Book struct {
	name string
	page int
}

// 若要改变函数需传指针
func changeBookName(this Book) {
	this.name = "changed"
}

func changeBOokName1(this *Book) {
	this.name = "changed1"
}

func main() {
	var book Book
	var book1 Book
	book2 := Book{name: "book3", page: 30}

	book.name = "book"
	book.page = 10

	book1.name = "book1"
	book1.page = 20

	changeBookName(book)
	changeBOokName1(&book1)

	fmt.Println("book", book)
	fmt.Println("book1", book1)
	fmt.Println("book2", book2)
}
