package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	//Methods Tests

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visitando o site ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Algo deu errado", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Pagina visitada: ", r.Request.URL)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		fmt.Printf("%v", e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, "Scraped")
	})

	c.Visit("https://br.wikipsdasdedia.org/wiki/Degemer")

}
