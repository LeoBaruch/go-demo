package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0

	for {
		i++
		fmt.Println("new goroutine: i = ", i)
		time.Sleep(time.Second * 1)
	}
}

// 当主goroutine退出, 其他工作routine也退出
func main() {
	go newTask()

	fmt.Println("main goroutine exit;")
}
