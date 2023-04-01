package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laridtm/minha_saude/internal/model"
)

func (h *handler) CreateProfile(c *gin.Context) {
	var profile model.Profile

	err := c.BindJSON(&profile)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Printf("profile recebido foi: [%+v]", profile)

	err = h.profileService.CreateProfile(profile)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusOK, "Profile created")
}
