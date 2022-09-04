package main

import (
	"fmt"
	"runtime"
)

// runtime.Goexit() 将立即终止当前 goroutine 执⾏，调度器确保所有已注册 defer 延迟调用被执行
func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			fmt.Println("b")
		}()

		fmt.Println("A")
	}()

	// 确保主goroutine一直保持执行
	for {

	}
}
