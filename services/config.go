package services

import (
	"ars_projekat/model"
)

type ConfigurationService struct {
	repo model.ConfigurationRepository
}

func NewConfigurationService(repo model.ConfigurationRepository) ConfigurationService {
	return ConfigurationService{
		repo: repo,
	}
}

func (s ConfigurationService) Add(config model.Configuration) {
	s.repo.Add(config)
}

func (s ConfigurationService) Get(name string, version int) (model.Configuration, error) {
	return s.repo.Get(name, version)
}
