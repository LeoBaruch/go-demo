package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func QrcodesController(c *gin.Context) {
	content := c.Query("content")

	if content == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.Data(http.StatusOK, "image/png", png)
}
