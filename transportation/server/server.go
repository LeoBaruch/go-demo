package server

import (
	"embed"
	"io/fs"
	"net/http"
	"transportation/config"
	c "transportation/server/controller"
	"transportation/server/ws"

	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist/*
var FS embed.FS

func Run() {
	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	hub := ws.NewHub()
	go hub.Run()

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/api/v1/addresses", c.AddressesController)
	r.POST("/api/v1/texts", c.TextsController)
	r.GET("/uploads/:path", c.UploadsController)
	r.GET("/api/v1/qrcodes", c.QrcodesController)
	r.POST("/api/v1/files", c.FilesController)

	r.GET("/ws", func(ctx *gin.Context) {
		ws.HttpController(ctx, hub)
	})

	r.StaticFS("/static", http.FS(staticFiles))

	r.NoRoute(func(ctx *gin.Context) {
		c.NoRoute(ctx, staticFiles)
	})

	r.Run("0.0.0.0:" + config.Port)
}
