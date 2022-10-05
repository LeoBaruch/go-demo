package main

import (
	"embed"
	"log"
	"transpotation/config"
	"transpotation/server"

	"github.com/zserge/lorca"
)

//go:embed server/frontend/dist/*
var FS embed.FS

func main() {
	go server.Run()

	ui, _ := openBrowser()
	<-ui.Done()
}

func openBrowser() (lorca.UI, error) {
	ui, err := lorca.New("http://127.0.0.1:"+config.Port+"/static/index.html", "", 1000, 800)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return ui, nil
}
