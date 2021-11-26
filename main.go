package main

import (
	"net/http"
	"os"
	"scrapper_echo/scrapper"

	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func main() {
	e := echo.New()
	e.GET("/", home)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

func handleHome(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func home(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := c.FormValue("term")
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
	// scrapper.Scrape("python")
}
