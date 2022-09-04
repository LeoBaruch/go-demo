package main

import (
	"fmt"
)

// make(chan Type, capacity) capacity 为0 或者无时,为无缓冲阻塞读写的
func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("子goroutine 结束")

		fmt.Println("子goroutine 进行...")

		c <- 666
	}()

	num, _ := <-c

	fmt.Println("num = ", num)
	fmt.Println("main go程结束")
}
