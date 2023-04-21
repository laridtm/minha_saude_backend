package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/laridtm/minha_saude/internal/model"
)

func (h *handler) CreateReminder(c *gin.Context) {
	userId := c.Param("user")
	if strings.TrimSpace(userId) == "" {
		h.abortWithError(c, http.StatusBadRequest, errors.New("missing user param"))
	}

	var reminder model.Reminder

	err := c.BindJSON(&reminder)
	if err != nil {
		h.abortWithError(c, http.StatusBadRequest, err)
		return
	}

	reminder.UserId = userId

	err = h.reminderService.CreateReminder(reminder)
	if err != nil {
		h.abortWithError(c, http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusCreated, "Reminder created")
}

func (h *handler) ListAllReminders(c *gin.Context) {
	userId := c.Param("user")
	if strings.TrimSpace(userId) == "" {
		h.abortWithError(c, http.StatusBadRequest, errors.New("missing user param"))
	}

	reminders, err := h.reminderService.GetAll(userId)
	if err != nil {
		h.abortWithError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, reminders)
}

func (h *handler) UpdateReminder(c *gin.Context) {
	reminderId := c.Param("reminderId")
	if strings.TrimSpace(reminderId) == "" {
		h.abortWithError(c, http.StatusBadRequest, errors.New("missing reminderId"))
	}

	userId := c.Param("user")
	if strings.TrimSpace(userId) == "" {
		h.abortWithError(c, http.StatusBadRequest, errors.New("missing user param"))
	}

	var reminder model.Reminder

	err := c.BindJSON(&reminder)
	if err != nil {
		h.abortWithError(c, http.StatusBadRequest, err)
		return
	}

	reminder.ID = reminderId
	reminder.UserId = userId

	if err := h.reminderService.UpdateReminder(reminderId, reminder); err != nil {
		h.abortWithError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, reminder)
}

func (h *handler) RemoveReminder(c *gin.Context) {
	reminderId := c.Param("reminderId")
	if strings.TrimSpace(reminderId) == "" {
		h.abortWithError(c, http.StatusBadRequest, errors.New("missing reminderId"))
	}

	err := h.reminderService.DeleteReminder(reminderId)
	if err != nil {
		h.abortWithError(c, http.StatusInternalServerError, err)
	}

	c.String(http.StatusOK, "Reminder removed")
}
