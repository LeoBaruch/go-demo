package main

import (
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))

	if err != nil {
		fmt.Println("err: ", err)
		return nil
	}

	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		conn:       conn,
	}

	return client
}

func main() {
	client := NewClient("127.0.0.1", 8888)

	if client == nil {
		fmt.Println("链接服务器失败!")
		return
	}

	fmt.Println("链接服务器成功!")

	select {}
}
