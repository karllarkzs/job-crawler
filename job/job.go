package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	// Instantiate default collector

	fName := ("companysss.csv")
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
	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	c := colly.NewCollector()
	c.WithTransport(t)
	c.IgnoreRobotsTxt = true

	c.OnHTML("#job_listing_panel", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText(".company-name"),
		})
	})

	files, err := ioutil.ReadDir("./downloads")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(fmt.Sprintf("file:///Users/a-fis/Projects/js/go-crawler/job-crawler/job/downloads/%s", file.Name()))
		c.Visit(fmt.Sprintf("file:///Users/a-fis/Projects/js/go-crawler/job-crawler/job/downloads/%s", file.Name()))
	}

	log.Printf("Scraping finished, check file %q for results\n")
}
