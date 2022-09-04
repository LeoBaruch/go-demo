package main

import "fmt"

func main() {
	map1 := map[string]string{
		"one": "123",
	}

	map1["two"] = "two"

	fmt.Println("map1: \n", map1)

	map2 := make(map[string]string)

	map2["one"] = "one"
	map2["two"] = "two"
	map2["three"] = "three"

	fmt.Println("map2: \n", map2)
}
