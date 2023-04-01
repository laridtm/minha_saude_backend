package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/laridtm/minha_saude/internal/services"
)

type handler struct {
	profileService services.ProfileService
}

func NewHandler(profileService services.ProfileService) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	h := &handler{
		profileService: profileService,
	}

	engine := gin.New()
	engine.Use(LogMiddleware())

	engine.GET("/ping", h.Ping)
	engine.POST("/profile", h.CreateProfile)

	return engine
}
