package main

import (
	"fmt"
	"github.com/Darklabel91/Crawler_TR/CSV"
	"github.com/Darklabel91/Crawler_TR/Crawler"
	"github.com/tebeka/selenium"
	"strconv"
)

const (
	webSiteTR      = "https://www.livrariart.com.br/#&search-term="
	bookPriceTR    = "//*[@id=\"smarthint-search-products\"]/section/div[2]/div/div/ul/div/li/div/div/div[1]/div[4]/div/span[1]"
	searchResultTR = "//*[@id=\"smarthint-search-result-message\"]"
	notFoundTR     = "Nenhum resultado para litigância climática no brasil. Pessoas que buscaram litigância climática no brasil também buscaram ou compraram os produtos abaixo:"
)

func main() {
	bookNames, err := CSV.ReadCsvFile("/Users/danielfillol/Desktop/Planilha sem título - Página1.csv")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(bookNames); i++ {
		driver, err := Crawler.SeleniumWebDriver()
		if err != nil {
			fmt.Println(err)
		}

		bookSearch := webSiteTR + bookNames[i]

		err = driver.Get(bookSearch)
		if err != nil {
			fmt.Println(err)
		}

		bookPrice, err := driver.FindElements(selenium.ByXPATH, bookPriceTR)
		if err != nil {
			fmt.Println(err)
		}

		resultElem, err := driver.FindElements(selenium.ByXPATH, searchResultTR)
		if err != nil {
			fmt.Println(err)
		}

		if len(resultElem) != 0 {
			restultText, err := resultElem[0].Text()
			if err != nil {
				fmt.Println(err)
			}

			if restultText == notFoundTR {
				fmt.Println("No book found")
			} else {
				fmt.Println(strconv.Itoa(i), " - ", bookSearch, " - ", bookNames[i])
				fmt.Println(bookPrice[0].Text())
			}
		}

		err = driver.Close()
		if err != nil {
			fmt.Println(err)
		}
	}

}
