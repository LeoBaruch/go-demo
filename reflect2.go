package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 1.23456

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)
	// 当明确知道value的类型,直接通过value的interface().(已知类型)方法强制转换
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println((convertPointer))
	fmt.Println((convertValue))
}
