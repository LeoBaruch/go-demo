package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
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
	user := NewUser(connect, this)

	// 用户上线
	user.Online()

	isLive := make(chan bool)

	go func() {
		buf := make([]byte, 4096)

		for {
			n, err := connect.Read(buf)

			if n == 0 {
				user.Offline()

				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("connect read err: ", err)

				return
			}

			// 提取用户消息(去掉\n)
			msg := string(buf[:n-1])

			user.DoMessage(msg)

			// 刷新用户活跃状态
			isLive <- true
		}
	}()

	for {
		select {
		case <-isLive:

		case <-time.After(time.Second * 10):
			user.SendMsg("超时被踢!")
			close(user.C)
			connect.Close()

			return
		}
	}
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
