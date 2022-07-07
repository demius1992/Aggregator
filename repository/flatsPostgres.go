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
	sum := make([]models.Flats, 0)
	query := fmt.Sprintf(
		"SELECT title, map_marker, price, inf, description, link, rooms\n"+
			"FROM %s\n"+
			"WHERE rooms = 1\n"+
			"LIMIT 20 OFFSET $1", Flat)

	if page > 1 {
		page = page * 10
	}

	rows, err := r.db.Query(context.Background(), query, page)
	if err != nil {
		return sum, err
	}
	for rows.Next() {
		var summary models.Flats
		err := rows.Scan(&summary.Title, &summary.MapMarker, &summary.Price, &summary.Info,
			&summary.Description, &summary.Link, &summary.Rooms)
		if err != nil {
			return sum, err
		}
		sum = append(sum, summary)
	}
	return sum, nil
}

func (r *FlatsPostgres) GetTwoRoomFlats(page int) ([]models.Flats, error) {
	sum := make([]models.Flats, 0)
	query := fmt.Sprintf(
		"SELECT title, map_marker, price, inf, description, link, rooms\n"+
			"FROM %s\n"+
			"WHERE rooms = 2\n"+
			"LIMIT 20 OFFSET $1", Flat)

	if page > 1 {
		page = page*20 - 1
	}

	rows, err := r.db.Query(context.Background(), query, page)
	if err != nil {
		return sum, err
	}
	for rows.Next() {
		var summary models.Flats
		err := rows.Scan(&summary.Title, &summary.MapMarker, &summary.Price, &summary.Info,
			&summary.Description, &summary.Link, &summary.Rooms)
		if err != nil {
			return sum, err
		}
		sum = append(sum, summary)
	}
	return sum, nil
}

func (r *FlatsPostgres) GetThreeRoomFlats(page int) ([]models.Flats, error) {
	sum := make([]models.Flats, 0)
	query := fmt.Sprintf(
		"SELECT title, map_marker, price, inf, description, link, rooms\n"+
			"FROM %s\n"+
			"WHERE rooms = 3\n"+
			"LIMIT 20 OFFSET $1", Flat)

	if page > 1 {
		page = page*20 - 1
	}

	rows, err := r.db.Query(context.Background(), query, page)
	if err != nil {
		return sum, err
	}
	for rows.Next() {
		var summary models.Flats
		err := rows.Scan(&summary.Title, &summary.MapMarker, &summary.Price, &summary.Info,
			&summary.Description, &summary.Link, &summary.Rooms)
		if err != nil {
			return sum, err
		}
		sum = append(sum, summary)
	}
	return sum, nil
}

func (r *FlatsPostgres) GetMoreThanThreeRoomFlats(page int) ([]models.Flats, error) {
	sum := make([]models.Flats, 0)
	query := fmt.Sprintf(
		"SELECT title, map_marker, price, inf, description, link, rooms\n"+
			"FROM %s\n"+
			"WHERE rooms = 4\n"+
			"LIMIT 20 OFFSET $1", Flat)

	if page > 1 {
		page = page*20 - 1
	}

	rows, err := r.db.Query(context.Background(), query, page)
	if err != nil {
		return sum, err
	}
	for rows.Next() {
		var summary models.Flats
		err := rows.Scan(&summary.Title, &summary.MapMarker, &summary.Price, &summary.Info,
			&summary.Description, &summary.Link, &summary.Rooms)
		if err != nil {
			return sum, err
		}
		sum = append(sum, summary)
	}
	return sum, nil
}

func (r *FlatsPostgres) GetAllFlats(page int) ([]models.Flats, error) {
	sum := make([]models.Flats, 0)
	query := fmt.Sprintf(
		"SELECT title, map_marker, price, inf, description, link, rooms\n"+
			"FROM %s\n"+
			"LIMIT 20 OFFSET $1", Flat)

	if page > 1 {
		page = page*20 - 1
	}

	rows, err := r.db.Query(context.Background(), query, page)
	if err != nil {
		return sum, err
	}
	for rows.Next() {
		var summary models.Flats
		err := rows.Scan(&summary.Title, &summary.MapMarker, &summary.Price, &summary.Info,
			&summary.Description, &summary.Link, &summary.Rooms)
		if err != nil {
			return sum, err
		}
		sum = append(sum, summary)
	}
	return sum, nil
}

func NewFlatsPostgres(db *pgxpool.Pool) *FlatsPostgres {
	return &FlatsPostgres{db: db}
}
