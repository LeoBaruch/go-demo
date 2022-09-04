package main

import (
	"fmt"
	"time"
)

// 无缓冲循环
func main() {
	c := make(chan int, 0)

	fmt.Printf("len(c) = %d, cap(c) = %d \n", len(c), cap(c))

	go func() {
		for i := 0; i < 3; i++ {
			c <- i

			fmt.Printf("子go程正在运行[%d]: len(c) = %d, cap(c) = %d \n", i, len(c), cap(c))
		}
	}()

	time.Sleep(time.Second * 2)

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}

	fmt.Println("main 结束`")
}
