package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"transportation/config"
	c "transportation/server/controller"

	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist/*
var FS embed.FS

func Run() {
	staticFiles, _ := fs.Sub(FS, "frontend/dist")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/api/v1/addresses", c.AddressesController)
	r.POST("/api/v1/texts", c.TextsController)
	r.GET("/uploads/:path", c.UploadsController)
	r.GET("/api/v1/qrcodes", c.QrcodesController)
	r.POST("/api/v1/files", c.FilesController)

	r.StaticFS("/static", http.FS(staticFiles))

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/static") {
			file, err := staticFiles.Open("index.html")
			if err != nil {
				log.Fatal(err.Error())
			}
			defer file.Close()

			stat, err := file.Stat()
			if err != nil {
				log.Fatal(err.Error())
			}

			c.DataFromReader(http.StatusOK, stat.Size(), "text/html;charset=utf-8", file, nil)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	r.Run("0.0.0.0:" + config.Port)
}
