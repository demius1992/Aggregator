package service

import (
	"Aggregator/models"
	"Aggregator/repository"
)

type FlatsService struct {
	repo repository.Flats
}

func (f *FlatsService) GetAllFlats(page int) ([]models.Flats, error) {
	return f.repo.GetAllFlats(page)
}

func (f *FlatsService) GetOneRoomFlats(page int) ([]models.Flats, error) {
	return f.repo.GetOneRoomFlats(page)
}

func (f *FlatsService) GetTwoRoomFlats(page int) ([]models.Flats, error) {
	return f.repo.GetTwoRoomFlats(page)
}

func (f *FlatsService) GetThreeRoomFlats(page int) ([]models.Flats, error) {
	return f.repo.GetThreeRoomFlats(page)
}

func (f *FlatsService) GetMoreThanThreeRoomFlats(page int) ([]models.Flats, error) {
	return f.repo.GetMoreThanThreeRoomFlats(page)
}

func NewFlatsService(repo repository.Flats) *FlatsService {
	return &FlatsService{repo: repo}
}
