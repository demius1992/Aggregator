package repository

import (
	"Aggregator/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CottagesPostgres struct {
	db *pgxpool.Pool
}

func (c *CottagesPostgres) GetAllHouses(page int) ([]models.Cottages, error) {
	sum := make([]models.Cottages, 0)
	query := fmt.Sprintf(
		"SELECT title, map_marker, price, inf, description, link, typ\n"+
			"FROM %s\n"+
			"WHERE typ = 'house'\n"+
			"LIMIT 20 OFFSET $1", Cottage)

	if page > 1 {
		page = page*20 - 1
	}

	rows, err := c.db.Query(context.Background(), query, page)
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

func (c *CottagesPostgres) GetAllAreas(page int) ([]models.Areas, error) {
	sum := make([]models.Areas, 0)
	query := fmt.Sprintf(
		"SELECT title, map_marker, price, inf, description, link, typ\n"+
			"FROM %s\n"+
			"WHERE typ = 'area'\n"+
			"LIMIT 20 OFFSET $1", Cottage)

	if page > 1 {
		page = page*20 - 1
	}

	rows, err := c.db.Query(context.Background(), query, page)
	if err != nil {
		return sum, err
	}
	for rows.Next() {
		var summary models.Areas
		err := rows.Scan(&summary.Title, &summary.MapMarker, &summary.Price, &summary.Info,
			&summary.Description, &summary.Link, &summary.Typ)
		if err != nil {
			return sum, err
		}
		sum = append(sum, summary)
	}
	return sum, nil
}

func (c *CottagesPostgres) GetAllDacha(page int) ([]models.Cottages, error) {
	sum := make([]models.Cottages, 0)
	query := fmt.Sprintf(
		"SELECT title, map_marker, price, inf, description, link, typ\n"+
			"FROM %s\n"+
			"WHERE typ = 'dacha'\n"+
			"LIMIT 20 OFFSET $1", Cottage)

	if page > 1 {
		page = page*20 - 1
	}

	rows, err := c.db.Query(context.Background(), query, page)
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

func NewCottagesPostgres(db *pgxpool.Pool) *CottagesPostgres {
	return &CottagesPostgres{db: db}
}
