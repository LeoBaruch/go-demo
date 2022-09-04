package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 1.2345
	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("typeName: ", reflect.TypeOf(num).Name())
	fmt.Println("value: ", reflect.ValueOf(num))

}
