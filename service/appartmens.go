package service

import (
	"Aggregator/models"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func CrawlDom() {
	allFlats := make([]models.Flats, 0)

	c := colly.NewCollector(colly.AllowedDomains("domovita.by", "https://domovita.by"))
	c.OnHTML(".found_full", func(e *colly.HTMLElement) {
		attr := e.ChildText("a")
		mapMarker := e.ChildText(".gr.mb-5.fs-12")
		price := strings.Split(e.ChildText(".price.dropdown-toggle"), ".")[0]
		info := e.ChildText(".col-md-7.col-lg-9.black.s-bold.fs-14.sm-mb-10")
		description := e.ChildText("p")
		link := e.ChildAttr("a", "href")

		flats := models.Flats{
			Title:       attr,
			MapMarker:   mapMarker,
			Price:       price,
			Info:        info,
			Description: description,
			Link:        link,
		}
		allFlats = append(allFlats, flats)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	for i := 1; i > 0; i++ {
		URL := fmt.Sprintf("https://domovita.by/vitebsk/flats/sale?page=%d", i)
		resp, err := http.Get(URL)
		if err != nil {
			log.Println("out of pages")
			return
		}
		defer resp.Body.Close()
		err = c.Visit(URL)
		if err != nil {
			return
		}
	}
	writeJSON(allFlats)
}

func writeJSON(data []models.Flats) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}
	_ = ioutil.WriteFile("flats.json", file, 0644)
}

//func funcName() {
//
//	links := make([]string, 0)
//
//	flatCollector := colly.NewCollector()
//	c := colly.NewCollector(
//		colly.AllowedDomains("domovita.by", "https://domovita.by"),
//		// Cache responses to prevent multiple download of pages
//		// even if the collector is restarted
//		colly.CacheDir("./domovita_cache"),
//	)
//
//	c.OnHTML(".found_full", func(e *colly.HTMLElement) {
//		link := e.ChildAttr("a", "href")
//		c.Visit(e.Request.AbsoluteURL(link))
//	})
//
//	c.OnRequest(func(r *colly.Request) {
//		log.Println("visiting", r.URL.String())
//		links = append(links, r.URL.String())
//	})
//
//	c.Visit("https://domovita.by/vitebsk/flats/sale")
//
//	fmt.Println(links)
//
//	for _, v := range links {
//		flatCollector.OnHTML(".content-container", func(e *colly.HTMLElement) {
//			fmt.Println(e.ChildText("p"))
//		})
//		flatCollector.Visit(v)
//	}
//}
