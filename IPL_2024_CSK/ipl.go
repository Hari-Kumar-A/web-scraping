package main

import (
    "fmt"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/gocolly/colly"
)

type Match struct {
    Title   string `json:"title"`
    Stadium string `json:"stadium"`
    Result  string `json:"result"`
}

func main() {
    var Matches []Match
 
    r := gin.Default()

    // Enabling CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://127.0.0.1:5500"},  
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},  
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    }))
 
    r.GET("/matches", func(c *gin.Context) {
        if len(Matches) == 0 {
            c.JSON(500, gin.H{"error": "No match data available"})
            return
        }
        c.JSON(200, Matches)
    })
 
    go scrapeMatches(&Matches)

    // Run the Gin web server
    r.Run(":8080")
}

func scrapeMatches(Matches *[]Match) {
    c := colly.NewCollector(
        colly.AllowedDomains("www.cricbuzz.com"),
    )

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Colly Web-Scraper visiting: ", r.URL)
    })

    c.OnError(func(r *colly.Response, err error) {
        fmt.Println("Error: ", err)
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Page visited: ", r.Request.URL)
    })

    c.OnHTML(".cb-series-matches", func(e *colly.HTMLElement) {
        match := Match{}
        match.Title = e.ChildText(".cb-col-60.cb-col.cb-srs-mtchs-tm a span")
        match.Stadium = e.ChildText(".text-gray")
        match.Result = e.ChildText(".cb-text-complete")
        *Matches = append(*Matches, match)
    })

    c.OnScraped(func(r *colly.Response) {
        fmt.Println(r.Request.URL, " scraped!")
    })

    c.Visit("https://www.cricbuzz.com/cricket-series/7607/indian-premier-league-2024/matches")
}
