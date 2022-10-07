package controller

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func getUploadsDir() (uploads string) {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err.Error())
	}
	dir := filepath.Dir(exe)
	uploads = filepath.Join(dir, "uploads")
	return
}

func UploadsController(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		c.Status(http.StatusNotFound)
		return
	}

	target := filepath.Join(getUploadsDir(), path)
	c.FileAttachment(target, path)
}
