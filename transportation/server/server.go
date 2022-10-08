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

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/api/v1/addresses", c.AddressesController)
	r.POST("/api/v1/texts", c.TextsController)
	r.GET("/uploads/:path", c.UploadsController)
	r.GET("/api/v1/qrcodes", c.QrcodesController)

	r.StaticFS("/static", http.FS(staticFiles))

	r.Run(":" + config.Port)
}
