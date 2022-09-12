package main

import (
	"flag"
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

var ServerIp string
var ServerPort int

func init() {
	flag.StringVar(&ServerIp, "ip", "127.0.0.1", "设置服务器ip地址(默认127.0.0.1)")
	flag.IntVar(&ServerPort, "port", 8888, "设置服务器端口(默认8888)")
}

func main() {
	flag.Parse()

	client := NewClient(ServerIp, ServerPort)

	if client == nil {
		fmt.Println("链接服务器失败!")
		return
	}

	fmt.Println("链接服务器成功!")

	select {}
}
