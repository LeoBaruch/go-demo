package server

import (
	"embed"
	"io/fs"
	"net/http"
	"transportation/config"
	c "transportation/server/controller"

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
	r.GET("/api/v1/addresses", c.AddressesController)
	r.POST("/api/v1/texts", c.TextsController)

	r.StaticFS("/static", http.FS(staticFiles))

	r.Run(":" + config.Port)
}
