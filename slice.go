package main

import (
	"fmt"
)

func main() {
	// 直接初始化
	sliceDemo := []int{1, 2, 3}

	fmt.Printf("slice demo: %T\n", sliceDemo)

	// make, [start: end]还是引用的原切片
	sliceDemo1 := make([]int, 3, 4)

	fmt.Printf("sliceDemo1 info: len(%d), cap(%d), info(%v)\n", len(sliceDemo1), cap(sliceDemo1), sliceDemo1)

	sliceDemo1[1] = 10
	sliceDemo1 = append(sliceDemo1, 5)
	sliceDemo1 = append(sliceDemo1, 6)
	sliceDemo2 := sliceDemo1[1:3]

	fmt.Printf("sliceDemo2 info: len(%d), cap(%d), info(%v)\n", len(sliceDemo2), cap(sliceDemo2), sliceDemo2)

	sliceDemo2[1] = 2

	fmt.Printf("new sliceDemo1 info: len(%d), cap(%d), info(%v)\n", len(sliceDemo1), cap(sliceDemo1), sliceDemo1)
	fmt.Printf("new sliceDemo2 info: len(%d), cap(%d), info(%v)\n", len(sliceDemo2), cap(sliceDemo2), sliceDemo2)

	// nil 判空
	var sliceNil []int

	if sliceNil == nil {
		fmt.Println("空切片!")
		fmt.Printf("len(%d), cap(%d), type(%T), info: %v\n", len(sliceNil), cap(sliceNil), sliceNil, sliceNil)
	}

	// append
	var sliceDemo3 = []int{1}
	sliceDemo3 = append(sliceDemo3, 2)
	fmt.Printf("sliceDemo3: %v\n", sliceDemo3)

	// copy 复制
	sliceCopied := []int{1, 2, 3}
	newSliceCopied := make([]int, len(sliceCopied), cap(sliceCopied))

	copy(newSliceCopied, sliceCopied)

	fmt.Printf("sliceCopoed: %v\n", sliceCopied)
	fmt.Printf("newSliceCoped: %v\n", newSliceCopied)

	sliceCopied[1] = 10

	fmt.Printf("sliceCopoed: %v\n", sliceCopied)
	fmt.Printf("newSliceCoped: %v\n", newSliceCopied)
}
