package main

import "fmt"

func main() {
	server := NewServer("127.0.0.1", 8888)
	fmt.Println("start")
	server.Start()
}
