package repository

import (
	"Aggregator/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CottagesPostgres struct {
	db *pgxpool.Pool
}

func (c *CottagesPostgres) GetAllHouses(page int) ([]models.Cottages, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CottagesPostgres) GetAllAreas(page int) ([]models.Areas, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CottagesPostgres) GetAllDacha(page int) ([]models.Cottages, error) {
	//TODO implement me
	panic("implement me")
}

func NewCottagesPostgres(db *pgxpool.Pool) *CottagesPostgres {
	return &CottagesPostgres{db: db}
}
