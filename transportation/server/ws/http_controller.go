package ws

import "github.com/gin-gonic/gin"

func HttpController(c *gin.Context, hub *Hub) {
	serveWs(hub, c.Writer, c.Request)
}
