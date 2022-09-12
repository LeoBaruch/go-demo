package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	go user.ListenMessage()

	return user
}

func (this *User) Online() {
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "已上线")
}

func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "下线")
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg + "\n"))
}

func (this *User) DoMessage(msg string) {
	// 当msg为who时, 查询当前所有在线用户
	if msg == "who" {
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"

			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()

	} else if (len(msg) > 7) && strings.HasPrefix(msg, "rename|") {
		newName := strings.Split(msg, "|")[1]

		_, ok := this.server.OnlineMap[newName]

		if ok {
			this.SendMsg("当前用户名已被使用!")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("您已经更新新用户名: " + this.Name + "\n")
		}

		// 添加私聊功能
	} else if strings.HasPrefix(msg, "to|") && len(msg) > 4 {
		remoteName := strings.Split(msg, "|")[1]
		tips := "请使用形如\"to|张三|你好啊!\"的格式"

		if remoteName == "" {
			this.SendMsg("消息格式不正确," + tips + "\n")
			return
		}

		remoteUser, ok := this.server.OnlineMap[remoteName]

		if !ok {
			this.SendMsg("当前用户不存在!")
			return
		}

		content := strings.Split(msg, "|")[2]

		if content == "" {
			this.SendMsg("消息为空," + tips + "\n")
			return
		}

		remoteUser.SendMsg(this.Name + "对您说:" + content + "\n")

	} else {
		this.server.BroadCast(this, msg)
	}

}

func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
