package jobs

import (
	"Aggregator/models"
	"Aggregator/repository"
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"strings"
)

func areaCollector(db *pgxpool.Pool) error {
	c := colly.NewCollector(
		colly.AllowedDomains("domovita.by", "https://domovita.by"),
		colly.CacheDir("https"))
	c.OnHTML(".found_full", func(e *colly.HTMLElement) {
		attr := e.ChildText("a")
		mapMarker := e.ChildText(".gr.mb-5.fs-12")
		price := strings.Split(e.ChildText(".price.dropdown-toggle"), ".")[0]
		info := e.ChildText(".col-md-7.col-lg-9.black.s-bold.fs-14.sm-mb-10")
		description := e.ChildText("p")
		link := e.ChildAttr("a", "href")
		typ := "area"

		area := models.Areas{
			Title:       attr,
			MapMarker:   mapMarker,
			Price:       price,
			Info:        info,
			Description: description,
			Link:        link,
			Typ:         typ,
		}
		err := writeDataAr(area, db)
		if err != nil {
			logrus.Printf("failed to write data %s", err.Error())
		}
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	for i := 1; i > 0; i++ {
		URL := fmt.Sprintf("https://domovita.by/vitebsk/area/sale?page=%d", i)
		err := c.Visit(URL)
		if err != nil {
			if err.Error() != "Not Found" {
				logrus.Printf("an error uccured in areaCollector while visiting URL %s", err.Error())
			}
			break
		}
	}
	return nil
}

func writeDataAr(data models.Areas, db *pgxpool.Pool) error {
	query := fmt.Sprintf("INSERT INTO %s (title, map_marker, price, inf, description, link, typ) values ($1, $2, $3, $4, $5, $6, $7)", repository.Cottage)
	_, err := db.Exec(context.Background(), query, data.Title, data.MapMarker, data.Price, data.Info, data.Description, data.Link, data.Typ)
	if err != nil {
		logrus.Printf("failed to write data %s", err.Error())
	}
	return nil
}
