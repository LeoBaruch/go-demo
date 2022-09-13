package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn

	flag int
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
		flag:       999,
	}

	return client
}

func (this *Client) Menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		this.flag = flag
		return true
	}

	fmt.Println("请输入合法数字!")
	return false
}

func (this *Client) DealResponse() {
	io.Copy(os.Stdout, this.conn)
}

func (this *Client) UpdateName() bool {
	var name string
	fmt.Println("请输入用户名: ")
	fmt.Scanln(&name)
	senMsg := "rename|" + name + "\n"

	_, err := this.conn.Write([]byte(senMsg))

	if err != nil {
		fmt.Println("conn.Write err: ", err)

		return false
	}

	this.Name = name
	return true
}

func (this *Client) PublicChat() {
	var chatMsg string
	fmt.Println("请输入公聊内容:")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		sendMsg := chatMsg + "\n"
		this.conn.Write([]byte(sendMsg))

		fmt.Println("请输入公聊内容:")
		fmt.Scanln(&chatMsg)
	}
}

func (this *Client) SelectUser() {
	sendMsg := "who\n"

	_, err := this.conn.Write([]byte(sendMsg))

	if err != nil {
		fmt.Println("conn err: ", err)
	}
}

func (this *Client) PriviteChat() {
	this.SelectUser()
	var chatMsg string

	var userName string
	fmt.Println("请输入私聊用户名:")
	fmt.Scanln(&userName)

	fmt.Println("请输入私聊内容:")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		sendMsg := "to|" + userName + "|" + chatMsg + "\n"
		this.conn.Write([]byte(sendMsg))

		fmt.Println("请输入私聊内容:")
		fmt.Scanln(&chatMsg)
	}
}

func (this *Client) Run() {
	for this.flag != 0 {
		for this.Menu() != true {
		}

		switch this.flag {
		case 1:
			this.PublicChat()
			break
		case 2:
			this.PriviteChat()
			break
		case 3:
			this.UpdateName()
			break
		}
	}
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

	go client.DealResponse()
	client.Run()
}
