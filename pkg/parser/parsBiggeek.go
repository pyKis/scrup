package parsBiggeek

import (
	"database/sql"
	"fmt"
	"log"

	"scrup/db"
	"scrup/models"
	"strings"

	"github.com/gocolly/colly/v2"
)


func ParsBiggeek(database *sql.DB) {
	c :=colly.NewCollector(
		colly.AllowedDomains("biggeek.ru"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36"),
	)

	c.OnHTML("div.prod-pagination", func(e *colly.HTMLElement) {

		page:= e.ChildAttrs("a.prod-pagination__item","href")
		for _, p := range page {
			nextVisited := "https://biggeek.ru" + p
			c.Visit(nextVisited)
			
		}
		
	})


var items []models.Item
    
	c.OnHTML("div.catalog-card", func(e *colly.HTMLElement)  {
		name := e.ChildAttr("img","alt")

		linc := e.ChildAttr("a","href")
		if linc != "" {
			linc = "https://biggeek.ru" + linc
		}

		price := e.ChildText("b")
		if price != "" {
			for _, r := range price {
				if r == ' ' {
					price = strings.Replace(price, " ", "", -1)
				}
				price = strings.Replace(price, "₽", "", -1)
			}
		}

		oldPrice := e.ChildText("span.old-price")
		if oldPrice != "" {
			for _, r := range oldPrice {
				if r == ' ' {
					oldPrice = strings.Replace(oldPrice, " ", "", -1)
				}
			}

		}
		item := models.Item{
			Name: name,
			Linc: linc,
			Price: price,
			OldPrice: oldPrice,
			
		}
			items = append(items, item)
			
		//discont := ((oldPrice - price) / oldPrice) * 100
		err := db.InsertItemBiggeek(database, item)
		if err != nil {
			log.Printf("Ошибка вставки элемента: %v\n", err)
		}
		
			
        
		
	
	})
	

	c.OnScraped(func(r *colly.Response) {
		log.Println("Парсинг завершен")
	})
	
	
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		
	})
	
	c.Visit("https://biggeek.ru/sale")
}
