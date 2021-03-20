package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

type Movies struct {
	FilmName string `json:"filmName"`
	Rating   string `json:"rating"`
}

func main() {
	movies := make([]Movies, 0)

	c := colly.NewCollector(
		colly.AllowedDomains("imdb.com", "www.imdb.com"),
	)

	c.OnHTML(".lister-list tr", func(h *colly.HTMLElement) {
		var movie Movies

		filmName := h.DOM.Find(".titleColumn").Children().Text()
		rating := h.DOM.Find(".ratingColumn.imdbRating").Children().Text()

		movie.FilmName = filmName
		movie.Rating = rating

		movies = append(movies, movie)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.imdb.com/chart/top/")

	b, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("something wrong %+v \n", err)
	}

	err = ioutil.WriteFile("data.json", b, 0644)
	if err != nil {
		log.Fatalf("something wrong %+v \n", err)
	}
}
