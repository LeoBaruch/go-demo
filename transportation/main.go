package main

import (
	"embed"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"time"
	"transportation/config"
	"transportation/server"
)

//go:embed server/frontend/dist/*
var FS embed.FS

func main() {
	go openBrowser()

	server.Run()

}

func openBrowser() {
	url := "http://127.0.0.1:" + config.Port + "/static/index.html"
	commands := map[string]string{
		"windows": "start",
		"darwin":  "open",
		"linux":   "xdg-open",
	}
	os := runtime.GOOS

	command, ok := commands[os]
	if !ok {
		log.Fatal("未知操作系统！")
	}

	var cmd *exec.Cmd

	if strings.HasPrefix(os, "window") {
		cmd = exec.Command(`cmd`, `/c`, command, url)
	} else {
		cmd = exec.Command(command, url)
	}

	time.After(time.Second * 1)
	cmd.Start()
}
