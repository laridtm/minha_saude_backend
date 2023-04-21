package controller

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

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

	filter := model.NewFilter()
	filterType := c.Query("filter")
	if strings.TrimSpace(filterType) != "" {
		if err := model.MedicalRecordType(filterType).Validate(); err != nil {
			h.abortWithError(c, http.StatusBadRequest, errors.New("filter query paramter invalid"))
		}
		filter.Fields["type"] = filterType
	}

	size := c.Query("size")
	if strings.TrimSpace(size) != "" {
		intSize, err := strconv.Atoi(size)
		if err != nil {
			h.abortWithError(c, http.StatusBadRequest, errors.New("size query paramter invalid"))
		}
		filter.Size = intSize
	}

	recent := c.Query("recent")
	if strings.TrimSpace(recent) != "" {
		recentBool, err := strconv.ParseBool(recent)
		if err != nil {
			h.abortWithError(c, http.StatusBadRequest, errors.New("recent query paramter invalid"))
		}
		if recentBool {
			now := time.Now().UTC()
			filter.FromDate = &now
		}
	}

	records, err := h.historyService.GetAll(userId, filter)
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

	c.String(http.StatusOK, "Record removed")
}
