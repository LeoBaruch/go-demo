package main

import "fmt"

func getType(a interface{}) string {
	if a == nil {
		return "null"
	}

	value, ok := a.(string)

	if !ok {
		fmt.Println("It's not a string")

		return ""
	} else {
		fmt.Println("The value is ", value)
		return value
	}
}

func main() {
	getType("ss")
}
