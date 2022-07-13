package Crawler

import "github.com/tebeka/selenium"

const (
	webSiteTR              = "https://www.livrariart.com.br/#&search-term="
	bookOpenLink           = "//*[@id=\"smarthint-search-products\"]/section/div[2]/div/div/ul/div/li/div/div/div[1]/div[3]/h3/a"
	searchResultTR         = "//*[@id=\"smarthint-search-result-message\"]"
	notFoundTR             = "NENHUM RESULTADO PARA DIREITOS FUNDAMENTAIS E JURISDIÇÃO CONSTITUCIONAL: ANÁLISE, CRÍTICA E CONTRIBUIÇÕES. PESSOAS QUE BUSCARAM DIREITOS FUNDAMENTAIS E JURISDIÇÃO CONSTITUCIONAL: ANÁLISE, CRÍTICA E CONTRIBUIÇÕES TAMBÉM BUSCARAM OU COMPRARAM OS PRODUTOS ABAIXO:"
	productSpecificationTR = "//*[@id=\"caracteristicas\"]/div[2]/div"
)

type Book struct {
	SearchName    string
	ISBN          string
	AvailableDate string
	Pages         string
	PubYear       string
	Link          string
}

func Craw(driver selenium.WebDriver, bookName string) (Book, error) {
	bookSearch := webSiteTR + bookName

	err := driver.Get(bookSearch)
	if err != nil {
		//This error happens very often in this website
		//println("something wrong with loading status craw")
	}

	foundBooks, err := hasFoundBooks(driver, searchResultTR, notFoundTR)
	if err != nil {
		return Book{}, err
	}

	var book Book
	if foundBooks {
		book, err = bookFounded(driver, bookName)
		if err != nil {
			return Book{}, err
		}
	}

	return book, nil
}
