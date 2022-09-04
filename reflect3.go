package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("reflect called")
}

func main() {
	user := User{1, "leo", 27}

	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get type is: ", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get value is: ", getValue)
	// 未知原有类型【遍历探测其Filed】
	for i := 0; i < getType.NumField(); i++ {
		filed := getType.Field(i)
		value := getValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", filed.Name, filed.Type, value)
	}
}
