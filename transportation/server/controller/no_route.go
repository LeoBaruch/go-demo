package controller

import (
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context, staticFiles fs.FS) {
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
}
