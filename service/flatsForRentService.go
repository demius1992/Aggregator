package service

import (
	"Aggregator/models"
	"Aggregator/repository"
)

type FlatsForRentService struct {
	repo repository.FlatsForRent
}

func (f *FlatsForRentService) GetFlatsForDay(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FlatsForRentService) GetFlatsLongRent(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FlatsForRentService) GetCottagesForDay(page int) ([]models.Cottages, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FlatsForRentService) GetCottagesLongRent(page int) ([]models.Cottages, error) {
	//TODO implement me
	panic("implement me")
}

func NewFlatsForRentService(repo repository.FlatsForRent) *FlatsForRentService {
	return &FlatsForRentService{repo: repo}
}
