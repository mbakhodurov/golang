package services

import (
	"log"
	"wheather/internal/domain"
	"wheather/internal/repositories"
)

type WeatherService interface {
	Insert(dates domain.Weather) (int64, error)
	SelectNotConf() ([]domain.Weather, error)
	UpdateStatus(id int64, status int) error
}
type weatherService struct {
	repo repositories.WeatherRepo
}

func (s *weatherService) UpdateStatus(id int64, status int) error { // Implement the method
	return s.repo.UpdateStatus(id, status)
}

func NewWeatherService(repo repositories.WeatherRepo) WeatherService {
	return &weatherService{repo}
}

func (s weatherService) Insert(dates domain.Weather) (int64, error) {
	return s.repo.Insert(dates)
}

func (s weatherService) SelectNotConf() ([]domain.Weather, error) {
	items, err := s.repo.SelectNotConf()
	if err != nil {
		log.Fatalf("Error selecting weather services: %v", err)
	}
	return items, nil
}
