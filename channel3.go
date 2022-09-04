package main

import (
	"fmt"
	"time"
)

// 有缓冲
func main() {
	c := make(chan int, 3)
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("子go程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(time.Second * 2)

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num =", num)
	}

	fmt.Println("主程结束!")
}
