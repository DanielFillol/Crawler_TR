package main

import (
	"fmt"
	"github.com/Darklabel91/Crawler_TR/CSV"
	"github.com/Darklabel91/Crawler_TR/Crawler"
)

func main() {
	bookNames, err := CSV.ReadCsvFile("CSV/bookTitles.csv")
	if err != nil {
		fmt.Println(err)
	}

	var data []Crawler.Book
	for i := 0; i < 10; i++ {
		book, err := Crawler.Craw(bookNames[i])
		if err != nil {
			fmt.Println(err)
		}

		data = append(data, book)
	}

	fmt.Println(data)
}

func oldMain() {
	//webSiteTR              := "https://www.livrariart.com.br/#&search-term="
	//bookOpenLink           := "//*[@id=\"smarthint-search-products\"]/section/div[2]/div/div/ul/div/li/div/div/div[1]/div[3]/h3/a"
	//searchResultTR         := "//*[@id=\"smarthint-search-result-message\"]"
	//notFoundTR             := "NENHUM RESULTADO PARA DIREITOS FUNDAMENTAIS E JURISDIÇÃO CONSTITUCIONAL: ANÁLISE, CRÍTICA E CONTRIBUIÇÕES. PESSOAS QUE BUSCARAM DIREITOS FUNDAMENTAIS E JURISDIÇÃO CONSTITUCIONAL: ANÁLISE, CRÍTICA E CONTRIBUIÇÕES TAMBÉM BUSCARAM OU COMPRARAM OS PRODUTOS ABAIXO:"
	//productSpecificationTR := "//*[@id=\"caracteristicas\"]/div[2]/div"
	//->main
	//var data []Crawler.Book
	//->main
	//bookNames, err := CSV.ReadCsvFile("/Users/danielfillol/Desktop/Planilha sem título - Página1.csv")
	//if err != nil {
	//fmt.Println(err)
	//}
	//-> main
	//for i := 0; i < 10; i++ {
	//->bookFounded
	//var isbn string
	//var dtDis string
	//var pgs string
	//var yr string
	//end
	//-> Craw
	//driver, err := Crawler.SeleniumWebDriver()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//-> Craw
	//bookSearch := webSiteTR + bookNames[i]
	//-> Craw
	//err = driver.Get(bookSearch)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//end
	//->hasSearchReturn
	//resultElem, err := driver.FindElements(selenium.ByXPATH, searchResultTR)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if len(resultElem) != 0 {
	//	resultText, err := resultElem[0].Text()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//end
	//->findBooks
	//	if resultText != notFoundTR {
	//		bookLink, err := driver.FindElements(selenium.ByXPATH, bookOpenLink)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//end
	//->onlyOneBook
	//if len(bookLink) > 1 {
	//	fmt.Println("Mais de um livro foi encontrado")
	//} else if len(bookLink) != 0 {
	//end
	//->amountSpecification
	//	link, err := bookLink[0].GetAttribute("href")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	err = driver.Get("https:" + link)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	specification, err := driver.FindElements(selenium.ByXPATH, productSpecificationTR)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	if len(specification) != 0 {
	//end
	//->bookFounded
	//for j := 0; j < len(specification); j++ {
	//-> title
	//xpath := "//*[@id=\"caracteristicas\"]/div[2]/div[" + strconv.Itoa(j) + "]/div[1]"
	//titles, err := driver.FindElements(selenium.ByXPATH, xpath)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//if len(titles) != 0 {
	//	title, err := titles[0].Text()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//end
	//->bookFounded
	//				position := j + 1
	//				elemDesc, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position)+"]/div[2]")
	//				if err != nil {
	//					fmt.Println(err)
	//				}
	//
	//				if title == "Código ISBN" {
	//					elemDesc, _ = driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position-1)+"]/div[2]")
	//					isbn, err = elemDesc.Text()
	//					if err != nil {
	//						fmt.Println(err)
	//					}
	//
	//				} else if title == "Data de disponibilidade" {
	//					elemDesc, _ = driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position-1)+"]/div[2]")
	//					dtDis, err = elemDesc.Text()
	//					if err != nil {
	//						fmt.Println(err)
	//					}
	//
	//				} else if title == "Número de páginas" {
	//					elemDesc, _ = driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position-1)+"]/div[2]")
	//					pgs, err = elemDesc.Text()
	//					if err != nil {
	//						fmt.Println(err)
	//					}
	//
	//				} else if title == "Ano de publicação" {
	//					elemDesc, _ = driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position-1)+"]/div[2]")
	//					yr, err = elemDesc.Text()
	//					if err != nil {
	//						fmt.Println(err)
	//					}
	//				}
	//			}
	//		}
	//	}
	//}
	//}
	//}

	//-> main
	//data = append(data, Crawler.Book{
	//	SearchName:    bookNames[i],
	//	ISBN:          isbn,
	//	AvailableDate: dtDis,
	//	Pages:         pgs,
	//	PubYear:       yr,
	//})

	//driver.Close()
	//}

	//fmt.Println(data)
}
