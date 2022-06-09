package repository

import (
	"Aggregator/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type FlatsForRentPostgres struct {
	db *pgxpool.Pool
}

func (f *FlatsForRentPostgres) GetFlatsForDay(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FlatsForRentPostgres) GetFlatsLongRent(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FlatsForRentPostgres) GetCottagesForDay(page int) ([]models.Cottages, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FlatsForRentPostgres) GetCottagesLongRent(page int) ([]models.Cottages, error) {
	//TODO implement me
	panic("implement me")
}

func NewFlatsForRentPostgres(db *pgxpool.Pool) *FlatsForRentPostgres {
	return &FlatsForRentPostgres{db: db}
}
