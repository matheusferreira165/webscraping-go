package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type PokemonProduct struct {
	Url, Image, Name, Price string
}

func main() {

	var pokemonProducts []PokemonProduct

	c := colly.NewCollector()

	c.OnHTML("li.product", func(e *colly.HTMLElement) {

		pokemonProduct := PokemonProduct{}

		pokemonProduct.Url = e.ChildAttr("a", "href")
		pokemonProduct.Image = e.ChildAttr("img", "src")
		pokemonProduct.Name = e.ChildText("h2")
		pokemonProduct.Price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	c.Visit("https://scrapeme.live/shop/")

	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Falha para criar o arquivo CSV", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}

	writer.Write(headers)

	for _, pokemonProduct := range pokemonProducts {
		record := []string{
			pokemonProduct.Url,
			pokemonProduct.Image,
			pokemonProduct.Name,
			pokemonProduct.Price,
		}

		writer.Write(record)
	}
	defer writer.Flush()
}
