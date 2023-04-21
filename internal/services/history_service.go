package services

import (
	"github.com/laridtm/minha_saude/internal/model"
	"github.com/laridtm/minha_saude/internal/repository/mongo"
)

type HistoryService struct {
	repository mongo.HistoryRepository
}

func NewHistoryService(repository mongo.HistoryRepository) HistoryService {
	return HistoryService{
		repository: repository,
	}
}

func (hs *HistoryService) CreateRecord(record model.MedicalRecord) error {
	if err := record.Type.Validate(); err != nil {
		return err
	}

	return hs.repository.Insert(record)
}

func (hs *HistoryService) GetAll(userId string, filter *model.MedicalRecordType) ([]model.MedicalRecord, error) {
	return hs.repository.FindAll(userId, filter)
}

func (hs *HistoryService) UpdateRecord(recordId string, record model.MedicalRecord) error {
	if err := record.Type.Validate(); err != nil {
		return err
	}

	return hs.repository.Update(recordId, record)
}

func (hs *HistoryService) DeleteRecord(recordId string) error {
	return hs.repository.Delete(recordId)
}
