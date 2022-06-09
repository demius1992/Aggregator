package service

import (
	"Aggregator/models"
	"Aggregator/repository"
)

type FlatsService struct {
	repo repository.Flats
}

func (f *FlatsService) GetAllFlats(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FlatsService) GetOneRoomFlats(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FlatsService) GetTwoRoomFlats(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FlatsService) GetThreeRoomFlats(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FlatsService) GetMoreThanThreeRoomFlats(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func NewFlatsService(repo repository.Flats) *FlatsService {
	return &FlatsService{repo: repo}
}
