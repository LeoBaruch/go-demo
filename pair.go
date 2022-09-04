package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("write a book")
}

func main() {
	b := &Book{}

	var r Reader

	r = b

	r.ReadBook()

	var w Writer

	// 类型断言,实现了Writer
	w = r.(Writer)

	w.WriteBook()
}
