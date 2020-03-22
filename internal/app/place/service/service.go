package service

import (
	"github.com/rockethelper/call-handler/internal/app/model"
	repo "github.com/rockethelper/call-handler/internal/app/place/repository"
)

type Service struct {
	Repository *repo.Repository
}

func New(repository *repo.Repository) *Service {
	return &Service{Repository: repository}
}

func (s Service) FindMatchingGermanAddressInformationFor(attr string, searchValue string) (model.GermanAddressInformation, error) {
	return s.Repository.FindMatchingGermanAddressInformationFor(attr, searchValue)
}
