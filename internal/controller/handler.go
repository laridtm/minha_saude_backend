package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/laridtm/minha_saude/internal/services"
)

type handler struct {
	profileService services.ProfileService
	historyService services.HistoryService
}

func NewHandler(
	profileService services.ProfileService,
	historyService services.HistoryService,
) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	h := &handler{
		profileService: profileService,
		historyService: historyService,
	}

	engine := gin.New()
	engine.Use(LogMiddleware())

	engine.GET("/ping", h.Ping)
	engine.POST("/profile", h.CreateProfile)
	engine.GET("/profile/:user", h.GetProfile)

	engine.POST("/medical-record/:user", h.CreateMedicalRecord)
	engine.GET("/medical-record/:user", h.ListAllMedicalRecords)
	engine.PUT("/medical-record/:user/:recordId", h.UpdateMedicalRecord)
	engine.DELETE("/medical-record/:recordId", h.RemoveMedicalRecord)

	return engine
}

func (h *handler) abortWithError(c *gin.Context, code int, err error) {
	log.Println(err.Error())
	c.AbortWithError(code, err)
}
