package services

import (
	"github.com/laridtm/minha_saude/internal/model"
	"github.com/laridtm/minha_saude/internal/repository/mongo"
)

type ReminderService struct {
	repository mongo.ReminderRepository
}

func NewReminderService(repository mongo.ReminderRepository) ReminderService {
	return ReminderService{
		repository: repository,
	}
}

func (rs *ReminderService) CreateReminder(reminder model.Reminder) error {
	if err := reminder.Type.Validate(); err != nil {
		return err
	}

	return rs.repository.Insert(reminder)
}

func (rs *ReminderService) GetAll(userId string) ([]model.Reminder, error) {
	return rs.repository.FindAll(userId)
}

func (rs *ReminderService) UpdateReminder(reminderId string, reminder model.Reminder) error {
	if err := reminder.Type.Validate(); err != nil {
		return err
	}

	return rs.repository.Update(reminderId, reminder)
}

func (rs *ReminderService) DeleteReminder(reminderId string) error {
	return rs.repository.Delete(reminderId)
}
