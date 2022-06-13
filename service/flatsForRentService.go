package service

import (
	"Aggregator/models"
	"Aggregator/repository"
)

type FlatsForRentService struct {
	repo repository.FlatsForRent
}

func (f *FlatsForRentService) GetFlatsForDay(page int) ([]models.Flats, error) {
	return f.repo.GetFlatsForDay(page)
}

func (f *FlatsForRentService) GetFlatsLongRent(page int) ([]models.Flats, error) {
	return f.repo.GetFlatsLongRent(page)
}

func (f *FlatsForRentService) GetCottagesForDay(page int) ([]models.Cottages, error) {
	return f.repo.GetCottagesForDay(page)
}

func (f *FlatsForRentService) GetCottagesLongRent(page int) ([]models.Cottages, error) {
	return f.repo.GetCottagesLongRent(page)
}

func NewFlatsForRentService(repo repository.FlatsForRent) *FlatsForRentService {
	return &FlatsForRentService{repo: repo}
}
