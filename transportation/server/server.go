package server

import (
	"embed"
	"io/fs"
	"net/http"
	"transpotation/config"

	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist/*
var FS embed.FS

func Run() {
	staticFiles, _ := fs.Sub(FS, "frontend/dist")

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("hi!"))
	})

	r.StaticFS("/static", http.FS(staticFiles))

	r.Run(":" + config.Port)
}
