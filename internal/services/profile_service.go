package services

import (
	"errors"
	"strings"

	"github.com/laridtm/minha_saude/internal/model"
	"github.com/laridtm/minha_saude/internal/repository/mongo"
)

type ProfileService struct {
	repository mongo.ProfileRepository
}

func NewProfileService(repository mongo.ProfileRepository) ProfileService {
	return ProfileService{
		repository: repository,
	}
}

func (ps *ProfileService) CreateProfile(profile model.Profile) error {
	if strings.TrimSpace(profile.CPF) == "" {
		return errors.New("CPF é obrigatório")
	}

	existedProfile, err := ps.repository.Find(profile.CPF)
	if err != nil {
		if err.Error() != "profile does not exist" {
			return err
		}
	}

	if existedProfile != nil {
		return ps.repository.Update(profile)
	}

	return ps.repository.Insert(profile)
}

func (ps *ProfileService) GetProfile(userId string) (*model.Profile, error) {
	return ps.repository.Find(userId)
}
