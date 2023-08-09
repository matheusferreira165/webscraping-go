package main

import (
	"github.com/gocolly/colly"
)

func main() {
	var pokemonProducts []PokemonProduct

	c := colly.NewCollector()

	c.Visit("https://scrapeme.live/shop/")

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}

		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

}
