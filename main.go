package main

import (
	"fmt"
	"sync"
	"github.com/gocolly/colly" //import Colly library :)
)


type Product struct{
	Url string
	Image string
	Name string
	Price string
}


func main() {

	//Our logic starts..
	var products[] Product //array to store Product elements
	var visitedUrls sync.Map //hashmap type to store visited urls or not
	fmt.Println("Hello, World!")
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
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
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		 product:=Product{}
		 product.Url=e.ChildAttr("a", "href")
		 product.Image=e.ChildAttr("img","src")
		 product.Name=e.ChildText(".product-name")
		 product.Price=e.ChildText(".price")

		 //insert our product to products
		 products = append(products, product)
	})

	// Pagination functionality
    c.OnHTML("a.next", func(e *colly.HTMLElement) {
 
        nextPageUrl := e.Attr("href") //get the attribute stored in href = "/page/3" aisa

        // check if the nextPageUrl URL has been visited
	
        if _, flag := visitedUrls.Load(nextPageUrl); !flag { //if URL not visited
            fmt.Println("scraping:", nextPageUrl) 
            visitedUrls.Store(nextPageUrl, struct{}{}) //mark as visited
            e.Request.Visit(nextPageUrl) //visit this new URL page
        }
    })
 
	//once job done, print out the array
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
		fmt.Println()

		// Print each product
	for i, product := range products {
		fmt.Println("Product Details:", i+1) 
		fmt.Println("  Name: ", product.Name) 
		fmt.Println("  Price: ", product.Price)
		fmt.Println()
	}

	fmt.Println("Total Products:", len(products))
	fmt.Println()
	})
	
	c.Visit("https://www.scrapingcourse.com/ecommerce")

	

}
