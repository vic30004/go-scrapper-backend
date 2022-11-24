package scrapper

import (
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

type Controller struct {
	Collector *colly.Collector
}

func New(collector *colly.Collector) *Controller {
	return &Controller{
		Collector: collector,
	}
}

func (c *Controller) Register(router *mux.Router) {
	router.HandleFunc("/", c.Scrape()).Methods(http.MethodGet)
}
