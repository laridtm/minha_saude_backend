package controller

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/laridtm/minha_saude/internal/model"
)

func (h *handler) CreateProfile(c *gin.Context) {
	var profile model.Profile

	err := c.BindJSON(&profile)
	if err != nil {
		h.abortWithError(c, http.StatusBadRequest, err)
		return
	}

	log.Printf("profile recebido foi: [%+v]", profile)

	err = h.profileService.CreateProfile(profile)
	if err != nil {
		h.abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusCreated, "Profile created")
}

func (h *handler) GetProfile(c *gin.Context) {
	userId := c.Param("user")
	if strings.TrimSpace(userId) == "" {
		h.abortWithError(c, http.StatusBadRequest, errors.New("missing user param"))
	}

	profile, err := h.profileService.GetProfile(userId)
	if err != nil {
		if err.Error() == "profile does not exist" {
			h.abortWithError(c, http.StatusNotFound, err)
			return
		}
		h.abortWithError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, profile)
}
