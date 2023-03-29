package http

import (
	"github.com/gin-gonic/gin"
)

type handler struct{}

func NewHandler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	h := &handler{}

	engine := gin.New()
	engine.Use(LogMiddleware())

	engine.GET("/ping", h.Ping)

	return engine
}
