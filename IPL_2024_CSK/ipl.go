package main

import (
	"fmt"
	// "sync"
	"github.com/gocolly/colly" //import Colly library :)
)

type Product struct {
	Url   string
	Image string
	Name  string
	Price string
}

type Match struct {
	Title    string 
	Stadium  string
	Result   string
}

func main() {

	//Our logic starts..
	var Matches []Match
	// var visitedUrls sync.Map //hashmap type to store visited urls or not
	fmt.Println("Hello, World!")
	c := colly.NewCollector(
		colly.AllowedDomains("www.cricbuzz.com"),
	)

	//before every request, yeh callback hoga
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Colly Web-Scraper visiting: ", r.URL)
	})

	//agar koi error aayega toh
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error: ", err)
	})

	//response ko track karke print karega
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	// OnHTML callback on matching with li.product
	c.OnHTML(".cb-series-matches", func(e *colly.HTMLElement) {
		match := Match{}
		match.Title = e.ChildText(".cb-col-60.cb-col.cb-srs-mtchs-tm a span") 
		match.Stadium = e.ChildText(".text-gray")
		match.Result = e.ChildText(".cb-text-complete")  
		Matches = append(Matches, match)
	})

	 

	//once job done, print out the array
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
		fmt.Println()

		// Print each product
		for i, match := range Matches {
			fmt.Println("Match No:", i+1)
			fmt.Println("  Title: ", match.Title) 
			fmt.Println("  Stadium: ", match.Stadium)
			fmt.Println("  Result: ", match.Result)
			fmt.Println()
		}

		fmt.Println("Total Matches:", len(Matches))
		fmt.Println()
	})

	c.Visit("https://www.cricbuzz.com/cricket-series/7607/indian-premier-league-2024/matches")

}
