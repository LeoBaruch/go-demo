package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Json struct {
	Raw string `form:"raw" json:"raw" xml:"raw"  binding:"required"`
	// Raw string `json: "raws" binding:"required"`
}

func TextsController(c *gin.Context) {
	var json Json

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("11111111111111received: " + json.Raw)
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": json.Raw}) // 返回文件的绝对路径（不含 exe 所在目录）
}
