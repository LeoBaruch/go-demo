package main

import "fmt"

type AnimalIF interface {
	Sleep()
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("The doggie is sleeping")
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("The cat is sleeping")
}

func (this *Cat) Eating() {
	fmt.Println("The cat is eating")
}

func main() {
	var animal AnimalIF

	animal = &Dog{"red"}
	animal.Sleep()
	fmt.Println(animal)
	animal = &Cat{"black"}
	fmt.Printf("cat, %v\n", animal)
}
