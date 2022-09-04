package main

import "fmt"

type Man struct {
	name string
	age  int
}

func (this *Man) Walk() {
	fmt.Println("The man is walking")
}

type SuperMan struct {
	Man
	level int
}

func (this *SuperMan) Walk() {
	fmt.Println("The man is running")
}

func (this *SuperMan) Fly() {
	fmt.Println("The man is Flying")
}

func main() {
	a := SuperMan{Man{"mortal", 10}, 10}

	var b SuperMan

	b.name = "superman"
	b.age = 20
	b.level = 100

	a.Walk()
	b.Walk()
	b.Fly()

	fmt.Printf("superman, {%v}\n", b)
	fmt.Printf("man, {%v}", a)
}
