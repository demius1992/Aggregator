package jobs

import (
	"Aggregator/models"
	"Aggregator/repository"
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
	"strings"
)

func flatDayRentCollector(db *pgxpool.Pool) error {
	c := colly.NewCollector(
		colly.AllowedDomains("domovita.by", "https://domovita.by"),
		colly.CacheDir("https://domovita.by"))
	c.OnHTML(".item-cell", func(e *colly.HTMLElement) {
		mapMarker := e.ChildText("div.fs-12.gr")
		price := e.ChildText(".pop_desc__price-orange") + " " +
			e.ChildText("span.fs-12.gr")
		info := e.ChildText(".text")
		description := e.ChildText("a")
		link := e.ChildAttr("a", "href")
		rooms, err := strconv.Atoi(strings.Split(description, "-")[0])
		if err != nil {
			log.Printf("failed to int convertation %s", err.Error())
		}

		flats := models.Flats{
			Title:       "",
			MapMarker:   mapMarker,
			Price:       price,
			Info:        info,
			Description: description,
			Link:        link,
			Rooms:       rooms,
			Typ:         "flatDayRent",
		}
		if err = writeDataDayRent(flats, db); err != nil {
			logrus.Printf("failed to write data %s", err.Error())
		}
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	for i := 1; i > 0; i++ {
		URL := fmt.Sprintf("https://domovita.by/vitebsk/flats-for-day/rent?page=%d", i)
		if err := c.Visit(URL); err != nil {
			break
		}
	}
	return nil
}

func writeDataDayRent(data models.Flats, db *pgxpool.Pool) error {
	query := fmt.Sprintf("INSERT INTO %s (title, map_marker, price, inf, description, link, rooms, typ) values ($1, $2, $3, $4, $5, $6, $7, $8)", repository.FlatForRent)
	_, err := db.Exec(context.Background(), query, data.Title, data.MapMarker, data.Price, data.Info, data.Description, data.Link, data.Rooms, data.Typ)
	if err != nil {
		logrus.Printf("failed to write data %s", err.Error())
	}
	return nil
}
