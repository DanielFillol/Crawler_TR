package Crawler

import (
	"github.com/tebeka/selenium"
)

const (
	webSiteTR              = "https://www.livrariart.com.br/#&search-term="
	bookOpenLink           = "//*[@id=\"smarthint-search-products\"]/section/div[2]/div/div/ul/div/li/div/div/div[1]/div[3]/h3/a"
	searchResultTR         = "//*[@id=\"smarthint-search-result-message\"]"
	notFoundTR             = "NENHUM RESULTADO"
	productSpecificationTR = "//*[@id=\"caracteristicas\"]/div[2]/div"
)

type Book struct {
	SearchName    string
	ISBN          string
	AvailableDate string
	Pages         string
	PubYear       string
	Link          string
	Error         string
}

func Craw(driver selenium.WebDriver, bookName string) (Book, error) {
	bookURL := webSiteTR + bookName
	driver.Get(bookURL)

	found, err := getSearchResult(driver, searchResultTR, notFoundTR)
	if err != nil {
		return Book{}, err
	}

	if found {
		book, err := getBook(driver, bookName)
		if err != nil {
			return Book{}, err
		}
		return book, nil
	}

	return Book{SearchName: bookName, Error: "Book not Found"}, nil
}
