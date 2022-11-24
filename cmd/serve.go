package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"
	scrapper "github.com/vic30004/go-scrapper-backend/internal/scrapper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the application over HTTP",
	Long:  `Starts an HTTP server on port 8080 to serve up the scraper application`,
	Run: func(cmd *cobra.Command, args []string) {
		app := negroni.Classic()

		c := colly.NewCollector(
			colly.AllowedDomains("amazon.com", "bestbuy.com"),
		)

		scrapperController := scrapper.New(c)
		router := mux.NewRouter()

		scrapperController.Register(router)

		cor := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
		})
		app.Use(cor)

		app.UseHandler(router)

		fmt.Println("Listening on port 8080")
		log.Fatal(http.ListenAndServe(":8080", app))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
