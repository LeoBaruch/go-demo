package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message

		this.mapLock.Lock()

		for _, client := range this.OnlineMap {
			client.C <- msg
		}

		this.mapLock.Unlock()
	}
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
}

func (this *Server) Handler(connect net.Conn) {
	user := NewUser(connect)

	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	this.BroadCast(user, "已上线")

	go func() {
		buf := make([]byte, 4096)

		for {
			n, err := connect.Read(buf)

			if n == 0 {
				this.BroadCast(user, "下限")

				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("connect read err: ", err)

				return
			}

			// 提取用户消息(去掉\n)
			msg := string(buf[:n-1])

			this.BroadCast(user, msg)
		}
	}()

	// 阻塞当前goroutine
	select {}
}

func (this *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Println("net.Listen err: ", err)

		return
	}

	defer listener.Close()

	// 注意新开一个go程去监听
	go this.ListenMessage()

	for {
		connect, err := listener.Accept()

		if err != nil {
			fmt.Println("listener accept err: ", err)

			continue
		}

		go this.Handler(connect)

	}
}
