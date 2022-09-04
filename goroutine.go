package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0

	for {
		i++
		fmt.Printf("new goroutine: i = %d\n", i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go newTask()

	i := 0
	for {
		i++
		fmt.Println("main goroutine: i = ", i)
		time.Sleep(time.Second * 1)
	}
}
