package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	fName := ("company-details.csv")
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"company title"})

	for i := 5; i <= 15; i++ {
		// Instantiate default collector
		c := colly.NewCollector()

		c.OnHTML("#job_listing_panel", func(e *colly.HTMLElement) {
			writer.Write([]string{
				e.ChildText(".company-name"),
			})
		})

		c.Visit(fmt.Sprintf("https://www.jobstreet.com.ph/en/job-search/job-vacancy/%d/", i))

	}
	log.Printf("Scraping finished, check file %q for results\n", fName)
}
