package model

import "errors"

type ReminderType string

const (
	EveryDay       ReminderType = "everyDay"
	Once           ReminderType = "once"
	MondayToFriday ReminderType = "mondayToFriday"
	Weekends       ReminderType = "weekends"
)

type Reminder struct {
	ID     string       `json:"id" bson:"_id"`
	Name   string       `json:"name"`
	Time   string       `json:"time"`
	UserId string       `json:"-"`
	Type   ReminderType `json:"type"`
}

func (rt ReminderType) Validate() error {
	switch rt {
	case EveryDay:
	case Once:
	case MondayToFriday:
	case Weekends:
	default:
		return errors.New("Tipo de lembrete invalido")
	}

	return nil
}
