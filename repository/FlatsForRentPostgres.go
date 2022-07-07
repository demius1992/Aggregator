package repository

import (
	"Aggregator/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type FlatsForRentPostgres struct {
	db *pgxpool.Pool
}

func (f *FlatsForRentPostgres) GetFlatsForDay(page int) ([]models.Flats, error) {
	sum := make([]models.Flats, 0)
	query := fmt.Sprintf(
		"SELECT title, map_marker, price, inf, description, link, rooms, typ\n"+
			"FROM %s\n"+
			"WHERE typ = 'flatDayRent'\n"+
			"LIMIT 20 OFFSET $1", FlatForRent)

	if page > 1 {
		page = page*20 - 1
	}

	rows, err := f.db.Query(context.Background(), query, page)
	if err != nil {
		return sum, err
	}
	for rows.Next() {
		var summary models.Flats
		err := rows.Scan(&summary.Title, &summary.MapMarker, &summary.Price, &summary.Info,
			&summary.Description, &summary.Link, &summary.Rooms, &summary.Typ)
		if err != nil {
			return sum, err
		}
		sum = append(sum, summary)
	}
	return sum, nil
}

func (f *FlatsForRentPostgres) GetFlatsLongRent(page int) ([]models.Flats, error) {
	sum := make([]models.Flats, 0)
	query := fmt.Sprintf(
		"SELECT title, map_marker, price, inf, description, link, rooms, typ\n"+
			"FROM %s\n"+
			"WHERE typ = 'flatLongRent'\n"+
			"LIMIT 20 OFFSET $1", FlatForRent)

	if page > 1 {
		page = page*20 - 1
	}

	rows, err := f.db.Query(context.Background(), query, page)
	if err != nil {
		return sum, err
	}
	for rows.Next() {
		var summary models.Flats
		err := rows.Scan(&summary.Title, &summary.MapMarker, &summary.Price, &summary.Info,
			&summary.Description, &summary.Link, &summary.Rooms, &summary.Typ)
		if err != nil {
			return sum, err
		}
		sum = append(sum, summary)
	}
	return sum, nil
}

func (f *FlatsForRentPostgres) GetCottagesForDay(page int) ([]models.Cottages, error) {
	sum := make([]models.Cottages, 0)
	query := fmt.Sprintf(
		"SELECT title, map_marker, price, inf, description, link, typ\n"+
			"FROM %s\n"+
			"WHERE typ = 'cottageDayRent'\n"+
			"LIMIT 20 OFFSET $1", FlatForRent)

	if page > 1 {
		page = page*20 - 1
	}

	rows, err := f.db.Query(context.Background(), query, page)
	if err != nil {
		return sum, err
	}
	for rows.Next() {
		var summary models.Cottages
		err := rows.Scan(&summary.Title, &summary.MapMarker, &summary.Price, &summary.Info,
			&summary.Description, &summary.Link, &summary.Typ)
		if err != nil {
			return sum, err
		}
		sum = append(sum, summary)
	}
	return sum, nil
}

func (f *FlatsForRentPostgres) GetCottagesLongRent(page int) ([]models.Cottages, error) {
	sum := make([]models.Cottages, 0)
	query := fmt.Sprintf(
		"SELECT title, map_marker, price, inf, description, link, typ\n"+
			"FROM %s\n"+
			"WHERE typ = 'cottageLongRent'\n"+
			"LIMIT 20 OFFSET $1", FlatForRent)

	if page > 1 {
		page = page*20 - 1
	}

	rows, err := f.db.Query(context.Background(), query, page)
	if err != nil {
		return sum, err
	}
	for rows.Next() {
		var summary models.Cottages
		err := rows.Scan(&summary.Title, &summary.MapMarker, &summary.Price, &summary.Info,
			&summary.Description, &summary.Link, &summary.Typ)
		if err != nil {
			return sum, err
		}
		sum = append(sum, summary)
	}
	return sum, nil
}

func NewFlatsForRentPostgres(db *pgxpool.Pool) *FlatsForRentPostgres {
	return &FlatsForRentPostgres{db: db}
}
