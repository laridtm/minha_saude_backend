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
	ID           string `json:"id" bson:"_id"`
	Date         time.Time
	Hospital     string
	Professional string
	Observation  string
	UserId       string `json:"-"`
	Type         MedicalRecordType
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
