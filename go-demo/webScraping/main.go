package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {

	fmt.Println("Парсинг")
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Visited %s\n", r.Request.URL)
	})
	c.OnHTML("article.items .item", func(e *colly.HTMLElement) {
		// Извлекаем название вакансии
		title := e.ChildText(".name")
		// Извлекаем название компании
		//company := e.ChildText(".vacancy-serp-item__meta-info-company")
		// Извлекаем зарплату, если она указана
		//salary := e.ChildText(".vacancy-serp__vacancy-compensation")

		fmt.Printf("Сотрудник: %s\n", title)
	})
	c.Visit("https://sfedu.ru/www/stat_pages22.show?p=UNI/N11899/P")
}
