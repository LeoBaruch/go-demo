package controller

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Json struct {
	Raw string `form:"raw" json:"raw" xml:"raw"  binding:"required"`
}

func TextsController(c *gin.Context) {
	var json Json

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err.Error())
	}

	dir := filepath.Dir(exe) // 可执行文件目录
	uploads := filepath.Join(dir, "uploads")
	filename := uuid.New().String()
	err = os.MkdirAll(uploads, fs.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}

	fullpath := filepath.Join("uploads", filename+".txt") // 文件上传路径
	err = os.WriteFile(path.Join(dir, fullpath), []byte(json.Raw), fs.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"url": "/" + fullpath}) // 返回文件的绝对路径（不含 exe 所在目录）
}
