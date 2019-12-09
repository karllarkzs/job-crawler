package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gocolly/colly"
)

type CompanyList struct {
	Name string `json:"name"`
}

func main() {
	var page int
	siteContainer := []CompanyList{}

	// Instantiate default collector

	//fName := ("job.csv")
	//file, err := os.Create(fName)
	//if err != nil {
	//log.Fatalf("Cannot create file %q: %s\n", fName, err)
	//return
	//}
	//defer file.Close()
	//writer := csv.NewWriter(file)
	//defer writer.Flush()

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	c := colly.NewCollector()
	c.WithTransport(t)
	c.IgnoreRobotsTxt = true

	c.OnHTML("#job_listing_panel", func(e *colly.HTMLElement) {
		e.ForEach("div.panel-body", func(_ int, e *colly.HTMLElement) {
			var name string

			name = e.ChildText(".company-name")

			fmt.Println("----------------------")
			fmt.Printf("Company Name: %s \n", name)
			fmt.Println("----------------------")

			site := CompanyList{}
			site.Name = name
			siteContainer = append(siteContainer, site)
		})

		page++

		if page == 55 {
			page = 58
		}

		if page > 58 {
			file, err := json.MarshalIndent(siteContainer, "", " ")
			if err != nil {
				panic("Not cool")
			}
			ioutil.WriteFile("jobstreet.json", file, 0644)
		} else {
			c.Visit(fmt.Sprintf("file:///Users/a-fis/Projects/js/go-crawler/job-crawler/job/downloads/jobVacancy%d.html", page))
		}
	})

	fmt.Println("visiting")
	site := fmt.Sprintf("file:///Users/a-fis/Projects/js/go-crawler/job-crawler/job/downloads/jobVacancy%d.html", page)

	c.Visit(site)
	c.Wait()
}
