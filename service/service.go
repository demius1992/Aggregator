package service

import (
	"Aggregator/models"
	"Aggregator/repository"
)

type Flats interface {
	GetAllFlats(page int) ([]models.Flats, error)
	GetOneRoomFlats(page int) ([]models.Flats, error)
	GetTwoRoomFlats(page int) ([]models.Flats, error)
	GetThreeRoomFlats(page int) ([]models.Flats, error)
	GetMoreThanThreeRoomFlats(page int) ([]models.Flats, error)
}

type Cottages interface {
	GetAllHouses(page int) ([]models.Cottages, error)
	GetAllAreas(page int) ([]models.Areas, error)
	GetAllDacha(page int) ([]models.Cottages, error)
}

type FlatsForRent interface {
	GetFlatsForDay(page int) ([]models.Flats, error)
	GetFlatsLongRent(page int) ([]models.Flats, error)
	GetCottagesForDay(page int) ([]models.Cottages, error)
	GetCottagesLongRent(page int) ([]models.Cottages, error)
}

type Service struct {
	Flats
	Cottages
	FlatsForRent
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Flats:        NewFlatsService(repos.Flats),
		Cottages:     NewCottagesService(repos.Cottages),
		FlatsForRent: NewFlatsForRentService(repos.FlatsForRent),
	}
}
