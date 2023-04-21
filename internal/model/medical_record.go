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
	Date         time.Time         `json:"date" bson:"date"`
	Hospital     string            `json:"hospital" bson:"hospital"`
	Professional string            `json:"professional" bson:"professional"`
	Name         string            `json:"name" bson:"name"`
	Observation  string            `json:"observation" bson:"observation"`
	UserId       string            `json:"-" bson:"userId"`
	Type         MedicalRecordType `json:"type" bson:"type"`
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
