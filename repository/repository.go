package repository

import (
	"Aggregator/models"
	"github.com/jackc/pgx/v4/pgxpool"
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

type Repository struct {
	Flats
	Cottages
	FlatsForRent
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Flats:        NewFlatsPostgres(db),
		Cottages:     NewCottagesPostgres(db),
		FlatsForRent: NewFlatsForRentPostgres(db),
	}
}
