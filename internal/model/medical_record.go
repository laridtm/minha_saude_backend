package model

import (
	"errors"
	"time"
)

type MedicalRecordType string

const (
	Vaccine     MedicalRecordType = "vaccine"
	Appointment MedicalRecordType = "appointment"
	Exam        MedicalRecordType = "exam"
)

type MedicalRecord struct {
	ID           string            `json:"id" bson:"_id"`
	Date         time.Time         `json:"date"`
	Hospital     string            `json:"hospital"`
	Professional string            `json:"professional"`
	Name         string            `json:"name"`
	Observation  string            `json:"observation"`
	UserId       string            `json:"-"`
	Type         MedicalRecordType `json:"type"`
}

func (mrt MedicalRecordType) Validate() error {
	switch mrt {
	case Vaccine:
	case Appointment:
	case Exam:
	default:
		return errors.New("Tipo de registro medico invalido")
	}

	return nil
}
