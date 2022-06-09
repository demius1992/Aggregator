package repository

import (
	"Aggregator/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type FlatsPostgres struct {
	db *pgxpool.Pool
}

func (r *FlatsPostgres) GetOneRoomFlats(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func (r *FlatsPostgres) GetTwoRoomFlats(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func (r *FlatsPostgres) GetThreeRoomFlats(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func (r *FlatsPostgres) GetMoreThanThreeRoomFlats(page int) ([]models.Flats, error) {
	//TODO implement me
	panic("implement me")
}

func (r *FlatsPostgres) GetAllFlats(page int) ([]models.Flats, error) {
	flats := make([]models.Flats, 0)
	query := fmt.Sprintf("SELECT * FROM %s", "") //TODO

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var flt models.Flats
		err := rows.Scan("") //TODO implement me
		if err != nil {
			return nil, err
		}

		flats = append(flats, flt)
	}
	return flats, nil
}

func NewFlatsPostgres(db *pgxpool.Pool) *FlatsPostgres {
	return &FlatsPostgres{db: db}
}
