package main

import (
	"fmt"
)

func Demo() {
	defer fmt.Println("1a")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4aa")

}

func main() {
	defer Demo()
	fmt.Printf("34324214\n")
}
