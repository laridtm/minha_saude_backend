package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "PONG")
}
