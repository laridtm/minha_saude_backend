package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/laridtm/minha_saude/internal/model"
)

func (h *handler) CreateMedicalRecord(c *gin.Context) {
	userId := c.Param("user")
	if strings.TrimSpace(userId) == "" {
		h.abortWithError(c, http.StatusBadRequest, errors.New("missing user param"))
	}

	var record model.MedicalRecord

	err := c.BindJSON(&record)
	if err != nil {
		h.abortWithError(c, http.StatusBadRequest, err)
		return
	}

	record.UserId = userId

	err = h.historyService.CreateRecord(record)
	if err != nil {
		h.abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusCreated, "Record created")
}

func (h *handler) UpdateMedicalRecord(c *gin.Context) {
	recordId := c.Param("recordId")
	if strings.TrimSpace(recordId) == "" {
		h.abortWithError(c, http.StatusBadRequest, errors.New("missing recordId"))
	}

	userId := c.Param("user")
	if strings.TrimSpace(userId) == "" {
		h.abortWithError(c, http.StatusBadRequest, errors.New("missing user param"))
	}

	var record model.MedicalRecord

	err := c.BindJSON(&record)
	if err != nil {
		h.abortWithError(c, http.StatusBadRequest, err)
		return
	}

	record.ID = recordId
	record.UserId = userId

	if err := h.historyService.UpdateRecord(recordId, record); err != nil {
		h.abortWithError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, record)
}

func (h *handler) ListAllMedicalRecords(c *gin.Context) {
	userId := c.Param("user")
	if strings.TrimSpace(userId) == "" {
		h.abortWithError(c, http.StatusBadRequest, errors.New("missing user param"))
	}

	records, err := h.historyService.GetAll(userId)
	if err != nil {
		h.abortWithError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, records)
}

func (h *handler) RemoveMedicalRecord(c *gin.Context) {
	recordId := c.Param("recordId")
	if strings.TrimSpace(recordId) == "" {
		h.abortWithError(c, http.StatusBadRequest, errors.New("missing recordId"))
	}

	err := h.historyService.DeleteRecord(recordId)
	if err != nil {
		h.abortWithError(c, http.StatusInternalServerError, err)
	}

	c.String(http.StatusCreated, "Record removed")
}
