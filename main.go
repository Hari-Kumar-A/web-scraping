package main

import (
	"fmt"

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
	var products[] Product
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
 
	//once job done, print out the array
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
		fmt.Println()

		// Print each product
	for _, product := range products {
		fmt.Println("Product Details:") 
		fmt.Println("  Name: ", product.Name) 
		fmt.Println("  Price: ", product.Price)
		fmt.Println()
	}
	})
	
	c.Visit("https://www.scrapingcourse.com/ecommerce")

	

}
